package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	args := os.Args

	fmt.Println(args)

	argsWithoutProgPath := args[1:]

	fmt.Println(argsWithoutProgPath)
}
