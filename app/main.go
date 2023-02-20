package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/door-bell/codecrafters-docker-go/app/helper"
	"github.com/door-bell/codecrafters-docker-go/app/isolation"
	"github.com/door-bell/codecrafters-docker-go/app/registry"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[1]
	switch command {
	case "pull":
		registry.Pull(os.Args[2])
	case "run":
		handleRun()
	default:
		log.Fatal("Unknown command:", command)
	}
}

func handleRun() {
	// if !registry.ExistsLocally(image) {
	//		registry.Pull(image)
	// }
	// buildRunCommand(image, args)

	image := os.Args[3]
	cmd := buildRunCommand(image, os.Args[3:])
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()

	os.Exit(cmd.ProcessState.ExitCode())
}

func buildRunCommand(image string, commandAndArgs []string) *exec.Cmd {
	rootName := isolation.CreateRoot(image)
	commandName := commandAndArgs[0]
	commandArgs := commandAndArgs[1:]
	cmd := exec.Command(commandName, commandArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot:     rootName,
		Cloneflags: syscall.CLONE_NEWPID,
	}
	helper.DebugLog("Final command:", cmd.Path, cmd.Args)
	return cmd
}
