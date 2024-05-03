package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	args := os.Args

	fmt.Println(args)

	if len(args) > 1 {
		HandleFirstCLIFlag(args)
	} else {
		fmt.Println("Please enter valid CLI flags.")
	}

	// fileSize := GetFileSize(filePath)

	// fmt.Println("The file size is:", fileSize)

}
