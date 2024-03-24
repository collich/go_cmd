package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)


func main() {
	files_add := []string{"bin", "cmd", "internal", "tests"}
	file_perm := 751

	filepath, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}

	for _, file := range files_add{
		total_path := filepath + file
		err := os.Mkdir(total_path, fs.FileMode(file_perm))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added file %s in %s with permission %s\n", file, filepath, fs.FileMode(file_perm).String())
	}

	
	// fmt.Println(filepath)
}
