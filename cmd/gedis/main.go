package main

import (
	"fmt"

	"gedis/internal/server"
)

func main() {
	fmt.Println("Gedis is started at port 6379.....")

	server.RunAsyncServer()
}
