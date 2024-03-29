package compression

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CompressTar(source, target string) error {
	tarfile, err := os.Create(path.Clean(target))
	if err != nil {
		return fmt.Errorf("create %s : %w", target, err)
	}
	defer func() {
		err = tarfile.Close()
		if err != nil {
			log.Printf("[ERROR] Failed to close destination file")
		}
	}()
	tarball := tar.NewWriter(tarfile)
	defer func() {
		err = tarball.Close()
		if err != nil {
			log.Printf("[ERROR] Failed to close destination file writer")
		}
	}()
	info, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("stat on %s : %w", source, err)
	}
	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}
	return filepath.Walk(source, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("previous error : %w", err)
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return fmt.Errorf("create file info header %s : %w", info.Name(), err)
		}
		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(currentPath, source))
		}
		if err := tarball.WriteHeader(header); err != nil {
			return fmt.Errorf("write header : %w", err)
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path.Clean(currentPath))
		if err != nil {
			return fmt.Errorf("open %s : %w", currentPath, err)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				log.Printf("[ERROR] Failed to close destination file descriptor")
			}
		}()
		_, err = io.Copy(tarball, file)
		if err != nil {
			return fmt.Errorf("copy error : %w", err)
		}
		return nil
	})
}
