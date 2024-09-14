package main

// TODO: Replace the temporary types with actual types
type Lesson struct {
	number int
	hours  string
	//data   string
	subject    string
	teacher    string
	room       string
	day        int
	group      int
	groupOutOf int
}

func InitEmptyLesson() Lesson {
	return Lesson{0, "", "", "", "", 0, 0, 0}
}
