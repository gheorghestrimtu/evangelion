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

type AlphaToken interface {
	Name(ctx context.Context) (string, error)

	SetDevFeePercent(ctx context.Context, devFee *big.Int) error
	SetTaxFeePercent(ctx context.Context, taxFee *big.Int) error
	SetLiquidityFeePercent(ctx context.Context, liquidityFee *big.Int) error
	SetDevWalletAddress(ctx context.Context, devWalletAddress common.Address) error
	ExcludeFromFee(ctx context.Context, account common.Address) error
	ExcludeFromReward(ctx context.Context, account common.Address) error
	Transfer(ctx context.Context, recipient common.Address, amount *big.Int) error

	Approve(ctx context.Context, spender common.Address, amount *big.Int) error
	UniswapV2Router(ctxt context.Context) (common.Address, error)
	BalanceOf(ctxt context.Context, account common.Address) (*big.Int, error)
}

type Router interface {
	SwapExactTokensForTokens(ctx context.Context, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error
	SwapExactETHForTokens(ctx context.Context, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error
	GetAmountsOut(ctx context.Context, amountIn *big.Int, path []common.Address) ([]*big.Int, error)
	SwapExactETHForTokensSupportingFeeOnTransferTokens(ctx context.Context, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error
	WETH(ctx context.Context) (common.Address, error)
	SwapExactTokensForETHSupportingFeeOnTransferTokens(ctx context.Context, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error
}
