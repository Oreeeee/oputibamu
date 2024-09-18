package main

type Lesson struct {
	Number      int                   `json:"number"`
	Hours       string                `json:"hours"`
	Subject     string                `json:"subject"`
	Teacher     Teacher               `json:"teacher"`
	Room        Room                  `json:"room"`
	Day         int                   `json:"day"`
	Group       Group                 `json:"group"`
	Replacement ElektronikReplacement `json:"replacement"`
}

func InitEmptyLesson() Lesson {
	return Lesson{0, "", "", Teacher{}, Room{}, 0, Group{}, ElektronikReplacement{}}
}
