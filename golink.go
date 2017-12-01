package golink

import (
	"os/exec"
	"fmt"
)

func CreateSymLink(source string, destination string) bool {
	cmdRemove := exec.Command("rm", "-rf", source);
	errRemove:=cmdRemove.Run();
	if errRemove != nil ==true{
		fmt.Println("Unable to remove file "+string(errRemove))
		return false
	}

	cmdLink := exec.Command("ln", "-s", destination);
	errLink:=cmdLink.Run();

	if errLink != nil  {
		fmt.Println(errLink)
		return false
	}

	return true
}