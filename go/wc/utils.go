package main

import (
	"fmt"
	"log"
	"os"
)

func GetFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	modificationTime := fileInfo.ModTime()

	fmt.Println("The file ", fileInfo.Name(), " was last modified at:", modificationTime)

	fileSize := fileInfo.Size()

	return fileSize
}
