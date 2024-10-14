package main

import (
	"flag"
	"log"
	"os"
)

// function to create
// Create/Open new file
// check if file exists in directory,
// prompt warning if so Y/N?
// SaveFile
// refactor to be released as bin and used in CMD

func main(){
	mkfile := flag.String("mkfile","mkfile", "file to create")
	flag.Parse()

	file, err := os.Create(*mkfile)
	if err != nil {
		log.Fatal(err)
	}	
	defer file.Close()
}