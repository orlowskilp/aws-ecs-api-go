package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

const defaultPort string = "8080"

func getKernelInfo() string {
	cmd := exec.Command("uname", "-a")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func getHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "Unknown"
		log.Fatal(err)
	}

	return hostname
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hostname", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, getHostname())
	})

	router.GET("/kernel", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, getKernelInfo())
	})

	return router
}

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

	router := setupRouter()
	router.Run(":" + portString)
	log.Printf("Started server at port %s", portString)
}
