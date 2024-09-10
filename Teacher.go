package main

import "fmt"

type Teacher struct {
	id   int
	name string
}

func (t *Teacher) getUrl() string {
	return fmt.Sprintf("/plany/n%d.html", t.id)
}
