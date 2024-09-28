package controllers

import (
	"sort"

	"github.com/ekanshthakur15/hotel-service/internal/models"
	"github.com/ekanshthakur15/hotel-service/internal/storage"
	"github.com/ekanshthakur15/hotel-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateRequest(c *gin.Context) {
	var newRequest models.RoomServiceRequest
	if err := c.ShouldBindJSON(&newRequest); err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
		return
	}

	requests , err := storage.ReadFromFile()
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
		return
	}

	newRequest.ID = utils.GenerateID()
	newRequest.Status = "recieved"

	requests = append(requests, newRequest)

	if err := storage.WriteToFile(requests); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
		return
	}

	c.JSON(201, newRequest)
}

func GetAllRequests(c *gin.Context) {

	requests, err := storage.ReadFromFile()
	if err != nil {
		c.JSON(500, gin.H{"errors":err.Error()})
		return
	}

	sort.Slice(requests, func(i, j int) bool {
		// If one of the requests is completed, move it to the end
		if requests[i].Status == "completed" && requests[j].Status != "completed" {
			return false
		}
		if requests[i].Status != "completed" && requests[j].Status == "completed" {
			return true
		}
		// Otherwise, sort by priority (lower numbers have higher priority)
		return requests[i].Priority < requests[j].Priority
	})

	c.JSON(200, requests)
}

func GetRequestByID(c *gin.Context){
	id := c.Param("id")
	requests, err := storage.ReadFromFile()
	if err != nil{
		c.JSON(500, gin.H{"errors":err.Error()})
		return
	}
	for _, req := range requests {
		if req.ID == id {
			c.JSON(200, req)
			return
		}
	}
	c.JSON(404, gin.H{"errors": "Request not found"})
}

func UpdateRequest(c *gin.Context) {

	id := c.Param("id")
	var updatedRequest models.RoomServiceRequest

	if err := c.ShouldBindJSON(&updatedRequest); err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
		return
	}

	requests, err := storage.ReadFromFile()
	if err != nil{
		c.JSON(500, gin.H{"errors":err.Error()})
		return
	}
	for i, req := range requests {
		if req.ID == id {
			updatedRequest.ID = id
			requests[i] = updatedRequest

			if err := storage.WriteToFile(requests); err != nil {
				c.JSON(500, gin.H{"error":"Unable to write date"})
				return
			}
			c.JSON(200, updatedRequest)
			return
		}
	}
	c.JSON(404, gin.H{"errors": "Request not found"})
}

func DeleteRequest(c *gin.Context) {

	id := c.Param("id")

	requests, err := storage.ReadFromFile()
	if err != nil{
		c.JSON(500, gin.H{"errors":err.Error()})
		return
	}

	for i, req := range requests {
		if req.ID == id {
			requests = append(requests[:i],requests[i+1:]... )

			if err := storage.WriteToFile(requests); err != nil {
				c.JSON(500, gin.H{"error":"Unable to write date"})
				return
			}
			c.JSON(200, gin.H{"message":"Request deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"errors": "Request not found"})
}

func CompleteRequest(c *gin.Context) {

	id := c.Param("id")

	requests, err := storage.ReadFromFile()
	if err != nil{
		c.JSON(500, gin.H{"errors":err.Error()})
		return
	}
	for i, req := range requests {
		if req.ID == id {
			req.Status = "completed"
			requests[i] = req

			if err := storage.WriteToFile(requests); err != nil {
				c.JSON(500, gin.H{"error":"Unable to write date"})
				return
			}
			c.JSON(200, req)
			return
		}
	}
	c.JSON(404, gin.H{"errors": "Request not found"})
}