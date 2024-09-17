package main

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
)

type ElektronikReplacement struct {
	lessonNumber int
	class        Class
	group        Group
	room         Room
}

type ReplacementsResponse struct {
	Replacements []map[string]string `json:"rows"`
}

func (s *voScraper) getReplacementData() {
	// TODO: fix
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
