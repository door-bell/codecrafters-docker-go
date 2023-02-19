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
	commands = append(commands, copyFolders(rootDir)...)
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
	"/bin/ps",
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

var foldersToCopy []string = []string{
	"/lib",
	"/bin",
}

func copyFolders(rootDir string) []*exec.Cmd {
	var commands []*exec.Cmd
	for _, binName := range foldersToCopy {
		commands = append(commands,
			exec.Command(
				"cp", "-a",
				binName,
				fmt.Sprintf("%s%s", rootDir, binName)))
	}
	return commands
}

func runCommands(commands []*exec.Cmd) error {
	for _, command := range commands {
		if os.Getenv("DEBUG") != "" {
			log.Println("Running command to confgure chroot: ", command)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
		}
		err := command.Run()
		if err != nil {
			log.Println("Error on command", command)
			log.Fatal(err)
			return err
		}
	}
	return nil
}
