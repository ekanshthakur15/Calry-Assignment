package main

import (
	"github.com/ekanshthakur15/hotel-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/requests", controllers.CreateRequest)
	r.GET("/requests", controllers.GetAllRequests)
	r.GET("/requests/:id", controllers.GetRequestByID)
	r.PUT("/requests/:id", controllers.UpdateRequest)
	r.DELETE("/requests/:id", controllers.DeleteRequest)
	r.POST("/requests/:id/complete", controllers.CompleteRequest)

	// Start the server and listen on port 8080 (default)
	if err := r.Run(":3000"); err != nil {
		panic(err) // Handle error if the server fails to start
	}
}
