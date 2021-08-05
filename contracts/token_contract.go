package contracts

import (
	"chainlink-sdet-golang-project/client"
	"chainlink-sdet-golang-project/contracts/ethereum"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EthereumAlphaToken acts as a conduit for the ethereum version of the storage contract
type EthereumAlphaToken struct {
	client       *client.EthereumClient
	alpha        *ethereum.ALPHA
	callerWallet client.BlockchainWallet
}

// NewEthereumAlphaToken creates a new instance of the aggregator contract for ethereum chains
func NewEthereumAlphaToken(
	client *client.EthereumClient,
	alpha *ethereum.ALPHA,
	callerWallet client.BlockchainWallet,
) AlphaToken {
	return &EthereumAlphaToken{
		client:       client,
		alpha:   alpha,
		callerWallet: callerWallet,
	}
}

func NewEthereumAlphaTokenContract(ethClient *client.EthereumClient, fromWallet client.BlockchainWallet, contractAddress string) (AlphaToken, error) {
	instance, err := ethereum.NewALPHA(common.HexToAddress(contractAddress), ethClient.Client)
	if err != nil {
		return nil, err
	}
	return NewEthereumAlphaToken(ethClient, instance, fromWallet), nil
}

// Name retrieves a set value from the storage contract
func (e *EthereumAlphaToken) Name(ctxt context.Context) (string, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.alpha.Name(opts)
}

func (e *EthereumAlphaToken) SetDevFeePercent(ctxt context.Context, devFee *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.alpha.SetDevFeePercent(opts, devFee)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) SetTaxFeePercent(ctxt context.Context, taxFee *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.alpha.SetTaxFeePercent(opts, taxFee)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) SetLiquidityFeePercent(ctxt context.Context, liquidityFee *big.Int) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.alpha.SetLiquidityFeePercent(opts, liquidityFee)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) SetDevWalletAddress(ctxt context.Context, devWalletAddress common.Address) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.alpha.SetDevWalletAddress(opts, devWalletAddress)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) ExcludeFromFee(ctxt context.Context, address common.Address) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.alpha.ExcludeFromFee(opts, address)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) ExcludeFromReward(ctxt context.Context, address common.Address) error {
	opts, err := e.client.TransactionOpts(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{})
	if err != nil {
		return err
	}

	_, err = e.alpha.ExcludeFromReward(opts, address)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) Transfer(ctx context.Context, recipient common.Address, amount *big.Int) error {
	opts, err := e.client.TransactionOptsWithCtx(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{}, ctx)
	if err != nil {
		return err
	}

	_, err = e.alpha.Transfer(opts, recipient, amount)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) Approve(ctx context.Context, spender common.Address, amount *big.Int) error {
	opts, err := e.client.TransactionOptsWithCtx(e.callerWallet, common.Address{}, big.NewInt(0), common.Hash{}, ctx)
	if err != nil {
		return err
	}

	_, err = e.alpha.Approve(opts, spender, amount)
	if err != nil {
		return err
	}
	//return e.client.WaitForTransaction(transaction.Hash())
	return nil
}

func (e *EthereumAlphaToken) UniswapV2Router(ctxt context.Context) (common.Address, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.alpha.UniswapV2Router(opts)
}

func (e *EthereumAlphaToken) BalanceOf(ctxt context.Context, account common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(e.callerWallet.Address()),
		Pending: true,
		Context: ctxt,
	}
	return e.alpha.BalanceOf(opts, account)
}