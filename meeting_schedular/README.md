# Meeting Scheduler

## Objective

The goal of this project is to develop a function that optimizes meeting room bookings by merging overlapping or consecutive booking times in a workspace management application.

## Function

### `optimizeBookings(bookings [][]int) [][]int`

- **Input**: A 2D slice `bookings` where each inner slice contains two integers representing the start and end times of a meeting.
- **Output**: A 2D slice of optimized meeting bookings, sorted by their starting times.

### Example

**Input**:
```go
[[9, 12], [11, 13], [14, 17], [16, 18]]
```

**Output**:
```go
[[9, 13], [14, 18]]
```

## How to Run

1. **Ensure Go is installed**: Make sure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

2. **Clone the repository**: Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/ekanshthakur15/Calry-Assignment.git
   ```
3. **Navigate to the project directory**: Change into the project directory
   ```bash
   cd meeting_schedular
   ```
5. **Run the application**: Execute the main Go file:
   ```go
   go run .
```
