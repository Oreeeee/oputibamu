package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

func getAllTableLessons(td *colly.HTMLElement, l Lesson) []Lesson {
	var ls []Lesson
	/*
		retNow := false

		les := l
	*/

	td.ForEach("span", func(i int, sp *colly.HTMLElement) {
		//if subject.Text == "&nbsp;" {
		//	// Empty lesson here, nothing much to do
		//	retNow = true
		//	return
		//}
		//les.subject = subject.Text
		//
		//td.ForEach("span .n", func(i int, lessonData *colly.HTMLElement) {
		//	l.teacher = lessonData.Text
		//})
		//
		//td.ForEach("span .s", func(i int, lessonData *colly.HTMLElement) {
		//	l.room = lessonData.Text
		//})
		//les := l
		//les.subject = sp.ChildText(".p")
		//les.teacher = sp.ChildText(".n")
		//les.room = sp.ChildText(".s")
		subject := sp.ChildText(".p")
		teacher := sp.ChildText(".n")
		room := sp.ChildText(".s")
		fmt.Println(subject, teacher, room)
		//ls = append(ls, les)
	})

	return ls

	//if retNow {
	//return ls
	//}

}

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
				// TODO: A lot
				//htmlData, _ := td.DOM.Html()
				//fmt.Printf("Lesson data: %v\n", htmlData)
				//l.data = htmlData
				fmt.Printf("Day: %d\n", i+1)
				if td.Text == "\xc2\xa0" {
					fmt.Println("NBSP")
					return
				}
				//fmt.Printf("%x\n", td.Text)
				//fmt.Println(td.Text)
				//fmt.Println(getAllTableLessons(td, l))
				getAllTableLessons(td, l)
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
