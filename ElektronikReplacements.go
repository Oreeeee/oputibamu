package main

import (
	"context"
	"github.com/carlmjohnson/requests"
	"regexp"
	"strconv"
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
	Date         string                  `json:"date"`
	Replacements []ElektronikReplacement `json:"rows"`
}

type ReplacementsData struct {
	day          int
	replacements []ElektronikReplacement
}

func (s *voScraper) fetchReplacementData() ReplacementsResponse {
	var resp ReplacementsResponse
	err := requests.
		URL(s.elektronikApi + "/replacements.json").
		ToJSON(&resp).
		Fetch(context.Background())

	if err != nil {
		return resp
	}

	return resp
}

func (s *voScraper) getReplacementData() ReplacementsData {
	res := s.fetchReplacementData()
	re := regexp.MustCompile(`\b(pon|wt|sr|czw|pt)\b`)
	day := Days[re.FindString(res.Date)]
	return ReplacementsData{day, res.Replacements}
}

func (r *ReplacementsData) getCurrentLessonReplacements(day int, l Lesson, c Class, g Group) ElektronikReplacement {
	if r.day != day {
		// No data for this day
		return ElektronikReplacement{}
	}
	for _, rep := range r.replacements {
		lN, _ := strconv.Atoi(rep.LessonNumber)
		if lN == l.Number && rep.Classgroup[0] == c.NameShort && rep.Classgroup[1] == l.Group.GroupName {
			return rep
		}
	}
	return ElektronikReplacement{}
}
