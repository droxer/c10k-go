package main

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func handleConnection(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/plain; charset=utf8")

	path := ctx.Path()

	switch string(path) {
	case "/":

		fmt.Fprintf(ctx, "Hello from fasthttp server! Current time: %s", time.Now().Format(time.RFC3339))
	case "/echo":

		name := ctx.QueryArgs().Peek("name")
		if len(name) > 0 {

			ctx.WriteString("Hello, ")
			ctx.Write(name)
			ctx.WriteString("!")
		} else {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.WriteString("Please provide a 'name' query parameter, e.g., /echo?name=John")
		}
	case "/info":

		userAgent := ctx.UserAgent()
		remoteAddr := ctx.RemoteAddr().String()
		log.Println(userAgent, remoteAddr)

		ctx.WriteString("Server Info:\n")
		ctx.WriteString(fmt.Sprintf("User-Agent: %s\n", userAgent))
		ctx.WriteString(fmt.Sprintf("Remote Address: %s\n", remoteAddr))

		if ctx.IsPost() {
			body := ctx.PostBody()
			ctx.WriteString(fmt.Sprintf("Request Body Length: %d\n", len(body)))
		}

	case "/health":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.WriteString("OK")

	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.WriteString("404 Not Found")
	}

}

func main() {
	log.Println("Starting fasthttp server on :8081")

	// Create a fasthttp server instance
	s := &fasthttp.Server{
		Handler:            handleConnection,
		Concurrency:        256 * 1024,
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       5 * time.Second,
		IdleTimeout:        60 * time.Second,
		MaxRequestBodySize: 4 * 1024 * 1024, // 4MB
		DisableKeepalive:   false,
	}

	err := s.ListenAndServe(":8081")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
