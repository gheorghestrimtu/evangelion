package contracts

import (
	"chainlink-sdet-golang-project/client"
	"chainlink-sdet-golang-project/contracts/ethereum"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// EthereumAggregator acts as a conduit for the ethereum version of the storage contract
type EthereumAggregator struct {
	client       *client.EthereumClient
	aggregator   *ethereum.AccessControlledAggregator
	callerWallet client.BlockchainWallet
}

// NewEthereumAggregator creates a new instance of the aggregator contract for ethereum chains
func NewEthereumAggregator(
	client *client.EthereumClient,
	aggregator *ethereum.AccessControlledAggregator,
	callerWallet client.BlockchainWallet,
) Aggregator {
	return &EthereumAggregator{
		client:       client,
		aggregator:   aggregator,
		callerWallet: callerWallet,
	}
}

func NewEthereumAggregatorContract(ethClient *client.EthereumClient, fromWallet client.BlockchainWallet) (Aggregator, error) {
	instance, err := ethereum.NewAccessControlledAggregator(common.HexToAddress("F570deEffF684D964dc3E15E1F9414283E3f7419"), ethClient.Client)
	if err != nil {
		return nil, err
	}
	return NewEthereumAggregator(ethClient, instance, fromWallet), nil
}

// GetRoundData retrieves a set value from the storage contract
func (e *EthereumAggregator) GetRoundData(ctxt context.Context, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.aggregator.GetRoundData(opts, _roundId)
}
