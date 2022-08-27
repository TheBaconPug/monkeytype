package main

type SentenceList struct {
	Sentences []struct {
		Text       string `json:"text"`
		Source     string `json:"source"`
		Length     int    `json:"length"`
		ID         int    `json:"id"`
		ApprovedBy string `json:"approvedBy,omitempty"`
	} `json:"quotes"`
}
