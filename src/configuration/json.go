package configuration

import (
	"encoding/json"
	"os"
)

// ParseJSON is use to parse json configuration file
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
