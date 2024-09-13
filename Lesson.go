package main

// TODO: Replace the temporary types with actual types
type Lesson struct {
	number int
	hours  string
	//data   string
	subject string
	teacher string
	room    string
}

func InitEmptyLesson() Lesson {
	return Lesson{0, "", "", "", ""}
}
