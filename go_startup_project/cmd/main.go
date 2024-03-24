package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
)


func main() {
	// Variables declaration
	files_add := []string{"bin", "cmd", "internal", "tests"}
	file_perm := 0751
	verbose := false

	// Add verbose Flags
	flag.BoolFunc("verbose", "Print out output when invoked", func (s string) error {
		verbose = true
		return nil
	})
	flag.BoolFunc("v", "Print out output when invoked", func (s string) error {
		verbose = true
		return nil
	})
	flag.BoolFunc("readme", "Creates a README.md file", CreateREADME)
	flag.BoolFunc("re", "Creates a README.md file", CreateREADME)

	flag.Parse()

	// Adding a for each loop to find args that contains a "/"
	for i := 0; i < flag.NArg(); i++ {
		args := flag.Arg(i)
		if strings.Contains(args, "/"){
			// sentence = args
			cmd := exec.Command("go", "mod", "init", args)
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
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

		// Add main.go file under /cmd
		if file == "cmd"{
			f, err := os.Create(total_path + "/main.go")
			if err != nil {
				log.Fatal(err)
			}
			f.WriteString("package main")
			defer f.Close()
		}
		if verbose{
			fmt.Printf("Added %s directory.\n", file)
		}
	}

	// Read Directory after reading
	ReadCreatedDir(filepath)

}

func ReadCreatedDir(filepath string) {
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

func CreateREADME(s string) error {
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("# README")
	defer f.Close()
	return nil
	
}