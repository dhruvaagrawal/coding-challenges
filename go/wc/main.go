package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	args := os.Args

	fmt.Println(args)

	cliFlag := args[1]

	HandleFirstCLIFlag(cliFlag)

	// fileSize := GetFileSize(filePath)

	// fmt.Println("The file size is:", fileSize)

}
