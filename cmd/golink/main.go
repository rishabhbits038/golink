package main

import (
	"flag"
	"github.com/rishabhbits038/golink"
	"fmt"
)


func main() {
	source :=flag.String("source", "", "Location of the source directory")
	destination :=flag.String("destination", "", "Location of the destination directory")
	flag.Parse()

	if*source=="" {
		fmt.Println("Invalid Source")
	}
	if *destination=="" {
		fmt.Println("Invalid Destination")
	}

	var ans bool
	if *source!="" && *destination!="" {
		ans=golink.CreateSymLink(*source, *destination);
	}

	if ans==true {
		fmt.Println("Symlink created successfully")
	} else {
		fmt.Println("Unable to create symlink")
	}
}