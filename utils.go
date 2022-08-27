package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchSentences(sentenceType string) SentenceList {
	var (
		reqUrl    string
		sentences SentenceList
	)

	switch sentenceType {
	case "quote":
		reqUrl = "https://raw.githubusercontent.com/monkeytypegame/monkeytype/master/frontend/static/quotes/english.json"
	}

	res, err := http.Get(reqUrl)
	if err != nil {
		fmt.Println("Failed to fetch the list of sentences")
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read the sentence body")
		os.Exit(1)
	}

	json.Unmarshal(body, &sentences)

	return sentences
}
