package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)


func main() {
	// Hard 
	files_add := []string{"bin", "cmd", "internal", "tests"}
	file_perm := 751

	// Get current working directory
	filepath, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}

	// Add Directory after getting current directory
	for _, file := range files_add{
		total_path := filepath + "/" +file
		err := os.Mkdir(total_path, fs.FileMode(file_perm))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added %s folder.\n", file)
	}

	// Read Directory after reading
	files, err := os.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Present in ", filepath, ":")
	for _, file := range files{
		if file.Type().IsDir(){
			fmt.Println(file.Name())
		}
	}

}
