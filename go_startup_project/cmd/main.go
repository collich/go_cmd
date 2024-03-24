package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	// 751
	// files_add := []string{"bin", "cmd", "internal", "tests"}
	// file_perm := 751

	filepath, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(filepath)
}
