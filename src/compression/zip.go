package compression

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func CompressZip(source, target string) error {
	file, err := os.Create(path.Clean(target))
	if err != nil {
		return fmt.Errorf("create %s : %s", target, err.Error())
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("[ERROR] Failed to close destination file")
		}
	}()
	writer := zip.NewWriter(file)
	defer func() {
		err = writer.Close()
		if err != nil {
			log.Printf("[ERROR] Failed to close destination file writer")
		}
	}()
	return filepath.Walk(source, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("previous error : %s", err.Error())
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("create file info header : %s", err.Error())
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(source), currentPath)
		if err != nil {
			return fmt.Errorf("relative path for %s of %s : %s", currentPath, filepath.Dir(source), err.Error())
		}
		if info.IsDir() {
			header.Name += "/"
		}
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("create header : %s", err.Error())
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path.Clean(currentPath))
		if err != nil {
			return fmt.Errorf("open file %s : %s", currentPath, err.Error())
		}
		defer func() {
			err = file.Close()
			if err != nil {
				log.Printf("[ERROR] Failed to close file descriptor")
			}
		}()
		_, err = io.Copy(headerWriter, file)
		if err != nil {
			return fmt.Errorf("copy file : %s", err.Error())
		}
		return err
	})
}
