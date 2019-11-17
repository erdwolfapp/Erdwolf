package main

import (
	"github.com/erdwolfapp/Erdwolf/web"
)

func main() {
	config := web.ErdwolfConfig{
		Port: 8080,
	}
	server := web.NewWebServer(config)

	server.Start()
}