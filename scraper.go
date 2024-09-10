package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type voScraper struct {
	timetableDomain, timetableUrl string
}

func (s *voScraper) printSomeTimetable() {
	c := colly.NewCollector(
		colly.AllowedDomains(s.timetableDomain))

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", USER_AGENT)
		fmt.Println(r.Headers)
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	err := c.Visit(s.timetableUrl + "/plany/o1.html")
	if err != nil {
		return
	}
}
