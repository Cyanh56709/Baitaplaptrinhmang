package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Lấy địa chỉ IP của server
	host := "localhost"

	// Lấy cổng từ tham số dòng lệnh
	port := os.Args[1]

	// Kết nối đến server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Gửi chuỗi "abcxyz" đến server
	conn.Write([]byte("abcxyz"))

	fmt.Println("Data sent to server")
}
