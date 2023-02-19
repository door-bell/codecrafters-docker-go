package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/door-bell/codecrafters-docker-go/app/isolation"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	// Since we actulaly wanna run chroot, we
	// include the command part with the rest of the arguments.
	// command := os.Args[3]
	// args := os.Args[4:len(os.Args)]
	cmd := buildCommand(os.Args[3:len(os.Args)])
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()

	os.Exit(cmd.ProcessState.ExitCode())
}

func buildCommand(commandAndArgs []string) *exec.Cmd {
	rootName := isolation.CreateRoot()
	commandAndArgs = append([]string{rootName}, commandAndArgs...)
	cmd := exec.Command("chroot", commandAndArgs...)
	wireStdio(cmd)
	return cmd
}

func wireStdio(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
