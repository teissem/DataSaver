package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ParseJSON is use to parse json configuration file.
func ParseJSON(source string) (*Configuration, error) {
	byteValueJSON, err := os.ReadFile(filepath.Clean(source))
	if err != nil {
		return nil, fmt.Errorf("reading file : %s", err.Error())
	}
	var configuration Configuration
	err = json.Unmarshal(byteValueJSON, &configuration)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json : %s", err.Error())
	}
	configuration.Destination = CalculateDestination(configuration.Destination)
	return &configuration, nil
}
