package data

import (
	"fmt"
	"github.com/cdipaolo/sentiment"
)

func Analyze(text string){
	model, err := sentiment.Restore()
	if err != nil {
		fmt.Println(err)
	}

	analysis := model.SentimentAnalysis(text, sentiment.English)
	fmt.Println(analysis.Score)

}
