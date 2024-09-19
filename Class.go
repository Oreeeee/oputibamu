package oputibamu

import "strings"

type Class struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	NameShort string `json:"nameShort"`
	Url       string `json:"url"`
}

func GetClassNameShort(name string) string {
	return strings.Split(name, " ")[0]
}

func InitClass(id int, name string) Class {
	return Class{id, name, GetClassNameShort(name), IdToUrl("o", id)}
}

func InitClassFromHTML(htmlFile string, name string) Class {
	id := IdFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	return Class{id, name, GetClassNameShort(name), url}
}

func InitClassFromURL(url string, name string) Class {
	id := IdFromHTML(url)
	return Class{id, name, GetClassNameShort(name), url}
}
