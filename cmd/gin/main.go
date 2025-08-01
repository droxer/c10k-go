package main

import (
	"fmt"
	"log"
	"net/http" // Used for HTTP status codes like http.StatusOK
	"strconv"  // For converting string to int

	"github.com/gin-gonic/gin"
)

// User struct to demonstrate JSON binding and response
type User struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"` // 'required' validation tag
	Age  int    `json:"age"`
}

// In-memory "database" for demonstration purposes
var users = []User{
	{ID: "1", Name: "Alice", Age: 30},
	{ID: "2", Name: "Bob", Age: 24},
}

func main() {
	// Create a Gin router with default middleware (Logger and Recovery)
	router := gin.Default()

	// Set the server to release mode in production for better performance and less logging
	// gin.SetMode(gin.ReleaseMode)

	// --- Routes ---

	// 1. Basic GET request
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin example API!",
		})
	})

	// 2. GET request with a path parameter
	// Access: /users/1, /users/Alice
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id") // Get the 'id' parameter from the URL

		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user) // Return the user as JSON
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// 3. GET request with query parameters
	// Access: /search?name=Alice&min_age=25
	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")         // Get 'name' query parameter
		minAgeStr := c.Query("min_age") // Get 'min_age' query parameter as string

		filteredUsers := []User{}
		for _, user := range users {
			// Filter by name if provided
			if name != "" && user.Name != name {
				continue
			}

			// Filter by minimum age if provided
			if minAgeStr != "" {
				minAge, err := strconv.Atoi(minAgeStr)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_age parameter"})
					return
				}
				if user.Age < minAge {
					continue
				}
			}
			filteredUsers = append(filteredUsers, user)
		}

		if len(filteredUsers) > 0 {
			c.JSON(http.StatusOK, filteredUsers)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "No users found matching criteria"})
		}
	})

	// 4. POST request - handling JSON body and validation
	router.POST("/users", func(c *gin.Context) {
		var newUser User

		// Bind the JSON request body to the newUser struct
		// Gin automatically validates based on 'binding:"required"' tags
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simulate assigning a new ID (in a real app, this would be from a database)
		newUser.ID = fmt.Sprintf("%d", len(users)+1)
		users = append(users, newUser)

		c.JSON(http.StatusCreated, newUser) // Return the created user with 201 status
	})

	// 5. Route Grouping (for API versioning or common middleware)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "1.0", "author": "Go Gin Example"})
		})
		// More v1 routes...
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "2.0", "status": "beta"})
		})
		// More v2 routes...
	}

	// 6. Custom Middleware Example
	router.Use(func(c *gin.Context) {
		// This middleware will run BEFORE each request
		fmt.Println("Before request:", c.Request.URL.Path)
		c.Next() // Continue to the next middleware or handler
		// This middleware will run AFTER the request has been handled
		fmt.Println("After request:", c.Request.URL.Path, "Status:", c.Writer.Status())
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This route passed through custom middleware!"})
	})

	// Run the server on port 8080
	// This will block until the server is stopped (e.g., Ctrl+C)
	log.Println("Gin server starting on :8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
