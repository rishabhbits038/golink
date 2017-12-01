package golink

import (
	"os/exec"
	"log"
	"fmt"
)

func CreateSymLink(source string, destination string) bool {
	ans:=true
	cmd := exec.Command("rm", "-rf", "~/dev/src/event_enricher/vendor/bitbucket.org/myntra/go_template_manager*");
	err:=cmd.Run();
	if err != nil {
		ans=false
		log.Fatal(err)
	}
	fmt.Printf("Removed directory");
	return ans
}