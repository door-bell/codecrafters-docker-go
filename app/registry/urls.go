package registry

import "fmt"

// Local layer / container cache
const LOCAL_IMAGE_REPO = "/var/lib/mydocker"

// Remote docker repository
const REGISTRY_URL = "registry.hub.docker.com"
const REGISTRY_AUTH_URL = "auth.docker.io"

func registryAuthUrl(image string) string {
	return fmt.Sprintf("https://%s/token?service=registry.docker.io&scope=repository:library/%s:pull",
		REGISTRY_AUTH_URL, image)
}

// Get URL for manifest list or manifest
func getManifestUrl(name string, reference string) string {
	return fmt.Sprintf("https://%s/v2/library/%s/manifests/%s", REGISTRY_URL, name, reference)
}

func getLayerUrl(image, digest string) string {
	return fmt.Sprintf("https://%s/v2/library/%s/blobs/%s", REGISTRY_URL, image, digest)
}
