package datasource

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"teissem.fr/data_saver/src/configuration"
)

func GetFolders(folders *configuration.Folder, destination string) error {
	if folders == nil || folders.Path == nil {
		return fmt.Errorf("folders = %v, folders.Path = %v", folders, folders.Path)
	}
	for _, srcDest := range folders.Path {
		err := copyFolder(srcDest.Source, path.Join(destination, srcDest.Destination))
		if err != nil {
			secondErr := os.RemoveAll(path.Join(destination, srcDest.Destination))
			if secondErr != nil {
				log.Printf("[ERROR] Failed to remove destination folder, a user clean can be necessary")
			}
			return fmt.Errorf("copy folder : %w", err)
		}
		log.Printf("[INFO] Successfully copy %s to %s", srcDest.Source, path.Join(destination, srcDest.Destination))
	}
	return nil
}

func copyFolder(source string, destination string) error {
	const dirPermission = 0777
	err := os.MkdirAll(destination, dirPermission)
	if err != nil {
		return fmt.Errorf("mkdir all %s : %w", destination, err)
	}
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return fmt.Errorf("stat %s : %w", source, err)
	}
	files, err := os.ReadDir(source)
	if err != nil {
		return fmt.Errorf("read dir %s : %w", source, err)
	}
	for _, file := range files {
		if file.IsDir() {
			err := copyFolder(path.Join(source, file.Name()), path.Join(destination, file.Name()))
			if err != nil {
				return fmt.Errorf("copy folder : %w", err)
			}
		} else {
			sourceFileBuffer, err := os.Open(path.Join(path.Clean(source), file.Name()))
			if err != nil {
				return fmt.Errorf("open file %s : %w", path.Clean(source), err)
			}
			defer func() {
				err = sourceFileBuffer.Close()
				if err != nil {
					log.Printf("[ERROR] Failed to close source file buffer")
				}
			}()

			destinationFileBuffer, err := os.Create(path.Join(path.Clean(destination), file.Name()))
			if err != nil {
				return fmt.Errorf("create file %s : %w", path.Clean(destination), err)
			}
			defer func() {
				err = destinationFileBuffer.Close()
				if err != nil {
					log.Printf("[ERROR] Failed to close destination file buffer")
				}
			}()
			_, err = io.Copy(destinationFileBuffer, sourceFileBuffer)
			if err != nil {
				return fmt.Errorf("copy file : %w", err)
			}
		}
	}
	return nil
}
