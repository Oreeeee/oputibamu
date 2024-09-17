package main

import (
	"context"
	"github.com/carlmjohnson/requests"
)

type ElektronikReplacement struct {
	LessonNumber string `json:"lesson"` // TODO: Make this an int
	TeacherName  string `json:"teacher"`
	Subject      string `json:"subject"`
	//Class        Class
	//Group        Group
	Classgroup []string `json:"classgroup"`
	RoomName   string   `json:"room"`
	Deputy     string   `json:"deputy"`
	Notes      string   `json:"notes"`
}

type ReplacementsResponse struct {
	Replacements []ElektronikReplacement `json:"rows"`
}

func (s *voScraper) getReplacementData() []ElektronikReplacement {
	var resp ReplacementsResponse
	err := requests.
		URL(s.elektronikApi + "/replacements.json").
		ToJSON(&resp).
		Fetch(context.Background())

	if err != nil {
		return resp.Replacements
	}

	return resp.Replacements
}
