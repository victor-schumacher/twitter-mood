package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/victor-schumacher/twitter-mood/pkg/data"
	"github.com/victor-schumacher/twitter-mood/pkg/tw"
	"strings"
	"github.com/AlecAivazis/survey/v2"
)

func main() {

	track := ""
	prompt := &survey.Input{
		Message: "Enter your track word",
	}
	survey.AskOne(prompt, &track)



	client := tw.Auth()

	params := &twitter.StreamFilterParams{
		Track:         []string{track},
		StallWarnings: twitter.Bool(true),
		Language:      []string{"en"},
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		fmt.Println(err)
	}

	dm := twitter.NewSwitchDemux()
	tweetCount := 0.0
	goodTweetCount := 0.0
	dm.Tweet = func(tweet *twitter.Tweet) {
		tweetCount++
		if data.Analyze(tweet.Text) {
			goodTweetCount++
		}
		tweetCount100 := tweetCount/100
		goodDecimal := goodTweetCount / tweetCount100
		s := fmt.Sprintf("%.2f%%", goodDecimal)
		s = strings.Replace(s, ".", ",", -1)

		fmt.Printf("TWEET COUNT:%.0f\n\n", tweetCount)
		fmt.Printf("TRACK WORD:%s\n\n",track)
		fmt.Printf("POSITIVE: %v \n\n", s)

	}

	for message := range stream.Messages {
		dm.Handle(message)
	}

}
