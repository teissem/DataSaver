package datasource

import (
	"io"
	"log"
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
			secondErr := os.RemoveAll(path.Join(destination, srcDest.Destination))
			if secondErr != nil {
				log.Printf("[ERROR] Failed to remove destination folder, a user clean can be necessary")
			}
			return err
		}
	}
	return nil
}

func copyFolder(source string, destination string) error {
	const dirPermission = 0777
	err := os.MkdirAll(destination, dirPermission)
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
			sourceFileBuffer, err := os.Open(path.Join(path.Clean(source), file.Name()))
			if err != nil {
				return err
			}
			defer func() {
				err = sourceFileBuffer.Close()
				if err != nil {
					log.Printf("[ERROR] Failed to close source file buffer")
				}
			}()

			destinationFileBuffer, err := os.Create(path.Join(path.Clean(destination), file.Name()))
			if err != nil {
				return err
			}
			defer func() {
				err = destinationFileBuffer.Close()
				if err != nil {
					log.Printf("[ERROR] Failed to close destination file buffer")
				}
			}()
			_, err = io.Copy(destinationFileBuffer, sourceFileBuffer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
