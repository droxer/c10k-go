package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello from Echo server! Current time: " + time.Now().Format(time.RFC3339),
		})
	})

	e.GET("/echo", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name != "" {
			return c.String(http.StatusOK, "Hello, "+name+"!")
		}
		return c.String(http.StatusBadRequest, "Please provide a 'name' query parameter, e.g., /echo?name=John")
	})

	e.GET("/info", func(c echo.Context) error {
		userAgent := c.Request().UserAgent()
		remoteAddr := c.Request().RemoteAddr

		info := map[string]string{
			"server":      "Echo",
			"userAgent":   userAgent,
			"remoteAddr":  remoteAddr,
		}

		if c.Request().Method == http.MethodPost {
			// For POST requests, we could read the body length
			// but Echo makes it easier to work with the body directly
			info["method"] = "POST"
		}

		return c.JSON(http.StatusOK, info)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Start server on port 8083
	e.Logger.Info("Echo server starting on :8083")
	e.Logger.Fatal(e.Start(":8083"))
}