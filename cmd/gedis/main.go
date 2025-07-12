package main

import (
	"fmt"

	"github.com/TrNgTien/gedis.git/internal/server"
)


func main(){
  fmt.Println("Gedis is started at port 6379.....")

  server.RunAsyncServer()
}

