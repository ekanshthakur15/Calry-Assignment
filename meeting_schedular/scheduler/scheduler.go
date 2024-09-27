package scheduler

/*
	This file contains the logic of optimising the booking and club booking that are
	overlapping.

	Time Complexity : T(n) = O(NlogN)
	Space Complexity : T(n) = O(N)
*/

import (
	"sort"
)


func OptimizeBookings(bookings [][]int) [][]int {


	/*
	Args: 2D Array number[][] -> Each element represents start and end time 
		of the meetings.

	Return: 2B Array number[][] -> Optimised meeting schedule which combines
		overlapping meetings into one.
	*/



	// Sort the bookings based on their starting time
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i][0] < bookings[j][0]
	})
	// fmt.Println(bookings)

	// To store the optimised bookings
	merged := [][]int{}

	// Logic for combining overlapping bookings

	for _, booking := range bookings {

		if len(merged) ==  0 || merged[len(merged)-1][1] <= booking[0] - 1 {
			merged = append(merged, booking)
		} else if booking[0] <= merged[len(merged)-1][1]  {
			merged[len(merged)-1][1] = max(booking[1], merged[len(merged)-1][1])
		}
	}
	return merged

}