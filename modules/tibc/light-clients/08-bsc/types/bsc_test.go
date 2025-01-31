package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/bianjieai/tibc-go/simapp"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type BSCTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *simapp.SimApp
}

func (suite *BSCTestSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.app = app
}

func TestBSCTestSuite(t *testing.T) {
	suite.Run(t, new(BSCTestSuite))
}
