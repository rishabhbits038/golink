package cmd

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
		fmt.Print("Invalid Source")
	}
	if *destination=="" {
		fmt.Print("Invalid Destination")
	}

	var ans bool
	if *source!="" && *destination!="" {
		ans=golink.createSymLink(source, destination);
	}

	if ans==true {
		fmt.Print("Symlink created successfully")
	} else {
		fmt.Print("Unable to create symlink")
	}
}