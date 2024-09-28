# Hotel Room Service Request API

## Objective

The objective of this project is to develop a comprehensive set of RESTful APIs designed for managing and prioritizing hotel room service requests. This backend system is imlemented using **Golang** and **Gin-Gonic** and utilizes JSON files for temporary data storage, facilitating the efficient handling of requests based on urgency and guest status. The system aims to help hotel staff manage room service operations effectively, ensuring that urgent and high-priority requests are serviced promptly to enhance guest satisfaction and operational efficiency.

## Project Description

### API Design

- **Service Request Fields**:
  - `id`: string
  - `guestName`: string
  - `roomNumber`: number
  - `requestDetails`: string
  - `priority`: number (lower numbers indicate higher priority)
  - `status`: one of `'received'`, `'in progress'`, `'awaiting confirmation'`, `'completed'`, or `'canceled'`

- **API Endpoints**:
  - **POST** `/requests`: Add a new service request.
  - **GET** `/requests`: Retrieve all requests, sorted by priority.
  - **GET** `/requests/{id}`: Retrieve a specific request by its ID.
  - **PUT** `/requests/{id}`: Update details or the priority of an existing request.
  - **DELETE** `/requests/{id}`: Remove a completed or canceled request.
  - **POST** `/requests/{id}/complete`: Mark a request as completed.

### JSON Data Storage Handling

- Created JSON file (request.json) to store room service request data temporarily.
- Implemented a locking mechanism to manage access to the JSON file.


## How to Run

1. **Ensure Go is installed**: Make sure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

2. **Clone the repository**: Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/ekanshthakur15/Calry-Assignment.git
   ```
3. **Navigate to the project directory**: Change into the project directory
   ```bash
   cd hotel_room_service
   ```
4. **Install Dependencies**:
   ```go
   go mod tidy
   ```
5. **Run the application**: This command will start the server at port **3000**
   ```go
   go run cmd/main.go
   ```
  
