package main

import (

	filefolder "go_cmd/internal/FileFolder"
)


func main() {
	// Variables declaration
	files_add := []string{"bin", "cmd", "internal", "tests"}
	file_perm := 0751
	verbose := false
	verbosePtr := &verbose

	// Flag Function
	filefolder.FlagFunc(verbosePtr)

	// Create Go mod function
	filefolder.CreateGoMod()

	// Getting a file path
	filepath := filefolder.GetFilePath()

	// Directory adding Function
	filefolder.AddingDir(files_add, filepath, file_perm, verbosePtr)

	// Read Directory after reading function
	filefolder.ReadCreatedDir(filepath)

}
