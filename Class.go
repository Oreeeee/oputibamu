package main

import "fmt"

type Class struct {
	id   int
	name string
	url  string
}

func InitClass(id int, name string) Class {
	url := fmt.Sprintf("/plany/o%d.html", id)
	return Class{id, name, url}
}

func InitClassFromHTML(htmlFile string, name string) Class {
	id := idFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Class{id, name, url}
}
