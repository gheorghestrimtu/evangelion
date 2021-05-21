package contracts

import (
	"context"
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
}
