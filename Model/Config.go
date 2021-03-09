package Model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	TwitterConsumerKey    string `json:"twitterConsumerKey"`
	TwitterConsumerSecret string `json:"twitterConsumerSecret"`
	TwitterAccessToken    string `json:"twitterAccessToken"`
	TwitterAccessSecret   string `json:"twitterAccessSecret"`
	CryptoCompareApiKey   string `json:"cryptoCompareApiKey"`
	CoinbaseId            string `json:"coinbaseId"`
	CoinbaseSecret        string `json:"coinbaseSecret"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
