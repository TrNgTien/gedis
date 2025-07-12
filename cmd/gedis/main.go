package main

import (
	"fmt"
	"net"
	"os"
)


func main(){
  fmt.Println("Gedis is started at port 6379.....")

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

  defer l.Close()
}
