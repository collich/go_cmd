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

	// Add verbose Flags
	verbose := flag.Bool("verbose", false, "Print out output when invoked")
	verboseShortHand := flag.Bool("v", false, "Print out output when invoked")
	// readme := flag.Bool("readme", false, "Creates a README.md file")
	// readmeShortHand := flag.Bool("re", false, "Creates a README.md file")
	flag.BoolFunc("readme", "Creates a README.md file", CreateREADME)
	flag.BoolFunc("re", "Creates a README.md file", CreateREADME)

	flag.Parse()

	// Adding a for each loop and switch to check for flags
	for i := 0; i < flag.NArg(); i++ {
		var sentence string
		args := flag.Arg(i)
		if strings.Contains(args, "/"){
			sentence = args
		}
		// Add file go.mod file
		switch args{
			case "--verbose":
				*verbose = true
			case "-v":
				*verboseShortHand = true
			// Readme flags
			// case "--readme":
			// 	*readme = true
			// 	fmt.Println("Readme is : ", *readme)
			// case "-re":
			// 	*readmeShortHand = true
			// 	fmt.Println("Readme is : ", *readmeShortHand)
			// Added a go mod init case 
			case sentence:
				cmd := exec.Command("go", "mod", "init", sentence)
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

func CreateREADME(s string) error {
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("# README")
	defer f.Close()
	return nil
	
}