package oputibamu

type Teacher struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func InitTeacher(id int, name string) Teacher {
	return Teacher{id, name, IdToUrl("n", id)}
}

func InitTeacherFromHTML(htmlFile string, name string) Teacher {
	id := IdFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Teacher{id, name, url}
}
