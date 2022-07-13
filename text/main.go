package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(content))

}
