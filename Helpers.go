package main

import (
	"fmt"
	"strconv"
	"strings"
)

func idFromHTML(htmlFile string) int {
	var result int
	var tmp string

	tmp = htmlFile[1:]               // Remove the first letter
	tmp = strings.Split(tmp, ".")[0] // Remove the .html
	result, _ = strconv.Atoi(tmp)

	return result
}

func idToUrl(prefix string, id int) string {
	return fmt.Sprintf("/plany/%s%d.html", prefix, id)
}
