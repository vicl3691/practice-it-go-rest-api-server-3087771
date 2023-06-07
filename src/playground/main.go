package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := fmt.Println("hi there!")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("hi again !!")
}