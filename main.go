package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	mkfile := flag.String("mkfile","mkfile", "file to create")
	flag.Parse()

	fmt.Printf(flag.Args()[0])
	if flag.NArg() == 0 {
		fmt.Printf("could not create: %s", *mkfile)
	} else if flag.Arg(0) == "mkfile" {
		file, err := os.Create(*mkfile)
		if err != nil {
			fmt.Println(err)
		}	
		defer file.Close()
		fmt.Println(file.Name())
	}
}