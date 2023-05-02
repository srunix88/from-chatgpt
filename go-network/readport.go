package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a TCP listener on a specific address and port
	address := "localhost:12345"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	fmt.Println("Listening on", address)

	// Accept incoming connections
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		fmt.Println("Accepted connection from", connection.RemoteAddr())

		// Read all data from the connection
		buffer := make([]byte, 1024)
		for {
			n, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading from connection:", err.Error())
				break
			}

			if n > 0 {
				// Print the received data to the console
				fmt.Print(string(buffer[:n]))
			}
		}

		// Close the connection
		connection.Close()
	}
}

