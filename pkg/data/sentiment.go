package data

import (
	"fmt"
	"github.com/cdipaolo/sentiment"
)

func Analyze(text string) bool {
	model, err := sentiment.Restore()
	if err != nil {
		fmt.Println(err)
	}

	analysis := model.SentimentAnalysis(text, sentiment.English)
	if analysis.Score == 1 {
		return false
	} else {
		return true
	}
}

