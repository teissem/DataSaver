package configuration

import (
	"encoding/json"
	"os"
)

type GitConfiguration struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Repositories []string `json:"repositories"`
}

type Configuration struct {
	Destination string           `json:"destination"`
	Folder      []string         `json:"folder"`
	Git         GitConfiguration `json:"git"`
}

// parse is use to parse json configuration file
func ParseJSON(path string) (*Configuration, error) {
	byteValueJSON, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var configuration Configuration
	err = json.Unmarshal(byteValueJSON, &configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}
