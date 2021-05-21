package contracts

import (
	"chainlink-sdet-golang-project/client"
	"chainlink-sdet-golang-project/config"
	"context"
	"fmt"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var conf *config.Config

	BeforeEach(func() {
		var err error
		conf, err = config.NewWithPath(config.LocalConfig, "../config")
		Expect(err).ShouldNot(HaveOccurred())
	})

	DescribeTable("interact with the aggregator contract", func(
		initFunc client.BlockchainNetworkInit,
		value *big.Int,
	) {
		// Deploy contract
		networkConfig, err := initFunc(conf)
		Expect(err).ShouldNot(HaveOccurred())
		client, err := client.NewEthereumClient(networkConfig)
		Expect(err).ShouldNot(HaveOccurred())
		wallets, err := networkConfig.Wallets()
		Expect(err).ShouldNot(HaveOccurred())

		aggregatorInstance, err := NewEthereumAggregatorContract(client, wallets.Default())
		Expect(err).ShouldNot(HaveOccurred())

		// Interact with contract
		roundValue, err := aggregatorInstance.GetRoundData(context.Background(), value)
		Expect(err).ShouldNot(HaveOccurred())
		//val, err := storeInstance.Get(context.Background())
		fmt.Println("adsfd")
		fmt.Println(roundValue)
		//fmt.Println(&val)
		//Expect(err).ShouldNot(HaveOccurred())
		//Expect(val).To(Equal(value))
	},
		Entry("on Ethereum Mainnet with round id", client.NewMainnetNetwork, big.NewInt(18832)),
	)
})
