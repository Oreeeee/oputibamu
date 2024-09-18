package main

type Timetable struct {
	Class   Class    `json:"class"`
	Lessons []Lesson `json:"lessons"`
}
