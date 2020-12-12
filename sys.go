package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

// GetKernelInfo execute `uname -a` using system shell
// and returns output in a string
func GetKernelInfo() string {
	cmd := exec.Command("uname", "-a")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

// GetHostname returns hostname as a string, using system API
func GetHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "Unknown"
		log.Fatal(err)
	}

	return hostname
}
