package main

type Teacher struct {
	id   int
	name string
	url  string
}

func InitTeacher(id int, name string) Teacher {
	return Teacher{id, name, idToUrl("n", id)}
}

func InitTeacherFromHTML(htmlFile string, name string) Teacher {
	id := idFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Teacher{id, name, url}
}
