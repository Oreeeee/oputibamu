package main

import (
	"github.com/gocolly/colly"
	"strconv"
)

type voScraper struct {
	timetableUrl string
}

func (s *voScraper) getClasses() []Class {
	var cA []Class
	c := colly.NewCollector()

	// TODO: Make the OnHTML code more reusable
	// Basically, class, teacher, and room stuff is the same but with different CSS attributes
	c.OnHTML("[name=\"oddzialy\"]", func(e *colly.HTMLElement) {
		e.ForEach("[value]", func(i int, element *colly.HTMLElement) {
			id, _ := strconv.Atoi(element.Attr("value"))
			cA = append(cA, InitClass(id, element.Text))
		})
	})

	err := c.Visit(s.timetableUrl + "/lista.html")
	if err != nil {
		return nil
	}

	return cA
}

func (s *voScraper) getRooms() []Room {
	var cA []Room
	c := colly.NewCollector()

	c.OnHTML("[name=\"sale\"]", func(e *colly.HTMLElement) {
		e.ForEach("[value]", func(i int, element *colly.HTMLElement) {
			id, _ := strconv.Atoi(element.Attr("value"))
			cA = append(cA, InitRoom(id, element.Text))
		})
	})

	err := c.Visit(s.timetableUrl + "/lista.html")
	if err != nil {
		return nil
	}

	return cA
}

func (s *voScraper) getTeachers() []Teacher {
	var cA []Teacher
	c := colly.NewCollector()

	c.OnHTML("[name=\"nauczyciele\"]", func(e *colly.HTMLElement) {
		e.ForEach("[value]", func(i int, element *colly.HTMLElement) {
			id, _ := strconv.Atoi(element.Attr("value"))
			cA = append(cA, InitTeacher(id, element.Text))
		})
	})

	err := c.Visit(s.timetableUrl + "/lista.html")
	if err != nil {
		return nil
	}

	return cA
}
