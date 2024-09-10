package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

type voScraper struct {
	timetableDomain, timetableUrl string
}

func (s *voScraper) printSomeTimetable() {
	c := colly.NewCollector(
		colly.AllowedDomains(s.timetableDomain))

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	err := c.Visit(s.timetableUrl + "/plany/o1.html")
	if err != nil {
		return
	}
}

func (s *voScraper) getClasses() []Class {
	var cA []Class
	c := colly.NewCollector()

	// TODO: Make the OnHTML code more reusable
	// Basically, class, teacher, and room stuff is the same but with different CSS attributes
	c.OnHTML("[name=\"oddzialy\"]", func(e *colly.HTMLElement) {
		e.ForEach("[value]", func(i int, element *colly.HTMLElement) {
			id, _ := strconv.Atoi(element.Attr("value"))
			cA = append(cA, Class{id, element.Text})
		})
	})

	err := c.Visit(s.timetableUrl + "/lista.html")
	if err != nil {
		return nil
	}

	return cA
}
