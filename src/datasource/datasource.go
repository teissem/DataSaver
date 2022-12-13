package datasource

import (
	"fmt"
	"os"

	"teissem.fr/data_saver/src/configuration"
)

func GetData(configuration *configuration.Configuration) error {
	const dirPermission = 0777
	err := os.MkdirAll(configuration.Destination, dirPermission)
	if err != nil {
		return fmt.Errorf("mkdir all %s : %s", configuration.Destination, err.Error())
	}
	err = GetFolders(&configuration.FolderSource, configuration.Destination)
	if err != nil {
		return fmt.Errorf("get folders : %s", err.Error())
	}
	err = GetGitRepositories(&configuration.GitSource, configuration.Destination)
	if err != nil {
		return fmt.Errorf("get git repositories : %s", err.Error())
	}
	return nil
}
