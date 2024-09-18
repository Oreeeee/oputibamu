package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
	"strings"
)

func getLessonData(lessonElement *colly.HTMLElement, l *Lesson, c Class, rep ReplacementsData) {
	subjectRaw := lessonElement.ChildText(".p")

	// Grouping stuff
	groupRegex := regexp.MustCompile(`\b\d+/\d+\b`)
	groupMatches := groupRegex.FindAllString(subjectRaw, -1)
	if groupMatches == nil {
		// The lesson is not grouped, use the raw string
		l.Subject = subjectRaw
	} else {
		// The lesson is grouped
		groupData := strings.Split(groupMatches[0], "/")

		group, _ := strconv.Atoi(groupData[0])
		groupMax, _ := strconv.Atoi(groupData[1])

		l.Group.Group = group
		l.Group.GroupMax = groupMax
		l.Group.GroupName = groupNumberToName(group)

		// Set the subject without the group data
		l.Subject = strings.Split(subjectRaw, "-")[0]
	}

	teacherName := lessonElement.ChildText(".n")
	teacherHTML := lessonElement.ChildAttr(".n", "href")

	roomName := lessonElement.ChildText(".s")
	roomHTML := lessonElement.ChildAttr(".s", "href")

	l.Replacement = rep.getCurrentLessonReplacements(l.Day, *l, c, Group{})
	l.Teacher = InitTeacherFromHTML(teacherHTML, teacherName)
	l.Room = InitRoomFromHTML(roomHTML, roomName)
}

func (s *voScraper) getRawTable(url string) Timetable {
	c := colly.NewCollector()
	timetable := Timetable{}
	currentDay := 0
	replacements := s.getReplacementData()

	c.OnHTML(".tytulnapis", func(title *colly.HTMLElement) { // Gets the class name
		timetable.Class = InitClassFromURL(url, title.Text)
	})

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
				l.Number = lN
			})

			// DO NOT create a new lesson here
			tr.ForEach("td .g", func(i int, td *colly.HTMLElement) {
				// Hours of the lesson
				fmt.Printf("Hours: %v\n", td.Text)
				l.Hours = td.Text
			})

			// Here... it gets... complicated...
			tr.ForEach("td .l", func(i int, td *colly.HTMLElement) {
				// Lesson data field

				l.Day = currentDay
				currentDay++

				if td.Text == "\xc2\xa0" {
					fmt.Println("NBSP")
					return
				}

				isMultipleGroups := false

				// Multiple groups
				td.ForEach("[style]", func(i int, sp *colly.HTMLElement) {
					isMultipleGroups = true
					getLessonData(sp, &l, timetable.Class, replacements)
					timetable.Lessons = append(timetable.Lessons, l)
				})

				// Single group
				if !isMultipleGroups {
					getLessonData(td, &l, timetable.Class, replacements)
					timetable.Lessons = append(timetable.Lessons, l)
				}

			})
			currentDay = 0
		})
	})

	err := c.Visit(s.timetableUrl + "/plany/o11.html")
	if err != nil {
		return timetable
	}
	return timetable
}
