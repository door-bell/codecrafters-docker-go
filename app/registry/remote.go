package registry

import (
	"log"

	"github.com/door-bell/codecrafters-docker-go/app/helper"
)

// Pull fetches an image from registry.hub.docker.com
// and caches it locally.
func Pull(imageStr string) error {
	// 1. Get token to pull given image
	// 2. Get manifest list
	// 3. Search manifest list for given os / architecture
	// 4. Fetch manifest
	// 5. Pull image layers and cache them to be used in a container
	image := NewImage(imageStr)
	imgName := image.Name
	imgReference := image.Reference

	token, err := getDockerHubToken(imgName)
	if err != nil {
		log.Println("Error fetching docker hub token!")
		log.Fatal(err)
	}

	manifest, err := fetchManifest(imgName, imgReference, token)
	if err != nil {
		log.Println("Error fetching image manifest!")
		log.Fatal(err)
	}

	for _, layer := range manifest.Layers {
		helper.DebugLog("Pulling layer: ", layer.Digest)
		err = downloadLayer(imgName, imgReference, layer.Digest, token)
		if err != nil {
			log.Println("Error downloading layer!")
			log.Fatal(err)
		}
		err = extractLayer(imgName, imgReference, layer.Digest)
		if err != nil {
			log.Println("Error extracting layer!")
			log.Fatal(err)
		}
	}

	return nil
}
