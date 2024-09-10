package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Room struct {
	id   int
	name string
}

func (r *Room) getUrl() string {
	return fmt.Sprintf("/plany/s%d.html", r.id)
}

func (r *Room) getIsPRz() (bool, PRzRoom) {
	re, _ := regexp.Compile(`([A-Za-z])(\d{3})PRz`)
	match := re.FindStringSubmatch(r.name)

	building := match[1]
	room, _ := strconv.Atoi(match[2])

	if match != nil {
		return true, PRzRoom{building, room}
	} else {
		return false, PRzRoom{}
	}
}
