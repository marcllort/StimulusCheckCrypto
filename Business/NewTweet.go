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

func BasicTweet(config Model.Config, client *twitter.Client, debug bool) {
	publishInfoTweet(config, client, true, debug)
	publishInfoTweet(config, client, false, debug)
}

func AdTweet(client *twitter.Client, debug bool) {
	publishAdTweet(client, true, debug)
	publishAdTweet(client, false, debug)
}

func publishInfoTweet(config Model.Config, client *twitter.Client, firstStimulus bool, debug bool) {
	btc := generatePrice(config, firstStimulus, "BTC")
	eth := generatePrice(config, firstStimulus, "ETH")
	ada := generatePrice(config, firstStimulus, "ADA")
	doge := generatePrice(config, firstStimulus, "DOGE")
	ltc := generatePrice(config, firstStimulus, "LTC")
	xrp := generatePrice(config, firstStimulus, "XRP")
	cro := generatePrice(config, firstStimulus, "CRO")
	bnb := generatePrice(config, firstStimulus, "BNB")

	tweet := generateTweetDataString(firstStimulus, btc, eth, ada, doge, ltc, xrp, cro, bnb)
	fmt.Println(tweet)

	if !debug {
		_, _, err := client.Statuses.Update(tweet, nil)

		if err != nil {
			err.Error()
		}
	}
}

func publishAdTweet(client *twitter.Client, firstAd bool, debug bool) {
	tweet := generateTweetAdString(firstAd)
	fmt.Println(tweet)

	if !debug {
		_, _, err := client.Statuses.Update(tweet, nil)

		if err != nil {
			err.Error()
		}
	}
}

func generatePrice(config Model.Config, firstStimulus bool, currency string) float64 {
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
	} else {
		amountCrypto = past_apiResponse.Data.Data[0].Open * SECOND_STIMULUS_CHECK_DOLLARS
		actualDollars = amountCrypto * actual_apiResponse.USD
	}

	return actualDollars
}

func generateTweetDataString(firstStimulus bool, btcPrice float64, ethPrice float64, adaPrice float64, dogePrice float64,
	ltcPrice float64, xrpPrice float64, croPrice float64, bnbPrice float64) string {
	var tweet string

	btc := strconv.FormatFloat(btcPrice, 'f', 0, 64)
	eth := strconv.FormatFloat(ethPrice, 'f', 0, 64)
	ada := strconv.FormatFloat(adaPrice, 'f', 0, 64)
	doge := strconv.FormatFloat(dogePrice, 'f', 0, 64)
	//ltc := strconv.FormatFloat(ltcPrice, 'f', 0, 64)
	//xrp := strconv.FormatFloat(xrpPrice, 'f', 0, 64)
	cro := strconv.FormatFloat(croPrice, 'f', 0, 64)
	bnb := strconv.FormatFloat(bnbPrice, 'f', 0, 64)

	first_dollars := strconv.FormatFloat(FIRST_STIMULUS_CHECK_DOLLARS, 'f', 0, 64)
	second_dollars := strconv.FormatFloat(SECOND_STIMULUS_CHECK_DOLLARS, 'f', 0, 64)

	tweet = "If you had invested the "
	if firstStimulus {
		tweet += "first stimulus check (" + first_dollars + "$) in: \n"
	} else {
		tweet += "second stimulus check (" + second_dollars + "$) in: \n"
	}
	tweet += "Bitcoin ($BTC) --> " + btc + "$\n"
	tweet += "Etherum ($ETH) --> " + eth + "$\n"
	tweet += "Cardano ($ADA) --> " + ada + "$\n"
	tweet += "Dogecoin ($DOGE) --> " + doge + "$\n"
	tweet += "CryptoCom ($CRO) --> " + cro + "$\n"
	tweet += "Binance ($BNB) --> " + bnb + "$\n"
	//tweet += "Litecoin ($LTC) --> " + ltc + "$\n"
	//tweet += "XRP ($XRP) --> " + xrp + "$\n"

	return tweet
}

func generateTweetAdString(firstAd bool) string {
	var tweet string

	tweet = "Use the following link to get "
	if firstAd {
		tweet += "25$ and a METAL Card with up to 8% CASHBACK when investing in CRO.\n"
		tweet += "https://crypto.com/app/yqtw3a3t8j\n"
	} else {
		tweet += "a Card with up to 8% CASHBACK when investing in BNB, plus the best crypto rates in the market.\n"
		tweet += "https://www.binance.com/en/register?ref=70750475\n"
	}

	return tweet
}

func makeTimestamp(dateTime time.Time) string {
	return strconv.FormatInt(dateTime.UTC().UnixNano()/1000000000, 10)
}
