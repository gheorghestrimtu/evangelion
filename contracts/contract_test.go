package contracts

import (
	"chainlink-sdet-golang-project/client"
	"chainlink-sdet-golang-project/config"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Client", func() {
	var conf *config.Config
	var nrOfPreviousRounds = 5

	BeforeEach(func() {
		var err error
		conf, err = config.NewWithPath(config.LocalConfig, "../config")
		Expect(err).ShouldNot(HaveOccurred())
	})

	DescribeTable("interact with the aggregator contract", func(
		initFunc client.BlockchainNetworkInit,
		feedContractAddress string,
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
		latestRound, err := aggregatorInstance.LatestRound(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		for i := 0; i < nrOfPreviousRounds; i++ {
			medianValue, err := aggregatorInstance.GetAnswer(context.Background(), latestRound)
			Expect(err).ShouldNot(HaveOccurred())

			eventsIterator, err := aggregatorInstance.FilterSubmissionReceived(context.Background(), nil, []uint32{uint32(latestRound.Uint64())}, nil)
			Expect(err).ShouldNot(HaveOccurred())

			for eventsIterator.Next() {
				oracleSubmission := eventsIterator.Event.Submission

				medianValueFloat := new(big.Float).SetInt(medianValue)
				oracleSubmissionFloat := new(big.Float).SetInt(oracleSubmission)

				// calculate percentage difference
				subValue := big.NewFloat(0).Sub(medianValueFloat, oracleSubmissionFloat)
				absValue := big.NewFloat(0).Abs(subValue)
				addValue := big.NewFloat(0).Add(medianValueFloat, oracleSubmissionFloat)
				denominator := big.NewFloat(0).Quo(addValue, big.NewFloat(2))
				divResult := big.NewFloat(0).Quo(absValue, denominator)
				percentageDifference := big.NewFloat(0).Mul(divResult, big.NewFloat(100))

				Expect(percentageDifference.Float64()).Should(BeNumerically("<", deviationThreshold))
			}
			err = eventsIterator.Error()
			Expect(err).ShouldNot(HaveOccurred())

			latestRound.Sub(latestRound, big.NewInt(1))
		}
	},
		Entry("on Ethereum Mainnet with round id", client.NewMainnetNetwork, "0xF570deEffF684D964dc3E15E1F9414283E3f7419", 10),
	)
})
