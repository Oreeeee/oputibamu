package main

import "strings"

type Class struct {
	id        int
	name      string
	nameShort string
	url       string
}

func getClassNameShort(name string) string {
	return strings.Split(name, " ")[0]
}

func InitClass(id int, name string) Class {
	return Class{id, name, getClassNameShort(name), idToUrl("o", id)}
}

func InitClassFromHTML(htmlFile string, name string) Class {
	id := idFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Class{id, name, getClassNameShort(name), url}
}
