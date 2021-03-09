package Business

import (
	"StimulusCheckCrypto/API"
	"StimulusCheckCrypto/Model"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"strconv"
	"time"
)

const FIRST_STIMULUS_CHECK_DOLLARS float64 = 1200
const SECOND_STIMULUS_CHECK_DOLLARS float64 = 1400

func BasicTweet(config Model.Config, client *twitter.Client) {
	publishTweet(config, client, true, "ETH")
	publishTweet(config, client, true, "BTC")
}

func publishTweet(config Model.Config, client *twitter.Client, firstStimulus bool, currency string) {
	tweet := generateTweet(config, firstStimulus, currency)
	fmt.Println(tweet)
	_, _, err := client.Statuses.Update(tweet, nil)

	if err != nil {
		err.Error()
	}
}

func generateTweet(config Model.Config, firstStimulus bool, currency string) string {
	var tweet string
	var timestamp string
	var amountCrypto float64
	var actualDollars float64

	if firstStimulus {
		timestamp = makeTimestamp(time.Date(2020, 04, 18, 0, 0, 0, 0, time.UTC))
	} else {
		timestamp = makeTimestamp(time.Date(2021, 03, 14, 0, 0, 0, 0, time.UTC))
	}

	past_apiResponse := API.PastPrice(currency, timestamp, config.CryptoCompareApiKey)
	actual_apiResponse := API.ActualPrice(currency, "BTC,ETH,USD", config.CryptoCompareApiKey)

	if firstStimulus {
		amountCrypto = past_apiResponse.Data.Data[0].Open * FIRST_STIMULUS_CHECK_DOLLARS
		actualDollars = amountCrypto * actual_apiResponse.USD

		tweet = "If you had invested the stimulus check in " + currency + " (" +
			strconv.FormatFloat(FIRST_STIMULUS_CHECK_DOLLARS, 'f', 0, 64) + "$) now it would be " +
			strconv.FormatFloat(actualDollars, 'f', 0, 64) + "$"
	} else {
		amountCrypto = past_apiResponse.Data.Data[0].Open * SECOND_STIMULUS_CHECK_DOLLARS
		actualDollars = amountCrypto * actual_apiResponse.USD

		tweet = "If you had invested the stimulus check in " + currency + " (" +
			strconv.FormatFloat(SECOND_STIMULUS_CHECK_DOLLARS, 'f', 0, 64) + "$) now it would be " +
			strconv.FormatFloat(actualDollars, 'f', 0, 64) + "$"
	}

	return tweet
}

func makeTimestamp(dateTime time.Time) string {
	return strconv.FormatInt(dateTime.UTC().UnixNano()/1000000000, 10)
}
