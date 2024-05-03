package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	args := os.Args

	fmt.Println(args)

	filePath := args[1]

	fileSize := GetFileSize(filePath)

	fmt.Println("The file size is:", fileSize)

}
