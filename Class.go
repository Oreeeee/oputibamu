package main

import "fmt"

type Class struct {
	id   int
	name string
}

func (c *Class) getUrl() string {
	return fmt.Sprintf("/plany/o%d.html", c.id)
}
