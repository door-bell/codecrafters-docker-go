package registry

type DockerHubAuthResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

type DockerHubImageManifest struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type DockerHubFsLayer struct {
	BlobSum string `json:"blobSum"`
}
