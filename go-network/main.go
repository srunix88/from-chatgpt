package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

		// Read data from stdin and send it to the client
		go func(conn net.Conn) {
			reader := bufio.NewReader(os.Stdin)
			for {
				// Read data from stdin
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err.Error())
					break
				}

				// Send data to the client
				_, err = conn.Write([]byte(input))
				if err != nil {
					fmt.Println("Error writing to connection:", err.Error())
					break
				}
			}

			// Close the connection
			conn.Close()
		}(connection)

		// Read data from the client and write it to stdout
		buffer := make([]byte, 1024)
		for {
			n, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading from connection:", err.Error())
				break
			}

			if n > 0 {
				// Write data to stdout
				fmt.Print(string(buffer[:n]))
			}
		}

		// Close the connection
		connection.Close()
	}
}

