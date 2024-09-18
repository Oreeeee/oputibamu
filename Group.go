package main

import "fmt"

type Group struct {
	Group     int    `json:"group"`
	GroupMax  int    `json:"groupMax"`
	GroupName string `json:"groupName"`
}

func groupNumberToName(group int) string {
	return fmt.Sprintf("gr%d", group)
}
