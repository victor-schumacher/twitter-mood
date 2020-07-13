package tw

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Auth() *twitter.Client {
	accessToken:=""
	accessTokenSecret:=""
	key:=""
	secretKey:=""


	config := oauth1.NewConfig(key, secretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}
