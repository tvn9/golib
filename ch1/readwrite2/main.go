// Reading/Writing from the child process
package main

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := "ping"
	timeout := 20 * time.Second

	// The command line tool "ping" is executed for 2 seconds
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, cmd, "google.com")

	// The process output is obtained in form of io.ReadCloser.
	// The underlying implementation use the os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// Start the process
	proc.Start()

	// For more compfortable reading the bufio.Scanner is used.
	// The read call is blocking.
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
