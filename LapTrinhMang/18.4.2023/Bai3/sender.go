package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename IP port\n", os.Args[0])
		os.Exit(1)
	}

	fileName := os.Args[1]
	serverIP := os.Args[2]
	serverPort := os.Args[3]

	// read file content
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "Error reading file: %s", err.Error())
		if err2 != nil {
			return
		}
		os.Exit(1)
	}

	// create UDP address
	serverAddr, err := net.ResolveUDPAddr("udp", serverIP+":"+serverPort)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error resolving address: %s", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}

	// create UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error connecting to server: %s", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Error closing connection: %s", err.Error())
			if err != nil {
				return
			}
		}
	}(conn)

	// send file name and content to server
	_, err = conn.Write([]byte(fileName + "\n" + string(fileContent)))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error sending data to server: %s", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
