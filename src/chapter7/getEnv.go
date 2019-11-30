package main

import (
	"fmt"
	"os"
)

func main() {
	getenv := os.Getenv("PATH")
	fmt.Println(getenv)
	for key, value := range os.Environ() {
		fmt.Println(key, "=", value)
	}
}
