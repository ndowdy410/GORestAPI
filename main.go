package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ndowdy410/GORestAPI.git/models"
)


func main(){
	fmt.Println("Starting Go Rest API project 1")

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run("localhost:8080")
}

func getEvents(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{"message": "Hello from Go Rest API project 1 !"})
}

func createEvent(context *gin.Context){

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserId = 1

	context.JSON(http.StatusBadRequest, gin.H{"message":"Event Created.", "event":event})



}