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
		return nil, fmt.Errorf("reading file : %w", err)
	}
	var configuration Configuration
	err = json.Unmarshal(byteValueJSON, &configuration)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json : %w", err)
	}
	configuration.Destination = CalculatePath(configuration.Destination)
	configuration.Log = CalculatePath(configuration.Log)
	return &configuration, nil
}
