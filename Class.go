package main

type Class struct {
	id   int
	name string
	url  string
}

func InitClass(id int, name string) Class {
	return Class{id, name, idToUrl("o", id)}
}

func InitClassFromHTML(htmlFile string, name string) Class {
	id := idFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Class{id, name, url}
}
