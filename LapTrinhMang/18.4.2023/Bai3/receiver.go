package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// parse command line arguments
	if len(os.Args) != 2 {
		_, err := fmt.Fprintf(os.Stderr, "Usage: %s port\n", os.Args[0])
		if err != nil {
			return
		}
		os.Exit(1)
	}
	port := os.Args[1]

	// create UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "error resolving UDP address: %s", err.Error())
		if err2 != nil {
			return
		}
		os.Exit(1)
	}

	// listen for UDP packets
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "error listening for UDP : %s", err.Error())
		if err2 != nil {
			return
		}
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on port %s...\n", port)

	// receive and handle UDP packets
	for {
		buffer := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Error receiving : %s", err.Error())
			if err != nil {
				return
			}
			continue
		}

		// parse file name and content from UDP packet
		packet := string(buffer[:n])
		parts := strings.Split(packet, "\n")
		if len(parts) != 2 {
			_, err := fmt.Fprintf(os.Stderr, "Invalid packet: %s\n", packet)
			if err != nil {
				return
			}
			continue
		}
		fileName := parts[0]
		fileContent := []byte(parts[1])

		// write file content to disk
		err = os.WriteFile(fileName, fileContent, 0644)
		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "Error writing file: %s", err.Error())
			if err2 != nil {
				return
			}
			continue
		}

		// success message
		fmt.Printf("Received file %s from %s\n", fileName, clientAddr)
	}
}
