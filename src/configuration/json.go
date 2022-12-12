package configuration

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// ParseJSON is use to parse json configuration file.
func ParseJSON(source string) (*Configuration, error) {
	byteValueJSON, err := os.ReadFile(filepath.Clean(source))
	if err != nil {
		return nil, errors.New("[ERROR] Reading file : " + err.Error())
	}
	var configuration Configuration
	err = json.Unmarshal(byteValueJSON, &configuration)
	if err != nil {
		return nil, errors.New("[ERROR] Unmarshal JSON : " + err.Error())
	}
	configuration.Destination = CalculateDestination(configuration.Destination)
	return &configuration, nil
}
