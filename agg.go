package main

import (
	"fmt"
	"time"
)

func Agg(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_requests>")
	}

	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every", duration)

	ticker := time.NewTicker(duration)

	for {
		if err := scrapeFeeds(s); err != nil {
			fmt.Println("error:", err)
		}
		<-ticker.C
	}
}
