package main

import "fmt"

func HandleFirstCLIFlag(flag string) {
	switch flag {
	case "-c":
		fmt.Println("Let's count the number of bytes in a file.")
	case "-l":
		fmt.Println("Let's count the number of lines in a file.")
	case "-w":
		fmt.Println("Let's count the number of words in a file.")
	case "-m":
		fmt.Println("Let's count the number of characters in a file.")
	default:
		fmt.Println("Let's count the number of bytes, lines and words in a file.")
	}
}
