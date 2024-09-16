package main

import (
	"fmt"
	"net"

	"github.com/MeroFuruya/go-satisfactory/lightqueryapi"
)

func main() {
	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1")
		Port: 12345,
	}
	client, err := lightqueryapi.NewClient(addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	fmt.Println(client)
}