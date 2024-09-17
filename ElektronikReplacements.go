package main

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
)

// TODO: fix the types, they are only temporary (for the most part)
type ElektronikReplacement struct {
	LessonNumber string `json:"lesson"`
	Teacher      string `json:"teacher"`
	Subject      string `json:"subject"`
	//Class        Class
	//Group        Group
	Classgroup []string `json:"classgroup"`
	Room       string   `json:"room"`
	Deputy     string   `json:"deputy"`
	Notes      string   `json:"notes"`
}

type ReplacementsResponse struct {
	Replacements []ElektronikReplacement `json:"rows"`
}

func (s *voScraper) getReplacementData() {
	// TODO: fix
	//var resp map[string]interface{}
	var resp ReplacementsResponse
	err := requests.
		URL(s.elektronikApi + "/replacements.json").
		ToJSON(&resp).
		Fetch(context.Background())

	if err != nil {
		fmt.Println("Error fetching replacement data:", err)
		return
	}

	fmt.Println(resp)
}
