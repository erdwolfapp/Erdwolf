package main

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/app"
)

func handleError(err error) {
	fmt.Printf("Caught an error: %+v\n", err)
}

func main() {
	if err := loadConfig(&appConfig, CONFIG_APPLICATION); err != nil {
		handleError(err)
		return
	}

	if err := loadConfig(&dbConfig, CONFIG_DATABASE); err != nil {
		handleError(err)
		return
	}

	erdwolf := app.NewInstance(appConfig, dbConfig)
	if err := erdwolf.InitHttpServer(); err != nil {
		handleError(err)
		return
	}
	erdwolf.StartListening()
}