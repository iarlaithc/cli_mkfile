package main

import (
	"fmt"
	"mkfile/mkfile"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Printf("Usage: mkfile <filename>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Printf("Usage: mkfile <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("Error - File Already Exists")
		os.Exit(1)
	}
	err = mkfile.CreateFile(filename)
	if err != nil {
		fmt.Printf("Error - Failed To Create File: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%v Created", filename)	
}