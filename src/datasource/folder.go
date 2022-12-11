package datasource

import (
	"io"
	"os"
	"path"

	"teissem.fr/data_saver/src/configuration"
)

func GetFolders(folders *configuration.Folder, destination string) error {
	if folders == nil || folders.Path == nil {
		return nil
	}
	for _, srcDest := range folders.Path {
		err := copyFolder(srcDest.Source, path.Join(destination, srcDest.Destination))
		if err != nil {
			os.RemoveAll(path.Join(destination, srcDest.Destination))
			return err
		}
	}
	return nil
}

func copyFolder(source string, destination string) error {
	err := os.MkdirAll(destination, 0777)
	if err != nil {
		return err
	}
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}
	files, err := os.ReadDir(source)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			err := copyFolder(path.Join(source, file.Name()), path.Join(destination, file.Name()))
			if err != nil {
				return err
			}
		} else {
			sourceFileBuffer, err := os.Open(path.Join(source, file.Name()))
			if err != nil {
				return err
			}
			defer sourceFileBuffer.Close()

			destinationFileBuffer, err := os.Create(path.Join(destination, file.Name()))
			if err != nil {
				return err
			}
			defer destinationFileBuffer.Close()
			_, err = io.Copy(destinationFileBuffer, sourceFileBuffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
