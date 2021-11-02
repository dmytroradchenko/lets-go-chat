package main

import (
	"fmt"
	"os"

	"github.com/dmytroradchenko/lets-go-chat/pkg/hasher"
)

func hashValue(password string) string {
	hash, error := hasher.HashPassword(password)
	if error != nil {
		fmt.Println("Error:", error)
	}
	return hash
}

func checkValue(password, hash string) {
	hashedPassword := hashValue(password)
	if hashedPassword != "" {
		fmt.Printf("Is password correct: %t\n", hashedPassword == hash)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error! No arguments passed")
	} else if len(args) == 1 {
		fmt.Println(hashValue(args[0]))
	} else {
		checkValue(args[0], args[1])
	}
}
