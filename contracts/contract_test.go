package contracts

import (
	"chainlink-sdet-golang-project/client"
	"chainlink-sdet-golang-project/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
	"strings"
	"sync"
	"time"
)

func etherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

var _ = Describe("Client", func() {
	var conf *config.Config

	var initFunc client.BlockchainNetworkInit = client.NewRopstenNetwork
	var tokenAddress = "0xDA0F503a482c3d52b73f4fb710DF977A22c77703"
	var devWalletAddress = common.HexToAddress("0xB000C5b54E1fdf683Bebc94065eDeC1F1d5718b7")
	var deadAddress = common.HexToAddress("0x000000000000000000000000000000000000dead")

	var networkConfig client.BlockchainNetwork
	var ethereumClient *client.EthereumClient
	var wallets client.BlockchainWallets

	var tokenInstanceOwner AlphaToken

	//BuyTokensWithWETH := func(wallet client.BlockchainWallet, wethAmount * big.Float, wg *sync.WaitGroup) {
	//	defer wg.Done()
	//
	//	var recipientAddress = common.HexToAddress(wallet.Address())
	//
	//	routerAddress, err := tokenInstanceOwner.UniswapV2Router(context.Background())
	//	Expect(err).ShouldNot(HaveOccurred())
	//
	//	routerInstance, err := NewEthereumRouterContract(ethereumClient, wallet, routerAddress.String())
	//	Expect(err).ShouldNot(HaveOccurred())
	//
	//	wethAddress, err := routerInstance.WETH(context.Background())
	//	Expect(err).ShouldNot(HaveOccurred())
	//
	//	var amountIn = etherToWei(wethAmount)
	//	var tokenIn = wethAddress
	//	var tokenOut = common.HexToAddress(tokenAddress)
	//	var path = make([]common.Address, 2)
	//	path[0] = tokenIn
	//	path[1] = tokenOut
	//
	//	amountOutMin := big.NewInt(0)
	//
	//	deadline := big.NewInt(time.Now().Unix() + 1000*60*10)
	//
	//	err = routerInstance.SwapExactETHForTokensSupportingFeeOnTransferTokens(
	//		context.Background(),
	//		amountIn,
	//		amountOutMin,
	//		path,
	//		recipientAddress,
	//		deadline)
	//
	//	Expect(err).ShouldNot(HaveOccurred())
	//}

	BuyETHWithToken := func(wallet client.BlockchainWallet, wg *sync.WaitGroup) {
		defer wg.Done()

		tokenInstance, err := NewEthereumAlphaTokenContract(ethereumClient, wallet, tokenAddress)
		Expect(err).ShouldNot(HaveOccurred())

		tokenBalance, _ := tokenInstance.BalanceOf(context.Background(), common.HexToAddress(wallet.Address()))

		var recipientAddress = common.HexToAddress(wallet.Address())

		MaxInt := new(big.Int)
		MaxInt, ok := MaxInt.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		if !ok {
			fmt.Println("SetString: error")
			return
		}

		routerAddress, err := tokenInstance.UniswapV2Router(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		routerInstance, err := NewEthereumRouterContract(ethereumClient, wallet, routerAddress.String())
		Expect(err).ShouldNot(HaveOccurred())

		wethAddress, err := routerInstance.WETH(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		var tokenIn = common.HexToAddress(tokenAddress)
		var tokenOut = wethAddress
		var path = make([]common.Address, 2)
		path[0] = tokenIn
		path[1] = tokenOut

		amountOutMin := big.NewInt(0)

		deadline := big.NewInt(time.Now().Unix() + 1000*60*10)

		err = tokenInstance.Approve(context.Background(), routerAddress, MaxInt)
		Expect(err).ShouldNot(HaveOccurred())

		time.Sleep(2 * time.Minute)

		err = routerInstance.SwapExactTokensForETHSupportingFeeOnTransferTokens(
			context.Background(),
			tokenBalance,
			amountOutMin,
			path,
			recipientAddress,
			deadline)

		Expect(err).ShouldNot(HaveOccurred())
	}

	BeforeEach(func() {
		var err error
		conf, err = config.NewWithPath(config.LocalConfig, "../config")
		Expect(err).ShouldNot(HaveOccurred())

		// Instantiate contract
		networkConfig, err = initFunc(conf)
		Expect(err).ShouldNot(HaveOccurred())
		ethereumClient, err = client.NewEthereumClient(networkConfig)
		Expect(err).ShouldNot(HaveOccurred())
		wallets, err = networkConfig.Wallets()
		Expect(err).ShouldNot(HaveOccurred())

		ownerAddress, _ := wallets.Wallet(0)
		tokenInstanceOwner, err = NewEthereumAlphaTokenContract(ethereumClient, ownerAddress, tokenAddress)
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Set all fees", func() {
		name, err := tokenInstanceOwner.Name(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		fmt.Println(name)

		// DEV Fee
		err = tokenInstanceOwner.SetDevFeePercent(context.Background(), big.NewInt(4))
		Expect(err).ShouldNot(HaveOccurred())

		// TAX Fee
		err = tokenInstanceOwner.SetTaxFeePercent(context.Background(), big.NewInt(4))
		Expect(err).ShouldNot(HaveOccurred())

		// LIQUIDITY Fee
		err = tokenInstanceOwner.SetLiquidityFeePercent(context.Background(), big.NewInt(2))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Set dev address", func() {
		err := tokenInstanceOwner.SetDevWalletAddress(context.Background(), devWalletAddress)
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Exclude from fee", func() {
		err := tokenInstanceOwner.ExcludeFromFee(context.Background(), common.HexToAddress("0xA895Df836Ac98f31d96c2B82f07E2cc276f7f031"))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Exclude from reward", func() {
		err := tokenInstanceOwner.ExcludeFromReward(context.Background(), common.HexToAddress("0x5594bCcBA7019a661AC3fE25f4f7E97F2e4ed44c"))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Transfer", func() {
		x := big.NewInt(0)
		x.SetString("500000000000000000000000", 10)
		err := tokenInstanceOwner.Transfer(context.Background(), deadAddress, x)
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Buy tokens with WETH", func() {
		routerAddress, err := tokenInstanceOwner.UniswapV2Router(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		walletAddress, err := wallets.Wallet(2)
		Expect(err).ShouldNot(HaveOccurred())

		var recipientAddress = common.HexToAddress(walletAddress.Address())

		routerInstance, err := NewEthereumRouterContract(ethereumClient, walletAddress, routerAddress.String())
		Expect(err).ShouldNot(HaveOccurred())

		wethAddress, err := routerInstance.WETH(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		var amountIn = etherToWei(big.NewFloat(0.1))
		var tokenIn = wethAddress
		var tokenOut = common.HexToAddress(tokenAddress)
		var path = make([]common.Address, 2)
		path[0] = tokenIn
		path[1] = tokenOut

		//amounts, err := routerInstance.GetAmountsOut(context.Background(), amountIn, path)
		//amountOutMin := amounts[1]
		amountOutMin := big.NewInt(0)

		deadline := big.NewInt(time.Now().Unix() + 1000*60*10)

		//fmt.Println("amounts: ", amounts)
		fmt.Println("Amount In (WEI): ", amountIn)
		fmt.Println("Token In (WBNB): ", tokenIn.String())
		fmt.Println("Token Out: ", tokenOut.String())
		fmt.Println("Amount Out Min: ", amountOutMin)
		fmt.Println("deadline: ", deadline)

		err = routerInstance.SwapExactETHForTokensSupportingFeeOnTransferTokens(
			context.Background(),
			amountIn,
			amountOutMin,
			path,
			recipientAddress,
			deadline)

		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Buy ETH with tokens", func() {
		var recipientAddress = common.HexToAddress("0xa8A8C12098876a320Ee72f8dBB19616889B4c713")

		MaxInt := new(big.Int)
		MaxInt, ok := MaxInt.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		if !ok {
			fmt.Println("SetString: error")
			return
		}

		routerAddress, err := tokenInstanceOwner.UniswapV2Router(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		routerInstance, err := NewEthereumRouterContract(ethereumClient, wallets.Default(), routerAddress.String())
		Expect(err).ShouldNot(HaveOccurred())

		wethAddress, err := routerInstance.WETH(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		amountIn := new(big.Int)
		amountIn, ok = MaxInt.SetString("23868827809988000000000", 10)
		if !ok {
			fmt.Println("SetString: error")
			return
		}

		var tokenIn = common.HexToAddress(tokenAddress)
		var tokenOut = wethAddress
		var path = make([]common.Address, 2)
		path[0] = tokenIn
		path[1] = tokenOut

		//amounts, err := routerInstance.GetAmountsOut(context.Background(), amountIn, path)
		//amountOutMin := amounts[1]
		amountOutMin := big.NewInt(0)

		deadline := big.NewInt(time.Now().Unix() + 1000*60*10)

		//fmt.Println("amounts: ", amounts)
		fmt.Println("Amount In (WEI): ", amountIn)
		fmt.Println("Token In (WBNB): ", tokenIn.String())
		fmt.Println("Token Out: ", tokenOut.String())
		fmt.Println("Amount Out Min: ", amountOutMin)
		fmt.Println("deadline: ", deadline)

		err = tokenInstanceOwner.Approve(context.Background(), routerAddress, MaxInt)
		Expect(err).ShouldNot(HaveOccurred())

		err = routerInstance.SwapExactTokensForETHSupportingFeeOnTransferTokens(
			context.Background(),
			amountIn,
			amountOutMin,
			path,
			recipientAddress,
			deadline)

		Expect(err).ShouldNot(HaveOccurred())
	})

	FIt("Increase market cap", func() {
		//var nrOfWallets = 5

		walletsInAction := 2

		var wg sync.WaitGroup

		for i := 0; i < 1; i++ {
			for j := 1; j <= walletsInAction; j++ {
				nodeWallet, err := wallets.Wallet(j)
				Expect(err).ShouldNot(HaveOccurred())

				wg.Add(1)
				//go BuyTokensWithWETH(nodeWallet, big.NewFloat(0.1), &wg)
				go BuyETHWithToken(nodeWallet, &wg)
			}
			wg.Wait()
		}

		//time.Sleep(2 * time.Minute)
	})

})
