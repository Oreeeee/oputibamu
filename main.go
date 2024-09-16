package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"slices"
)

func main() {
	// This is the application part of the library
	// Its main use is testing for now, until it works as a library/API
	// Also might be used in daemon mode

	// Daemon mode placeholder
	if slices.Contains(os.Args, "-d") {
		fmt.Println("Daemon mode!!!!!")
		return
	}

	// Load the .env
	loadEnv := godotenv.Load()
	if loadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	timetableUrl := os.Getenv("TIMETABLE_URL")
	elektronikApi := os.Getenv("ELEKTRONIK_API")
	var elektronikMode bool
	fmt.Println(timetableUrl)

	if elektronikApi == "0" { // Elektronik API disabled
		elektronikMode = false
	} else {
		elektronikMode = true
	}

	s := voScraper{timetableUrl, elektronikMode, elektronikApi}

	//cl := InitClass(1, "1a Example Class")
	//fmt.Println(cl.url)
	//
	//room := InitRoom(1, "A100PRz")
	//fmt.Println(room.url)
	//
	//if room.isPRz {
	//	fmt.Println(room.prz)
	//}
	//
	//fmt.Println(s.getClasses())
	//fmt.Println(s.getRooms())
	//fmt.Println(s.getTeachers())
	fmt.Println(s.getRawTable("/plany/o11.html"))
	//s.getRawTable()
}
