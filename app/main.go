package main

import (
	"log"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// We can pass one of these SysProcAttr into os.StartProcess()
	// instead of using exec.command: cmd.SysProcAttr.Chroot
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()

	os.Exit(cmd.ProcessState.ExitCode())
}

func make_root() {
	// Create directory to chroot to, copy directories
}
