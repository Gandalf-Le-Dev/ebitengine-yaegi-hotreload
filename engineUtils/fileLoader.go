package engineUtils

import (
	"embed"
	"io/fs"
	"os"
)

func LoadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func FindFilesRecursive(src *embed.FS) ([]string, error) {
	var files []string

	err := fs.WalkDir(src, "src", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
