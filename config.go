package main

import (
	"log"
	"os"
)

func main() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("https://golang.org/pkg/os/#UserConfigDir failed: %v", err)
	}
	println(userConfigDir)
}
