package registry

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

type ManifestIndex struct {
	Manifests []ImageManifestRef `json:"manifests"`
}

type ImageManifestRef struct {
	Digest    string        `json:"digest"`
	MediaType string        `json:"mediaType"`
	Platform  ImagePlatform `json:"platform"`
}

type ImageManifest struct {
	Layers []ImageLayer `json:"layers"`
}

type ImageLayer struct {
	Digest string `json:"digest"`
}

type ImagePlatform struct {
	Architecture string `json:"architecture"`
	Os           string `json:"os"`
}
