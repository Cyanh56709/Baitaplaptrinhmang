package main

// Thi thoảng bị lỗi error sending data

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("listening on port 8080")

	// vòng lặp vô hạn để chấp nhận các kết nối đến từ client
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}
		fmt.Println("new client connected.")
		// sử dụng goroutine để xử lý kết nối từng client
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// đọc dữ liệu từ client và gửi lại phản hồi
	scanner := bufio.NewScanner(conn)
	for {
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		// in ra thông tin từ client
		fmt.Printf("Received message: %s\n", text)

		// đọc dữ liệu từ bàn phím
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter text: \n")
		scanner.Scan()
		text = scanner.Text()

	}

	fmt.Println("Client disconnected.")
}
