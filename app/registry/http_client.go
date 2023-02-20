package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/door-bell/codecrafters-docker-go/app/helper"
)

// Client used for calls to docker hub
var client = &http.Client{}

var savedToken string = ""

func fetchManifest(imgName string, imgReference string) {
	token, _ := getDockerHubToken()
	savedToken = token

	bodyBytes := getRawResponse("GET", getManifestUrl(imgName, imgReference))
	var manifest DockerHubImageManifest
	err := json.Unmarshal(bodyBytes, &manifest)
	helper.DebugLog("Manifest:", manifest.Name, manifest.Tag)
	if err != nil {
		log.Fatal(err)
	}
}

func getDockerHubToken() (string, error) {
	bodyBytes := getRawResponse("GET", REGISTRY_AUTH_URL)
	var responseBody DockerHubAuthResponse
	err := json.Unmarshal(bodyBytes, &responseBody)
	if err == nil {
		helper.DebugLog("Auth Token:", responseBody.Token)
		return responseBody.Token, nil
	} else {
		log.Println("Error deserializing registry token:", err)
		return "", err
	}
}

func getRawResponse(verb string, url string) []byte {
	req, _ := http.NewRequest(verb, url, nil)
	req.Header.Set("Accept", "application/json")
	if savedToken != "" {
		req.Header.Set("Authorization", "Bearer "+savedToken)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error completing request", verb, url, err)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response", verb, url, err)
	}
	helper.DebugLog(verb, url, string(bodyBytes))
	return bodyBytes
}

func getManifestUrl(name string, reference string) string {
	return fmt.Sprintf("https://%s/v2/%s/manifests/%s", REGISTRY_URL, name, reference)
}
