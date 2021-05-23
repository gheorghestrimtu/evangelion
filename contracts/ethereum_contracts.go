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

func NewEthereumAggregatorContract(ethClient *client.EthereumClient, fromWallet client.BlockchainWallet, contractAddress string) (Aggregator, error) {
	instance, err := ethereum.NewAccessControlledAggregator(common.HexToAddress(contractAddress), ethClient.Client)
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

func (e *EthereumAggregator) Description(ctxt context.Context) (string, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.aggregator.Description(opts)
}

func (e *EthereumAggregator) GetOracles(ctxt context.Context) ([]common.Address, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.aggregator.GetOracles(opts)
}

func (e *EthereumAggregator) OracleRoundState(ctxt context.Context, _oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.aggregator.OracleRoundState(opts, _oracle, _queriedRoundId)
}

func (e *EthereumAggregator) LatestRound(ctxt context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.aggregator.LatestRound(opts)
}

func (e *EthereumAggregator) GetAnswer(ctxt context.Context, _roundId *big.Int) (*big.Int, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.aggregator.GetAnswer(opts, _roundId)
}

func (e *EthereumAggregator) FilterSubmissionReceived(ctxt context.Context, submission []*big.Int, round []uint32, oracle []common.Address) (*ethereum.AccessControlledAggregatorSubmissionReceivedIterator, error) {
	opts := &bind.FilterOpts{
		Start:   1,
		End:     nil,
		Context: ctxt,
	}
	return e.aggregator.FilterSubmissionReceived(opts, submission, round, oracle)
}
