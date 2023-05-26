package main

import (
	"fmt"
	//s"io/ioutil"
	"log"
	"bufio"
	"os"
	
	//User-Defined Packages
)

func main() {
    filePath := "./README.md"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error reading file")
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)


	// Iterate over each line and write it to the response
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if scanner.Err() != nil {
		log.Fatalln("Failed to read file")
		return
	}
}