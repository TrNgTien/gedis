package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func RunAsyncServer() {
	l, err := net.Listen("tcp", "0.0.0.0:6377")

	if err != nil {
		fmt.Println("Failed to bind to port 6377")
		os.Exit(1)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		defer conn.Close()

		buf := make([]byte, 1024)
		addr := conn.RemoteAddr().String()
		length, err := conn.Read(buf)
		if err == io.EOF {
			log.Printf("[%s] Connection Closed\n", addr)
			return
		} else if err != nil {
			log.Printf("Error reading: %#v\n", err)
			return
		}

		rawMessage := string(buf[:length])
		fmt.Println("rawMessage: ", rawMessage)
	}
}
