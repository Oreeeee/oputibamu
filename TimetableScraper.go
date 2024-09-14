package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

func getLessonData(lessonElement *colly.HTMLElement) (string, string, string) {
	subject := lessonElement.ChildText(".p")
	teacher := lessonElement.ChildText(".n")
	room := lessonElement.ChildText(".s")
	return subject, teacher, room
}

func (s *voScraper) getRawTable() []Lesson {
	c := colly.NewCollector()
	var m []Lesson
	currentDay := 0

	c.OnHTML(".tabela", func(tabela *colly.HTMLElement) { // The main table
		fmt.Println("got tabela")
		tabela.ForEach("tr", func(i int, tr *colly.HTMLElement) { // Every row in the table
			fmt.Println("got tr")
			if i == 0 {
				// We skip the table header stuff, we don't need it
				fmt.Println("skipping...")
				return
			}

			l := InitEmptyLesson()

			// DO NOT create a new lesson here
			tr.ForEach("td .nr", func(i int, td *colly.HTMLElement) {
				// Lesson number
				lN, _ := strconv.Atoi(td.Text)
				fmt.Printf("Lesson: %v\n", td.Text)
				l.number = lN
			})

			// DO NOT create a new lesson here
			tr.ForEach("td .g", func(i int, td *colly.HTMLElement) {
				// Hours of the lesson
				fmt.Printf("Hours: %v\n", td.Text)
				l.hours = td.Text
			})

			// Here... it gets... complicated...
			tr.ForEach("td .l", func(i int, td *colly.HTMLElement) {
				// Lesson data field

				l.day = currentDay
				currentDay++

				if td.Text == "\xc2\xa0" {
					fmt.Println("NBSP")
					return
				}

				isMultipleGroups := false

				// Multiple groups
				td.ForEach("[style]", func(i int, sp *colly.HTMLElement) {
					isMultipleGroups = true

					subject, teacher, room := getLessonData(sp)
					l.subject = subject
					l.teacher = teacher
					l.room = room

					m = append(m, l)
				})

				// Single group
				if !isMultipleGroups {
					subject, teacher, room := getLessonData(td)
					l.subject = subject
					l.teacher = teacher
					l.room = room

					m = append(m, l)
				}

			})
			currentDay = 0
		})
	})

	err := c.Visit(s.timetableUrl + "/plany/o11.html")
	if err != nil {
		return m
	}
	return m
}
