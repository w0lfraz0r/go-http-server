package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Further file operations would go here.

	data := make([]byte, 8)
	for {
		count, err := file.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("read: %s\n", data[:count])
	}
}
