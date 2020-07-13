package tw

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func Auth() *twitter.Client {
	accessToken:="1130660432681799681-gKk2Qf1LJpuWGGBDhunnCM5TsBQPB9"
	accessTokenSecret:="7GNSMAnva0V78xCpqdJ433H07zeMbkFERdrYmbeomleH1"
	key:="2zr105kPhmM6qxmkqHeYsl6mC"
	secretKey:="7MjZeewsCjLaAZriEr68bJZaFEARmYO1OdrPznXY9sypHWTsVV"


	config := oauth1.NewConfig(key, secretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}
