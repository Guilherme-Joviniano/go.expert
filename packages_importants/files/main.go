package main;

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	f, err := os.Create("file.txt")
	
	if err != nil {
		panic(err)	
	}

	// size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World!"))

	if err != nil {
		panic(err)	
	}
	
	fmt.Printf("File created with success! Size: %d bytes \n", size)

	f.Close()

	file, err := os.ReadFile("file.txt")
	
	if err != nil {
		panic(err)	
	}

	
	fmt.Println(string(file))
	
	// read by chunks
	
	file2, err := os.Open("file.txt")
	
	if err != nil {
		panic(err)	
	}
	reader := bufio.NewReader(file2)

	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	// remove file
	err = os.Remove("file.txt")

	if err != nil {
		panic(err)
	}
}