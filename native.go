package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Serving new connection from %s", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection from %s closed by client.", conn.RemoteAddr())
			} else {
				log.Printf("Error reading from %s: %v", conn.RemoteAddr(), err)
			}
			return
		}

		request := string(buffer[:n])
		log.Printf("Received from %s: %s", conn.RemoteAddr(), request)

		response := fmt.Sprintf("Echo: %s", request)
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("Error writing to %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}
