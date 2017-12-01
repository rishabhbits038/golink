package golink

import (
	"os/exec"
	"fmt"
	"encoding/json"
)

func CreateSymLink(source string, destination string) bool {
	out, errCheck:= exec.Command("test", "-d", source, "&&", "echo", "true", "||", "echo", "false").Output();
	if errCheck !=nil {
		fmt.Println(errCheck);
		fmt.Println("Unable to check for the directory")
	}
	var checkDir bool
	json.Unmarshal(out, &checkDir)

	cmdRemove := exec.Command("rm", "-rf", source);
	errRemove:=cmdRemove.Run();
	if errRemove != nil && checkDir==true{
		fmt.Println(errRemove)
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