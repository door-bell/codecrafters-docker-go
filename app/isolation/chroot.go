package isolation

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/door-bell/codecrafters-docker-go/app/helper"
	"github.com/google/uuid"
)

// CreateRoot returns the name of a temporary
// folder prepared for chroot
func CreateRoot(image string) string {
	// Steps in creating valid chroot:
	// 1. Create Directory
	// 2. Copy any necessary binaries
	// 3. Copy any necessary dependencies (none in our case)
	rootDir := "mydocker-" + uuid.NewString()
	commands := []*exec.Cmd{}
	commands = append(commands, exec.Command("mkdir", "-p", rootDir))
	commands = append(commands, copyDockerExplorer(rootDir)...)
	commands = append(commands, copyImageContents(image, rootDir)...)
	runCommands(commands)
	return rootDir
}

func copyImageContents(image string, rootDir string) []*exec.Cmd {
	// Commands to copy image contents
	return []*exec.Cmd{}
}

func copyDockerExplorer(rootDir string) []*exec.Cmd {
	return []*exec.Cmd{
		exec.Command("mkdir", "-p", fmt.Sprintf("%s%s", rootDir, "/usr/local/bin")),
		exec.Command("cp", "/usr/local/bin/docker-explorer", fmt.Sprintf("%s%s", rootDir, "/usr/local/bin")),
	}
}

func runCommands(commands []*exec.Cmd) {
	for _, command := range commands {
		if helper.IsDebug() {
			log.Println("Running command to confgure chroot: ", command)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
		}
		err := command.Run()
		if err != nil {
			log.Println("Error on command", command)
			log.Fatal(err)
		}
	}
}
