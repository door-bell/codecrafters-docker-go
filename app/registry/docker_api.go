package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/door-bell/codecrafters-docker-go/app/helper"
)

// Client used for calls to docker hub
var client = &http.Client{}

// Downloads a compressed layer to the correct cache folder
func downloadLayer(imgName, imgReference, digest, token string) error {
	dirname := fmt.Sprintf("%s/%s/%s", LOCAL_IMAGE_REPO, imgName, imgReference)
	filename := fmt.Sprintf("%s/%s", dirname, digest)
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("GET", getLayerUrl(imgName, digest), nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/vnd.oci.image.layer.v1.tar+gzip")
	req.Header.Add("Accept-Encoding", "gzip")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	helper.DebugLog(file.Name(), size)

	return nil
}

func fetchManifest(imgName, imgReference, token string) (*ImageManifest, error) {
	index, err := fetchManifestIndex(imgName, imgReference, token)
	if err != nil {
		return nil, err
	}
	thisPlatform := ImagePlatform{runtime.GOARCH, runtime.GOOS}
	for _, manifestRef := range index.Manifests {
		if manifestRef.Platform == thisPlatform {
			return getImageManifestFromRef(imgName, token, &manifestRef)
		}
	}
	return nil, errors.New("no image manifest found for the current platform")
}

func getImageManifestFromRef(imgName, token string, manifestRef *ImageManifestRef) (*ImageManifest, error) {
	responseBody, err := responseBytes(
		"GET",
		getManifestUrl(imgName, manifestRef.Digest),
		"manifest",
		token,
	)
	if err != nil {
		return nil, err
	}
	var imgManifest *ImageManifest = &ImageManifest{}
	err = json.Unmarshal(responseBody, imgManifest)
	if err != nil {
		return nil, err
	}
	return imgManifest, nil
}

func fetchManifestIndex(imgName, imgReference, token string) (*ManifestIndex, error) {
	responseBody, err := responseBytes("GET", getManifestUrl(imgName, imgReference), "index", token)
	if err != nil {
		return nil, err
	}
	var imageIndex *ManifestIndex = &ManifestIndex{}
	err = json.Unmarshal(responseBody, imageIndex)
	if err != nil {
		return nil, err
	}
	return imageIndex, nil
}

func getDockerHubToken(image string) (string, error) {
	responseBody, err := responseBytes("GET", registryAuthUrl(image), "json", "")
	if err != nil {
		return "", err
	}
	var authResponse AuthResponse
	err = json.Unmarshal(responseBody, &authResponse)
	if err == nil {
		return authResponse.Token, nil
	} else {
		log.Println("Error deserializing registry token:", err)
		return "", err
	}
}

func responseBytes(verb, url, contentType, token string) ([]byte, error) {
	req, _ := http.NewRequest(verb, url, nil)
	acceptType := "*/*"
	switch contentType {
	case "index":
		acceptType = "application/vnd.oci.image.index.v1+json"
	case "manifest":
		acceptType = "application/vnd.oci.image.manifest.v1+json"
	case "json":
		acceptType = "application/json"
	}
	req.Header.Add("Accept", acceptType)

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response", verb, url, err)
		return nil, err
	}
	helper.DebugLog(string(bodyBytes))
	return bodyBytes, nil
}
