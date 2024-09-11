package main

import "fmt"

type Teacher struct {
	id   int
	name string
	url  string
}

func InitTeacher(id int, name string) Teacher {
	url := fmt.Sprintf("/plany/n%d.html", id)
	return Teacher{id, name, url}
}
