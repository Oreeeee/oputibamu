package main

type Teacher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func InitTeacher(id int, name string) Teacher {
	return Teacher{id, name, idToUrl("n", id)}
}

func InitTeacherFromHTML(htmlFile string, name string) Teacher {
	id := idFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Teacher{id, name, url}
}
