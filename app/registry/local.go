package registry

import (
	"fmt"
	"os"
	"os/exec"
)

func GetImageFsDir(image, reference string) string {
	return fmt.Sprintf("%s/%s/%s/img", LOCAL_IMAGE_REPO, image, reference)
}

func extractLayer(imgName, imgReference, digest string) error {
	imgDirname := GetImageFsDir(imgName, imgReference)
	filename := getLayerFilename(imgName, imgReference, digest)
	err := os.MkdirAll(imgDirname, os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("tar", "-xzf", filename, "--directory", imgDirname)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
