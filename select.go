package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// DetectFromLockfile detects package manager from lockfiles in the given directory
func DetectFromLockfile(dir string, config *Config) string {
	lockfiles := map[string]string{
		"bun.lockb":        config.Bun,
		"pnpm-lock.yaml":   config.PNPM,
		"yarn.lock":        config.Yarn,
		"package-lock.json": config.NPM,
	}

	for lockfile, pm := range lockfiles {
		if _, err := os.Stat(filepath.Join(dir, lockfile)); err == nil {
			return pm
		}
	}

	return ""
}

func Select(data *PackageJson, dir string, config *Config) string {
	packageManager := ""

	// Try to detect from engines field first
	if data.Engines != nil {
		switch {
		case data.Engines.NPM != "":
			packageManager = config.NPM
		case data.Engines.PNPM != "":
			packageManager = config.PNPM
		case data.Engines.Yarn != "":
			packageManager = config.Yarn
		case data.Engines.Bun != "":
			packageManager = config.Bun
		}
	}

	// If no package manager found in engines, try lockfile detection
	if packageManager == "" {
		packageManager = DetectFromLockfile(dir, config)
	}

	// If still no package manager, fall back to npm
	if packageManager == "" {
		fmt.Println("No package manager specified, falling back to npm")
		packageManager = config.NPM
	}

	return packageManager
}
