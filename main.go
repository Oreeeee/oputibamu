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

	timetableDomain := os.Getenv("TIMETABLE_DOMAIN")
	timetableUrl := os.Getenv("TIMETABLE_URL")
	fmt.Println(timetableDomain, timetableUrl)

	s := voScraper{timetableDomain: timetableDomain, timetableUrl: timetableUrl}
	s.printSomeTimetable()

	cl := Class{1, "1a Example Class"}
	fmt.Println(cl.getUrl())

	room := Room{1, "A100PRz"}
	fmt.Println(room.getUrl())

	isPRz, przRoom := room.getIsPRz()
	if isPRz {
		fmt.Println(przRoom)
	}
}
