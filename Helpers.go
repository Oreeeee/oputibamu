package oputibamu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IdFromHTML(htmlFile string) int {
	var result int

	// HACK: Go's regexp doesn't support lookahead assertions, we have to match
	// the number including .html and later remove that .html
	re := regexp.MustCompile(`(\d+)\.html`)
	tmp := re.FindString(htmlFile)
	tmp = strings.Split(tmp, ".")[0]

	result, _ = strconv.Atoi(tmp)
	return result
}

func IdToUrl(prefix string, id int) string {
	return fmt.Sprintf("/plany/%s%d.html", prefix, id)
}
