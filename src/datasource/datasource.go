package datasource

import (
	"os"

	"teissem.fr/data_saver/src/configuration"
)

func GetData(configuration *configuration.Configuration) error {
	err := os.MkdirAll(configuration.Destination, 0777)
	if err != nil {
		return err
	}
	err = GetFolders(&configuration.FolderSource, configuration.Destination)
	if err != nil {
		return err
	}
	err = GetGitRepositories(&configuration.GitSource, configuration.Destination)
	return err
}
