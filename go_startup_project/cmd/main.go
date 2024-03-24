package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)


func main() {
	// Hard 
	files_add := []string{"bin", "cmd", "internal", "tests"}
	file_perm := 751

	// Add verbose Flags
	verbose := flag.Bool("verbose", false, "Print out output when invoked")
	verboseShortHand := flag.Bool("v", false, "Print out output when invoked")

	flag.Parse()

	// Adding a for each loop and switch to check for flags
	for i := 0; i < flag.NArg(); i++ {
		switch flag.Arg(i){
			case "--verbose":
				*verbose = true
			case "-v":
				*verboseShortHand = true
		}
	}

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
		if *verbose || *verboseShortHand{
			fmt.Printf("Added %s folder.\n", file)
		}
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
