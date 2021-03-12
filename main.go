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

	debug := false

	go executeInfoCronJob(config, client, debug)
	go executeAdCronJob(client, debug)

	for {
	}
}

func executeInfoCronJob(config Model.Config, client *twitter.Client, debug bool) {
	Business.BasicTweet(config, client, debug)
	gocron.Every(1).Day().Do(Business.BasicTweet, config, client, debug)
	<-gocron.Start()
}

func executeAdCronJob(client *twitter.Client, debug bool) {
	Business.AdTweet(client, debug)
	gocron.Every(4).Day().Do(Business.AdTweet, client, debug)
	<-gocron.Start()
}
