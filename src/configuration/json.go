package configuration

import (
	"encoding/json"
	"errors"
	"os"
)

// ParseJSON is use to parse json configuration file.
func ParseJSON(path string) (*Configuration, error) {
	byteValueJSON, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("[ERROR] Reading file : " + err.Error())
	}
	var configuration Configuration
	err = json.Unmarshal(byteValueJSON, &configuration)
	if err != nil {
		return nil, errors.New("[ERROR] Unmarshal JSON : " + err.Error())
	}
	return &configuration, nil
}
