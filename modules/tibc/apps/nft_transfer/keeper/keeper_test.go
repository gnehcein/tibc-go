package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"

	"github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibctesting "github.com/bianjieai/tibc-go/modules/tibc/testing"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *tibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *tibctesting.TestChain
	chainB *tibctesting.TestChain
	chainC *tibctesting.TestChain
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = tibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(tibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(tibctesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(tibctesting.GetChainID(2))
}

func (suite *KeeperTestSuite) TestGetTransferMoudleAddr() {
	expectedMaccAddr := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	macc := suite.chainA.App.NftTransferKeeper.GetNftTransferModuleAddr(types.ModuleName)

	suite.Require().NotNil(macc)
	suite.Require().Equal(expectedMaccAddr, macc)
}

func NewTransferPath(scChain, destChain *tibctesting.TestChain) *tibctesting.Path {
	path := tibctesting.NewPath(scChain, destChain)
	return path
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
