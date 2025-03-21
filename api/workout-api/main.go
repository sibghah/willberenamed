package main

workouts := SeedWorkoutData()
//TODO: what is gin
func getWorkoutComponents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workouts)
}

