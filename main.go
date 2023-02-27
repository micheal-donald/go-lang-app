package main

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Hello function
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf("Hello %s", name)
	return message, nil
}

// Main function
func main() {
	message, err := Hello("Mich the awesome")
	if err != nil {
		log.Error("Error:", err)
	}
	fmt.Println(message)
	log.Info(message)
}
