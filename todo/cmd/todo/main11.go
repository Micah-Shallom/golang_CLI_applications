package main

import (
	"fmt"
	"time"
)

func timings() {
	currentTime := time.Now()

	// Print the current time
	fmt.Println("Current Time:", currentTime)

	// Extract and print the date components
	year, month, day := currentTime.Date()
	fmt.Printf("Date: %d-%02d-%02d\n", year, month, day)

	// Extract and print the time components
	hour, minute, second := currentTime.Clock()
	fmt.Printf("Time: %02d:%02d:%02d\n", hour, minute, second)
}
