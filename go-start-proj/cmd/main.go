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
	verbosePtr := &verbose

	// Flag Function
	FlagFunc(verbosePtr)

	// Create Go mod function
	CreateGoMod()

	// Getting a file path
	filepath := GetFilePath()

	// Directory adding Function
	AddingDir(files_add, filepath, file_perm, verbosePtr)

	// Read Directory after reading function
	ReadCreatedDir(filepath)

}

func AddingDir(files_add []string, filepath string, file_perm int, verbose *bool)  {
	// Add Directory after getting current directory
	goFile := ""
	for _, file := range files_add{
		total_path := filepath + "/" +file
		err := os.Mkdir(total_path, fs.FileMode(file_perm))
		if err != nil {
			log.Fatal(err)
		}

		// Add main.go file under /cmd
		if file == "cmd"{
			// Function to create main.go
			goFile = "main"
			CreateGoFile(total_path + "/main.go", goFile)
		} 
			
		if file == "tests"{
			goFile = "tests"
			CreateGoFile(total_path + "/test.go", goFile)
		}

		if *verbose{
			fmt.Printf("Added %s directory.\n", file)
		}
	}
}

func GetFilePath() string {
	// Get current working directory
	filepath, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	return filepath
}

func FlagFunc(verbose *bool) {
		// Add verbose Flags
		flag.BoolFunc("verbose", "Print out output when invoked", func (s string) error {
			*verbose = true
			return nil
		})
		flag.BoolFunc("v", "Print out output when invoked", func (s string) error {
			*verbose = true
			return nil
		})
		flag.BoolFunc("readme", "Creates a README.md file", CreateREADME)
		flag.BoolFunc("re", "Creates a README.md file", CreateREADME)
	
		flag.Parse()
}

func CreateGoMod()  {
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
}

func CreateGoFile(total_path string, goFile string) {
	f, err := os.Create(total_path)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("package " + goFile)
	defer f.Close()
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