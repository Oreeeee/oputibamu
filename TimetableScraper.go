package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

func (s *voScraper) getRawTable() []Lesson {
	c := colly.NewCollector()
	var m []Lesson

	c.OnHTML(".tabela", func(tabela *colly.HTMLElement) { // The main table
		fmt.Println("got tabela")
		tabela.ForEach("tr", func(i int, tr *colly.HTMLElement) { // Every row in the table
			fmt.Println("got tr")
			if i == 0 {
				// We skip the table header stuff, we don't need it
				fmt.Println("skipping...")
				return
			}

			l := Lesson{0, "", ""}

			tr.ForEach("td .nr", func(i int, td *colly.HTMLElement) {
				// Lesson number
				lN, _ := strconv.Atoi(td.Text)
				fmt.Printf("Lesson: %v\n", td.Text)
				l.number = lN
			})

			tr.ForEach("td .g", func(i int, td *colly.HTMLElement) {
				// Hours of the lesson
				fmt.Printf("Hours: %v\n", td.Text)
				l.hours = td.Text
			})

			tr.ForEach("td .l", func(i int, td *colly.HTMLElement) {
				// Lesson data field
				// TODO: A lot
				htmlData, _ := td.DOM.Html()
				fmt.Printf("Lesson data: %v\n", htmlData)
				l.data = htmlData
			})

			m = append(m, l)
		})
	})

	err := c.Visit(s.timetableUrl + "/plany/o11.html")
	if err != nil {
		return m
	}
	return m
}
