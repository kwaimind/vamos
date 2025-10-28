package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func FindPackageJson(rootDir string, packageName string, ignoreFiles []string, config *Config) (string, error) {
	var result string

	err := filepath.WalkDir(rootDir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			isRoot := path == rootDir
			shouldIgnore := slices.Contains(ignoreFiles, d.Name())
			skipDotFiles := strings.HasPrefix(d.Name(), ".")

			if !isRoot && d.IsDir() && (shouldIgnore || skipDotFiles) {
				return filepath.SkipDir
			}

			if d.Name() == config.PackageJSONName {

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				data, err := ParseJson(file)
				if err != nil {
					return err
				}

				if data.Name == packageName {
					result = path
					return nil
				}
			}

			return nil
		})

	return result, err
}
