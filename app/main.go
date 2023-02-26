package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
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
	// 	registry.Pull(image)
	// }
	// buildRunCommand(image, args)
	image := os.Args[2]
	command := os.Args[3]
	registry.Pull(image)
	cmd := buildRunCommand(image, command, os.Args[3:])
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()

	os.Exit(cmd.ProcessState.ExitCode())
}

func buildRunCommand(image, command string, commandAndArgs []string) *exec.Cmd {
	split := strings.Split(image, ":")
	imgName := split[0]
	imgReference := split[1]

	rootName := isolation.CreateRoot(imgName, imgReference)
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
