package main

import (
	"log"
	"os"
	"strconv"
)

const defaultPort string = "8080"

func main() {
	var portString string

	if len(os.Args) == 1 {
		portString = defaultPort
		log.Printf("Using default port %s", portString)
	} else {
		port, err := strconv.Atoi(os.Args[1])

		if err != nil {
			log.Fatalf("Invalid port value rolling back to default %s",
				defaultPort)
		}

		if port < 0 || port > 65535 {
			log.Fatalf("Port value out of range rolling back to default %s",
				defaultPort)
		}

		portString = strconv.Itoa(port)
	}

	router := SetupRouter()
	router.Run(":" + portString)
	log.Printf("Started server at port %s", portString)
}
