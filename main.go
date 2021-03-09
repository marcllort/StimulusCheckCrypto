package main

import (
	"StimulusCheckCrypto/API"
	"StimulusCheckCrypto/Business"
	"StimulusCheckCrypto/Model"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/jasonlvhit/gocron"
)

func main() {
	config := Model.LoadConfiguration("creds.json")
	client := API.GetTwitterClient(config)

	Business.BasicTweet(config, client)
	executeCronJob(config, client)
}

func executeCronJob(config Model.Config, client *twitter.Client) {
	gocron.Every(1).Day().Do(Business.BasicTweet, config, client)
	<-gocron.Start()
}
