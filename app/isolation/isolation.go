package isolation

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

// CreateRoot returns the name of a temporary
// folder prepared for chroot
func CreateRoot() string {
	// Steps in creating valid chroot:
	// 1. Create Directory
	// 2. Copy any necessary binaries
	// 3. Copy any necessary dependencies (none in our case)
	rootDir := "mydocker-" + uuid.NewString()
	commands := []*exec.Cmd{}
	commands = append(commands, exec.Command("mkdir", "-p", rootDir))
	commands = append(commands, buildDirStructure(rootDir)...)
	commands = append(commands, copyBins(rootDir)...)
	runCommands(commands)

	return rootDir
}

var directoryStructure []string = []string{
	"/usr/local/bin",
	"/bin",
	"/lib",
	"/lib64",
}

func buildDirStructure(rootDir string) []*exec.Cmd {
	var commands []*exec.Cmd
	for _, dirname := range directoryStructure {
		commands = append(commands,
			exec.Command(
				"mkdir", "-p",
				fmt.Sprintf("%s%s", rootDir, dirname)))
	}
	return commands
}

var minBins []string = []string{
	"/usr/local/bin/docker-explorer",
}

func copyBins(rootDir string) []*exec.Cmd {
	var commands []*exec.Cmd
	for _, binName := range minBins {
		commands = append(commands,
			exec.Command(
				"cp",
				binName,
				fmt.Sprintf("%s%s", rootDir, binName)))
	}
	return commands
}

func runCommands(commands []*exec.Cmd) {
	for _, command := range commands {
		if os.Getenv("DEBUG") != "" {
			log.Println("Running command: ", command)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
		}
		command.Run()
	}
}
