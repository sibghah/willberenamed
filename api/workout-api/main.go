package main

//TODO: this won't be a main file later

import (
	. "awesomeProject/api/workout-api/seed"
	"net/http"

	"github.com/gin-gonic/gin"
)

var workouts = SeedWorkoutData()

// gin is a web framework (a bit like express or flask) for Go, used to build APIs and web services
//** Syntax breakdown **
//  getWorkoutComponents is a user-defined HTTP handler function
//  c *gin.Context represents the request and response context
// **//

func getWorkoutComponents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workouts)
}

func main() {

	//creates a new Gin engine (router) with some default middleware
	r := gin.Default()

	r.GET("/workouts", getWorkoutComponents)

	r.Run(":8080")
}
