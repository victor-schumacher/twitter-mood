package twitter

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Auth() *twitter.Client {
	// 2zr105kPhmM6qxmkqHeYsl6mC
	// 7MjZeewsCjLaAZriEr68bJZaFEARmYO1OdrPznXY9sypHWTsVV - secret
	config := oauth1.NewConfig(os.Getenv("key"), os.Getenv("secretKey"))

	// 1130660432681799681-gKk2Qf1LJpuWGGBDhunnCM5TsBQPB9
	// 7GNSMAnva0V78xCpqdJ433H07zeMbkFERdrYmbeomleH1 - secret
	token := oauth1.NewToken(os.Getenv("accessToken"), os.Getenv("accessTokenSecret"))

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client

	params := twitter.StatusUpdateParams{
		Status:             "z",
		InReplyToStatusID:  0,
		PossiblySensitive:  nil,
		Lat:                nil,
		Long:               nil,
		PlaceID:            "",
		DisplayCoordinates: nil,
		TrimUser:           nil,
		MediaIds:           nil,
		TweetMode:          "",
	}

	fmt.Println(params)
}
