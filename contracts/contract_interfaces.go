package contracts

import (
	"chainlink-sdet-golang-project/contracts/ethereum"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Aggregator interface {
	GetRoundData(context.Context, *big.Int) (struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}, error)
	Description(context.Context) (string, error)
	GetOracles(context.Context) ([]common.Address, error)
	OracleRoundState(context.Context, common.Address, uint32) (struct {
		EligibleToSubmit bool
		RoundId          uint32
		LatestSubmission *big.Int
		StartedAt        uint64
		Timeout          uint64
		AvailableFunds   *big.Int
		OracleCount      uint8
		PaymentAmount    *big.Int
	}, error)
	FilterSubmissionReceived(context.Context, []*big.Int, []uint32, []common.Address) (*ethereum.AccessControlledAggregatorSubmissionReceivedIterator, error)
	LatestRound(context.Context) (*big.Int, error)
	GetAnswer(context.Context, *big.Int) (*big.Int, error)
}
