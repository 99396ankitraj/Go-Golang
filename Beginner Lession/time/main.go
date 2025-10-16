package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	// Formatting time
	formattedTime := currentTime.Format("2006/01/02 Monday")
	fmt.Println("Formatted Time:", formattedTime)

	// Add 1 more day
	nextDay := currentTime.Add(24 * time.Hour)
	fmt.Println("Next Day:", nextDay.Format("2006/01/02 Monday"))
}