package cmd

import (
	"fico_ar/delivery/container"
	"fico_ar/delivery/http"
	"fmt"
	"log"
	"os"
)

func Execute() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("Starting App...")
	container := container.SetupContainer()
	handler := http.SetupHandler(container)

	http := http.ServerHttp(handler)
	http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
}
