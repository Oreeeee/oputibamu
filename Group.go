package main

import "fmt"

type Group struct {
	group     int
	groupMax  int
	groupName string
}

func groupNumberToName(group int) string {
	return fmt.Sprintf("gr%d", group)
}
