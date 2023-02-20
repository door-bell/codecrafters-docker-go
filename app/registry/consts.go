package registry

// Local layer / container cache
const LOCAL_IMAGE_REPO = "/var/lib/mydocker"

// Remote docker repository
const REGISTRY_URL = "registry.hub.docker.com"
const REGISTRY_AUTH_URL = "https://auth.docker.io/token?service=" + REGISTRY_URL
