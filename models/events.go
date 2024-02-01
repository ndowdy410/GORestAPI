package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId 		int

}

var events = []Event{}


func (e Event)Save(){
	events = append(events, e)
}


func GetAllEvents(context *gin.Context) []Event {
	return events
}