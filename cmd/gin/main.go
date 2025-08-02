package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: "1", Name: "Alice", Age: 30},
	{ID: "2", Name: "Bob", Age: 24},
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin example API!",
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")
		minAgeStr := c.Query("min_age")

		filteredUsers := []User{}
		for _, user := range users {
			if name != "" && user.Name != name {
				continue
			}

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

	router.POST("/users", func(c *gin.Context) {
		var newUser User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = fmt.Sprintf("%d", len(users)+1)
		users = append(users, newUser)

		c.JSON(http.StatusCreated, newUser)
	})

	router.Use(func(c *gin.Context) {
		fmt.Println("Before request:", c.Request.URL.Path)
		c.Next()
		fmt.Println("After request:", c.Request.URL.Path, "Status:", c.Writer.Status())
	})

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This route passed through custom middleware!"})
	})

	log.Println("Gin server starting on :8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
