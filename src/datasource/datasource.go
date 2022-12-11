package datasource

import "teissem.fr/data_saver/src/configuration"

func GetData(configuration *configuration.Configuration) error {
	err := GetFolders(&configuration.FolderSource, configuration.Destination)
	if err != nil {
		return err
	}
	err = GetGitRepositories(&configuration.GitSource, configuration.Destination)
	return err
}
