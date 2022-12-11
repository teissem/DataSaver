package datasource

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"teissem.fr/data_saver/src/configuration"
)

func GetGitRepositories(gitConfig *configuration.Git, destination string) error {
	if gitConfig == nil || gitConfig.Repositories == nil {
		return nil
	}
	for _, srcDest := range gitConfig.Repositories {
		_, err := git.PlainClone(path.Join(destination, srcDest.Destination), false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: gitConfig.Username,
				Password: gitConfig.Password,
			},
			URL:      srcDest.Source,
			Progress: os.Stdout,
		})
		if err != nil {
			return err
		}
		err = os.RemoveAll(path.Join(destination, srcDest.Destination, ".git"))
		if err != nil {
			return err
		}
	}
	return nil
}
