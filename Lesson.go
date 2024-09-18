package main

type Lesson struct {
	number      int
	hours       string
	subject     string
	teacher     Teacher
	room        Room
	day         int
	group       Group
	replacement ElektronikReplacement
}

func InitEmptyLesson() Lesson {
	return Lesson{0, "", "", Teacher{}, Room{}, 0, Group{}, ElektronikReplacement{}}
}
