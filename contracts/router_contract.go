package contracts

import (
	"chainlink-sdet-golang-project/client"
	"chainlink-sdet-golang-project/contracts/ethereum"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EthereumRouter acts as a conduit for the ethereum version of the storage contract
type EthereumRouter struct {
	client       *client.EthereumClient
	router        *ethereum.UniswapV2Router02
	callerWallet client.BlockchainWallet
}

// NewEthereumRouter creates a new instance of the aggregator contract for ethereum chains
func NewEthereumRouter(
	client *client.EthereumClient,
	router *ethereum.UniswapV2Router02,
	callerWallet client.BlockchainWallet,
) Router {
	return &EthereumRouter{
		client:       client,
		router:   router,
		callerWallet: callerWallet,
	}
}

func NewEthereumRouterContract(ethClient *client.EthereumClient, fromWallet client.BlockchainWallet, contractAddress string) (Router, error) {
	instance, err := ethereum.NewUniswapV2Router02(common.HexToAddress(contractAddress), ethClient.Client)
	if err != nil {
		return nil, err
	}
	return NewEthereumRouter(ethClient, instance, fromWallet), nil
}

func (e *EthereumRouter) SwapExactTokensForTokens(ctx context.Context, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.router.SwapExactTokensForTokens(opts, amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumRouter) GetAmountsOut(ctxt context.Context, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.router.GetAmountsOut(opts, amountIn, path)
}

func (e *EthereumRouter) SwapExactETHForTokens(ctx context.Context, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.router.SwapExactETHForTokens(opts, amountOutMin, path, to, deadline)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumRouter) SwapExactETHForTokensSupportingFeeOnTransferTokens(ctx context.Context, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, amountIn, common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.router.SwapExactETHForTokensSupportingFeeOnTransferTokens(opts, amountOutMin, path, to, deadline)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumRouter) WETH(ctxt context.Context) (common.Address, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.router.WETH(opts)
}

func (e *EthereumRouter) SwapExactTokensForETHSupportingFeeOnTransferTokens(ctx context.Context, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.router.SwapExactTokensForETHSupportingFeeOnTransferTokens(opts, amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}