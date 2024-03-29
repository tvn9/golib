// Reading/Writing from the child process
package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)

	// Output will run the process terminates and returns the standard
	// output in the byte slice.
	buff, err := proc.Output()
	if err != nil {
		log.Fatal(err)
	}

	// The output of child process in form of byte slice printed as string
	fmt.Println(string(buff))
}
