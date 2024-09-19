package oputibamu

import "fmt"

type Group struct {
	Group     int    `json:"group"`
	GroupMax  int    `json:"groupMax"`
	GroupName string `json:"groupName"`
}

func GroupNumberToName(group int) string {
	return fmt.Sprintf("gr%d", group)
}
