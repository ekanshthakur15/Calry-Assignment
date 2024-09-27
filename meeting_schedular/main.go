package main

import (
	"fmt"

	scheduler "github.com/ekanshthakur15/meeting-scheduler/scheduler"
)


func main() {
	// Input 2D array representing meeting intervals
	testcases := GenerateTestCases()

	// To Get the Optimised Booking
	for _, tc := range testcases {
		fmt.Printf("Test Case: %s\n", tc.name)
		fmt.Printf("Input Booking: \n")
		fmt.Println(tc.bookings)
		output := scheduler.OptimizeBookings(tc.bookings)
		fmt.Printf("Output: \n")
		fmt.Println(output)
		fmt.Println()
	}
}