package API

import (
	"StimulusCheckCrypto/Model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const PAST_API_URL string = "https://min-api.cryptocompare.com/data/v2/histoday?fsym=USD&tsym=$crypto&limit=1&aggregate=1&toTs=$timestamp&api_key=$apikey"
const ACTUAL_API_URL string = "https://min-api.cryptocompare.com/data/price?fsym=$crypto_origin&tsyms=$crypto_dest&api_key=$apikey"

func PastPrice(currency string, timestamp string, api_key string) Model.PastPrice {
	var pastPrice Model.PastPrice

	url := strings.ReplaceAll(PAST_API_URL, "$crypto", currency)
	url = strings.ReplaceAll(url, "$timestamp", timestamp)
	url = strings.ReplaceAll(url, "$apikey", api_key)

	past_api_resp, err := http.Get(url)
	if err != nil {
		err.Error()
	}
	past_body_bytes, _ := ioutil.ReadAll(past_api_resp.Body)

	json.Unmarshal(past_body_bytes, &pastPrice)

	return pastPrice
}

func ActualPrice(currency string, currencies string, api_key string) Model.ActualPrice {
	var actualPrice Model.ActualPrice

	url := strings.ReplaceAll(ACTUAL_API_URL, "$crypto_origin", currency)
	url = strings.ReplaceAll(url, "$crypto_dest", currencies)
	url = strings.ReplaceAll(url, "$apikey", api_key)

	actual_api_resp, err := http.Get(url)
	if err != nil {
		err.Error()
	}
	actual_body_bytes, _ := ioutil.ReadAll(actual_api_resp.Body)

	json.Unmarshal(actual_body_bytes, &actualPrice)

	return actualPrice
}
