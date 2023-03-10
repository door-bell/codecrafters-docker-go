package isolation

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/door-bell/codecrafters-docker-go/app/helper"
	"github.com/door-bell/codecrafters-docker-go/app/registry"
	"github.com/google/uuid"
)

// CreateRoot returns the name of a temporary
// folder prepared for chroot
func CreateRoot(image, reference string) string {
	rootDir := "mydocker-" + uuid.NewString()
	commands := []*exec.Cmd{}
	commands = append(commands, exec.Command("mkdir", "-p", rootDir))
	commands = append(commands, copyDockerExplorer(rootDir)...)
	commands = append(commands, copyImageContents(image, reference, rootDir)...)
	runCommands(commands)
	return rootDir
}

func copyImageContents(image, reference, rootDir string) []*exec.Cmd {
	// Commands to copy image contents
	cmd := exec.Command("cp", "-a", registry.GetImageFsDir(image, reference)+"/.", rootDir)
	return []*exec.Cmd{cmd}
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
