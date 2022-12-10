package datasource

import "teissem.fr/data_saver/src/configuration"

func GetData(configuration *configuration.Configuration) error {
	err := GetFolders(configuration.FolderSource.Path, configuration.Destination)
	if err != nil {
		return err
	}
	err = GetGitRepositories(configuration.GitSource.Repositories, configuration.Destination)
	return err
}
