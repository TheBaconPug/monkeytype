package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	sentences, err := os.Create("data/sentences.json")
	if err != nil {
		fmt.Println("Failed to create sentences.json")
		os.Exit(1)
	}
	fileContent, _ := json.MarshalIndent(fetchSentences("quote"), "", " ")
	sentences.Write(fileContent)
	sentences.Close()
}
