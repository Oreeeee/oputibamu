package main

type Lesson struct {
	number int
	hours  string
	//data   string
	subject    string
	teacher    Teacher
	room       Room
	day        int
	group      int
	groupOutOf int
}

func InitEmptyLesson() Lesson {
	return Lesson{0, "", "", Teacher{}, Room{}, 0, 0, 0}
}
