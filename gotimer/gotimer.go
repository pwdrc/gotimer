package gotimer

import (
	"flag"
	"fmt"
	"time"
)

func RunTimer() {
	help := flag.Bool("help", false, "Show these instructions")
	seconds := flag.Int("s", 0, "Number of seconds to run the timer")
	minutes := flag.Int("m", 0, "Number of minutes to run the timer")
	hours := flag.Int("H", 0, "Number of hours to run the timer")

	flag.Parse()

	totalTimeInSeconds := (*seconds) + (*minutes * 60) + (*hours * 3600)

	if *help {
		printHelp()
		return
	}

	if len(flag.Args()) > 0 {
		fmt.Println("Invalid arguments")
		printHelp()
		return
	}

	if totalTimeInSeconds == 0 {
		fmt.Println("You must specify a time")
		printHelp()
		return
	}

	fmt.Println("Starting timer")
	endTime := time.Now().Add(time.Duration(totalTimeInSeconds) * time.Second)

	fmt.Println("And... go!")

	for {
		remainingTime := int(endTime.Sub(time.Now()).Seconds())

		if remainingTime <= 0 {
			break
		}

		h := remainingTime / 3600
		m := (remainingTime % 3600) / 60
		s := remainingTime % 60

		fmt.Printf("\r%02d:%02d:%02d", h, m, s)
	}

	fmt.Println("\nAnd went!\a")
}

func printHelp() {
	fmt.Println("Usage: gotimer [-s seconds] [-m minutes] [-H hours]")
	fmt.Println("You can use any combination of the flags, but at least one is required.")
	fmt.Println("Examples:")
	fmt.Println("		gotimer -s 10")
	fmt.Println("		gotimer -m 5")
	fmt.Println("		gotimer -H 1")
	fmt.Println("		gotimer -H 1 -m 30")
}
