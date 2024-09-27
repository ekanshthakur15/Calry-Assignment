package main

// Struct to hold test case information
type TestCase struct {
	name     string
	bookings [][]int
}

// Function to generate TestCases
func GenerateTestCases() []TestCase {
	return []TestCase{
		{
			name:     "Given bookings",
			bookings: [][]int{{9, 12}, {11, 13}, {14, 17}, {16,18}},
		},
		{
			name:     "Non-overlapping bookings",
			bookings: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:     "Consecutive bookings",
			bookings: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name:     "Already non-overlapping bookings",
			bookings: [][]int{{1, 3}, {4, 6}, {7, 9}},
		},
		{
			name:     "Empty list of bookings",
			bookings: [][]int{},
		},
		{
			name:     "Overlapping bookings",
			bookings: [][]int{{1, 5}, {2, 3}, {4, 6}},
		},
		{
			name:     "Complex overlapping bookings",
			bookings: [][]int{{1, 10}, {2, 6}, {8, 12}, {15, 20}, {18, 22}},
		},

	}
}