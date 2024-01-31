package main


import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main(){
	fmt.Println("Starting Go Rest API project 1")

	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run("localhost:8080")
}

func getEvents(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{"message": "Hello from Go Rest API project 1 !"})
}