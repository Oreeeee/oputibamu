package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Room struct {
	id    int
	name  string
	url   string
	isPRz bool
	prz   PRzRoom
}

func InitRoom(id int, name string) Room {
	url := fmt.Sprintf("/plany/s%d.html", id)
	isPRz, prz := getIsPRz(name)
	return Room{id, name, url, isPRz, prz}
}

func InitRoomFromHTML(htmlFile string, name string) Room {
	id := idFromHTML(htmlFile)
	url := "/plany/" + htmlFile
	isPRz, prz := getIsPRz(name)
	return Room{id, name, url, isPRz, prz}
}

func getIsPRz(name string) (bool, PRzRoom) {
	re, _ := regexp.Compile(`([A-Za-z])(\d{3})PRz`)
	match := re.FindStringSubmatch(name)

	if match == nil {
		return false, PRzRoom{}
	}

	building := match[1]
	room, _ := strconv.Atoi(match[2])

	return true, PRzRoom{building, room}
}
