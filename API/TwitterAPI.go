package API

import (
	"StimulusCheckCrypto/Model"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func GetTwitterClient(config Model.Config) *twitter.Client {
	config_twitter := oauth1.NewConfig(config.TwitterConsumerKey, config.TwitterConsumerSecret)
	token := oauth1.NewToken(config.TwitterAccessToken, config.TwitterAccessSecret)
	httpClient := config_twitter.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}
