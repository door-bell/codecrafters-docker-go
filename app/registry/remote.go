package registry

import (
	"strings"
)

// Pull fetches an image from registry.hub.docker.com
// and caches it locally.
func Pull(image string) error {
	split := strings.Split(image, ":")
	imgName := split[0]
	imgReference := split[1]
	fetchManifest(imgName, imgReference)

	return nil
}
