package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/door-bell/codecrafters-docker-go/app/isolation"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	// Since we actulaly wanna run chroot, we
	// include the command part with the rest of the arguments.
	// command := os.Args[3]
	// args := os.Args[4:len(os.Args)]
	cmd := buildRunCommand(os.Args[3:len(os.Args)])
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()

	os.Exit(cmd.ProcessState.ExitCode())
}

func buildRunCommand(commandAndArgs []string) *exec.Cmd {
	rootName := isolation.CreateRoot()
	commandStr := commandAndArgs[0]
	commandArgs := commandAndArgs[1:]
	cmd := exec.Command(commandStr, commandArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot:     rootName,
		Cloneflags: syscall.CLONE_NEWPID,
	}
	if os.Getenv("DEBUG") != "" {
		log.Println("Final command:", cmd.Path, cmd.Args)
	}
	return cmd
}
