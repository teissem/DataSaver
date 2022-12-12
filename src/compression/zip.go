package compression

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func CompressZip(source, target string) error {
	f, err := os.Create(path.Clean(target))
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Printf("[ERROR] Failed to close destination file")
		}
	}()
	writer := zip.NewWriter(f)
	defer func() {
		err = writer.Close()
		if err != nil {
			log.Printf("[ERROR] Failed to close destination file writer")
		}
	}()
	return filepath.Walk(source, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(source), currentPath)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path.Clean(currentPath))
		if err != nil {
			return err
		}
		defer func() {
			err = f.Close()
			if err != nil {
				log.Printf("[ERROR] Failed to close file descriptor")
			}
		}()
		_, err = io.Copy(headerWriter, f)
		return err
	})
}
