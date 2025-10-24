package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// A type that defines the user struct 
// Defines ID, FirstName, and LastName
type user struct{
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

// Initial User's List 
var users = []user{
	{ID: "1", FirstName:"John", LastName:"Smith"},
	{ID: "2", FirstName:"Jane", LastName:"Doe"},
}


// Function to Get ALL users
func getUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, users)
}

// Function to Get a SPECIFIC user
func getUserById(c *gin.Context){
	id := c.Param("id")

	// SLOW look up if large data set, but for simple test we can do this.
	// Find the user and return it with 200 status
	for _, a := range users{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// If user was not found return 404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User was not found."})
}



// Function to Create a new user
func postUsers(c *gin.Context){
	var newUser user

	if err := c.BindJSON(&newUser); err !=nil{
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// Main function, defining the endpoints and methods for the router
func main(){
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", postUsers)
	router.Run("localhost:8080")
}


