package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	packetType "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	routingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
)

const (
	PREFIX = "tibc/nft"

	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
)

func (k Keeper) SendNftTransfer(
	ctx sdk.Context,
	class, id string,
	sender sdk.AccAddress,
	receiver, destChain, relayChain string,
) error {
	// class must be existed
	_, found := k.nk.GetDenom(ctx, class)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "class %s not existed ", class)
	}
	// get nft
	nft, err := k.nk.GetNFT(ctx, class, id)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", id, class)
	}

	// decode the sender address
	sender, err = sdk.AccAddressFromBech32(sender.String())
	if err != nil {
		return err
	}

	moudleAddr := k.GetNftTransferModuleAddr(types.ModuleName)

	// sourceChain cannot be equal to destChain
	sourceChain := k.ck.GetChainName(ctx)
	if sourceChain == destChain {
		return sdkerrors.Wrapf(types.ErrScChainEqualToDestChain, "invalid destChain %s equals to scChain %s", destChain, sourceChain)
	}

	fullClassPath := class

	// deconstruct the nft class into the class trace info to determine if the sender is the source chain
	if strings.HasPrefix(class, "tibc-") {
		fullClassPath, err = k.ClassPathFromHash(ctx, class)
		if err != nil {
			return err
		}
	}

	// determine whether nft is sent from the source chain or sent back to the source chain from other chains
	awayFromOrigin := k.determineAwayFromOrigin(fullClassPath, destChain)

	// get the next sequence
	sequence, _ := k.pk.GetNextSequenceSend(ctx, sourceChain, destChain)

	if awayFromOrigin {
		// Two conversion scenarios
		// 1. nftClass -> tibc-hash(nft/A/B/nftClass)
		// 2. tibc-hash(nft/A/B/nftClass) -> tibc-hash(nft/A/B/C/nftClass)

		// Two things need to be done
		// 1. lock nft  |send to moduleAccount
		// 2. send packet
		// The nft attribute must be marked as unchanged (it cannot be changed in practice)
		// because the TransferOwner method will verify when UpdateRestricted is true
		if err := k.nk.TransferOwner(ctx, class, id, DoNotModify, DoNotModify, DoNotModify, sender, moudleAddr); err != nil {
			return err
		}
	} else {
		// burn nft
		if err := k.nk.BurnNFT(ctx, class, id, sender); err != nil {
			return err
		}
	}

	// constructs packet
	packetData := types.NewNonFungibleTokenPacketData(
		fullClassPath,
		id,
		nft.GetURI(),
		sender.String(),
		receiver,
		awayFromOrigin,
	)

	packet := packetType.NewPacket(packetData.GetBytes(), sequence, sourceChain, destChain, relayChain, string(routingtypes.NFT))

	// send packet
	if err := k.pk.SendPacket(ctx, packet); err != nil {
		return err
	}

	return nil
}

/*
OnRecvPacket
A->B->C  away_from_source == true
	B receive packet from A : class -> nft/A/B/class
	c receive packet from B : nft/A/B/class -> nft/A/B/C/class
C->B->A  away_from_source == flase
	B receive packet from C : nft/A/B/C/class -> nft/A/B/class
	A receive packet from B : nft/A/B/class -> class
*/
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet packetType.Packet, data types.NonFungibleTokenPacketData) error {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return err
	}

	// decode the sender address
	receiver, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		return err
	}

	moudleAddr := k.GetNftTransferModuleAddr(types.ModuleName)

	var newClass string
	if data.AwayFromOrigin {
		if strings.HasPrefix(data.Class, "nft") {
			// nft/A/B/class -> nft/A/B/C/class
			// [nft][A][B][class] -> [nft][A][B][C][class]
			classSplit := strings.Split(data.Class, "/")
			classSplit = append(classSplit[:len(classSplit)-1], append([]string{packet.DestinationChain}, classSplit[len(classSplit)-1:]...)...)
			newClass = strings.Join(classSplit, "/")
		} else {
			// class -> nft/A/B/class
			newClass = "nft" + "/" + packet.SourceChain + "/" + packet.DestinationChain + "/" + data.Class
		}

		// construct the class trace from the full raw class
		classTrace := types.ParseClassTrace(newClass)

		traceHash := classTrace.Hash()

		if !k.HasClassTrace(ctx, traceHash) {
			k.SetClassTrace(ctx, classTrace)
		}

		voucherClass := classTrace.IBCClass()

		_, found := k.nk.GetDenom(ctx, voucherClass)
		if !found {
			// The creator of cross-chain denom must be a module account,
			// and only the owner of Denom can issue NFT under this category,
			// and no one under this category can update NFT,
			// that is, updateRestricted is true and mintRestricted is true
			if err := k.nk.IssueDenom(ctx, voucherClass, "", "", "", moudleAddr, true, true); err != nil {
				return err
			}
		}

		// Only module accounts can mint nft, because mintRestricted is true,
		// you must first mint nft to the module account, and then transfer nft ownership to the receiver
		if err := k.nk.MintNFT(ctx, voucherClass, data.Id, "", data.Uri, "", moudleAddr); err != nil {
			return err
		}

		if err := k.nk.TransferOwner(ctx, voucherClass, data.Id, DoNotModify, DoNotModify, DoNotModify, moudleAddr, receiver); err != nil {
			return err
		}

	} else {
		if strings.HasPrefix(data.Class, "nft") {
			classSplit := strings.Split(data.Class, "/")

			if len(classSplit) == 4 {
				// nft/A/B/class -> class
				newClass = classSplit[len(classSplit)-1]
			} else {
				// nft/A/B/C/class -> nft/A/B/class
				classSplit = append(classSplit[:len(classSplit)-2], classSplit[len(classSplit)-1])
				newClass = strings.Join(classSplit, "/")
			}

			classTrace := types.ParseClassTrace(newClass)
			voucherClass := classTrace.IBCClass()
			// unlock
			if err := k.nk.TransferOwner(ctx, voucherClass, data.Id, DoNotModify, DoNotModify, DoNotModify, moudleAddr, receiver); err != nil {
				return err
			}
		} else {
			return sdkerrors.Wrapf(types.ErrInvalidDenom, "class has no prefix: %s", data.Class)
		}
	}
	return nil
}

func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, data types.NonFungibleTokenPacketData, ack packetType.Acknowledgement) error {
	switch ack.Response.(type) {
	case *packetType.Acknowledgement_Error:
		return k.refundPacketToken(ctx, data)
	default:
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	}
}

func (k Keeper) refundPacketToken(ctx sdk.Context, data types.NonFungibleTokenPacketData) error {
	// decode the sender address
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return err
	}

	moudleAddr := k.GetNftTransferModuleAddr(types.ModuleName)

	classTrace := types.ParseClassTrace(data.Class)
	voucherClass := classTrace.IBCClass()

	if data.AwayFromOrigin {
		// unlock
		if err := k.nk.TransferOwner(ctx, voucherClass, data.Id, DoNotModify, DoNotModify, DoNotModify,
			k.GetNftTransferModuleAddr(types.ModuleName), sender); err != nil {
			return err
		}
	} else {
		// mintNFT
		// Corresponding to burnNft, because the mintRestricted attribute of denom generated by any cross-chain nft is true,
		// so to re-mint nft, you must first mintnft to the module account, and then transfer the nft ownership to the sender account
		if err := k.nk.MintNFT(ctx, voucherClass, data.Id, "", data.Uri, "", moudleAddr); err != nil {
			return err
		}

		if err := k.nk.TransferOwner(ctx, voucherClass, data.Id, DoNotModify, DoNotModify, DoNotModify, moudleAddr, sender); err != nil {
			return err
		}
	}
	return nil
}

// determineAwayFromOrigin determine whether nft is sent from the source chain or sent back to the source chain from other chains
func (k Keeper) determineAwayFromOrigin(class, destChain string) (awayFromOrigin bool) {
	/*
		-- not has prefix
		1. A -> B  class:class | sourceChain:A  | destChain:B |awayFromOrigin = true
		-- has prefix
		1. B -> C    class:nft/A/B/class 	| sourceChain:B  | destChain:C |awayFromOrigin = true
		2. C -> B    class:nft/A/B/C/class  | sourceChain:C  | destChain:B |awayFromOrigin = false
		3. B -> A    class:nft/A/B/class 	| sourceChain:B  | destChain:A |awayFromOrigin = false
	*/
	if !strings.HasPrefix(class, "nft") {
		return true
	}

	classSplit := strings.Split(class, "/")
	if classSplit[len(classSplit)-3] == destChain {
		return false
	}
	return true
}

// ClassPathFromHash returns the full class path prefix from an ibc class with a hash
// component.
func (k Keeper) ClassPathFromHash(ctx sdk.Context, class string) (string, error) {
	// trim the class prefix, by default "tibc-"
	hexHash := class[len(types.ClassPrefix+"-"):]

	hash, err := types.ParseHexHash(hexHash)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidDenom, err.Error())
	}

	denomTrace, found := k.GetClassTrace(ctx, hash)
	if !found {
		return "", sdkerrors.Wrap(types.ErrTraceNotFound, hexHash)
	}

	fullDenomPath := denomTrace.GetFullClassPath()
	return fullDenomPath, nil
}