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
		feedContractAddress string,
		roundId uint32,
		deviationThreshold int,
	) {
		// Instantiate contract
		networkConfig, err := initFunc(conf)
		Expect(err).ShouldNot(HaveOccurred())
		client, err := client.NewEthereumClient(networkConfig)
		Expect(err).ShouldNot(HaveOccurred())
		wallets, err := networkConfig.Wallets()
		Expect(err).ShouldNot(HaveOccurred())

		aggregatorInstance, err := NewEthereumAggregatorContract(client, wallets.Default(), feedContractAddress)
		Expect(err).ShouldNot(HaveOccurred())

		// Interact with contract

		// Get oracles
		oracles, err := aggregatorInstance.GetOracles(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		// Get median value for round
		roundData, err := aggregatorInstance.GetRoundData(context.Background(), big.NewInt(int64(roundId)))
		Expect(err).ShouldNot(HaveOccurred())

		medianValue := roundData.Answer
		fmt.Println(medianValue)

		// Get round value for oracle
		// read events
		eventsIterator, err := aggregatorInstance.FilterSubmissionReceived(context.Background(), nil, []uint32{roundId}, oracles)
		for eventsIterator.Next() {
			fmt.Println(eventsIterator.Event.Oracle)
			fmt.Println(eventsIterator.Event.Submission)
		}

		//val, err := storeInstance.Get(context.Background())
		//fmt.Println(&val)
		//Expect(err).ShouldNot(HaveOccurred())
		//Expect(val).To(Equal(roundId))
	},
		Entry("on Ethereum Mainnet with round id", client.NewMainnetNetwork, "0xF570deEffF684D964dc3E15E1F9414283E3f7419", uint32(2400), 10),
		Entry("on Ethereum Mainnet with round id", client.NewMainnetNetwork, "0xF570deEffF684D964dc3E15E1F9414283E3f7419", uint32(10034), 10),
		Entry("on Ethereum Mainnet with round id", client.NewMainnetNetwork, "0xF570deEffF684D964dc3E15E1F9414283E3f7419", uint32(15342), 10),
	)
})
