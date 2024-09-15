package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
	"strings"
)

func getLessonData(lessonElement *colly.HTMLElement, l *Lesson) {
	subjectRaw := lessonElement.ChildText(".p")

	// Grouping stuff
	groupRegex := regexp.MustCompile(`\b\d+/\d+\b`)
	groupMatches := groupRegex.FindAllString(subjectRaw, -1)
	if groupMatches == nil {
		// The lesson is not grouped, use the raw string
		l.subject = subjectRaw
		l.group = 0
		l.groupOutOf = 0
	} else {
		// The lesson is grouped
		groupData := strings.Split(groupMatches[0], "/")

		group, _ := strconv.Atoi(groupData[0])
		groupOutOf, _ := strconv.Atoi(groupData[1])

		l.group = group
		l.groupOutOf = groupOutOf

		// Set the subject without the group data
		l.subject = strings.Split(subjectRaw, "-")[0]
	}

	teacherName := lessonElement.ChildText(".n")
	teacherHTML := lessonElement.ChildAttr(".n", "href")

	roomName := lessonElement.ChildText(".s")
	roomHTML := lessonElement.ChildAttr(".s", "href")

	l.teacher = InitTeacherFromHTML(teacherHTML, teacherName)
	l.room = InitRoomFromHTML(roomHTML, roomName)
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
					getLessonData(sp, &l)
					m = append(m, l)
				})

				// Single group
				if !isMultipleGroups {
					getLessonData(td, &l)
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
