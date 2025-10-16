package main

import (
	"fmt"
	"os"
)

func main() {
	/*	fmt.Println("File Handling in Go")

		file, err := os.Create("example.txt")

		if err != nil {
			fmt.Println("Error while creating the file: ", err)
		}
		defer file.Close()

		//steps to write into the files

		content := "hello world by Ankit"
		_, errors := io.WriteString(file, content+"\n")
		if errors != nil {
			fmt.Println("Error while writinf file: ", errors)
			return
		}

		fmt.Println("Sucessfully created file")

		//steps to read into the file
		file, error1 := os.Open("example.txt")
		if error1 != nil {
			fmt.Println("Error while openning files: ", error1)
		}

		//create a buffer to read the file content
		buffer := make([]byte, 1024)

		//read the file content into the buffer
		for {
			n, error1 := file.Read(buffer)
			if error1 == io.EOF {
				break
			}
			if error1 != nil {
				fmt.Println("Error while openning files: ", error1)
				return
			}
			fmt.Println(string(buffer[:n]))
		}
	*/

	// Read the entire file into a byte slice
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return
	}
	fmt.Println(string(content))
}
