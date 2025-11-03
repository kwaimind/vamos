package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type SelectResult struct {
	PackageManager  string
	DetectionMethod string
}

// DetectFromLockfile detects package manager from lockfiles in the given directory
func DetectFromLockfile(dir string) string {
	lockfiles := map[string]string{
		"bun.lockb":         bun,
		"pnpm-lock.yaml":    pnpm,
		"yarn.lock":         yarn,
		"package-lock.json": npm,
	}

	for lockfile, pm := range lockfiles {
		if _, err := os.Stat(filepath.Join(dir, lockfile)); err == nil {
			return pm
		}
	}

	return ""
}

func Select(data *PackageJson, dir string) SelectResult {
	result := SelectResult{}

	// Try to detect from engines field first
	if data.Engines != nil {
		switch {
		case data.Engines.NPM != "":
			result.PackageManager = npm
			result.DetectionMethod = "engines.npm in package.json"
		case data.Engines.PNPM != "":
			result.PackageManager = pnpm
			result.DetectionMethod = "engines.pnpm in package.json"
		case data.Engines.Yarn != "":
			result.PackageManager = yarn
			result.DetectionMethod = "engines.yarn in package.json"
		case data.Engines.Bun != "":
			result.PackageManager = bun
			result.DetectionMethod = "engines.bun in package.json"
		}
	}

	// If no package manager found in engines, try lockfile detection
	if result.PackageManager == "" {
		lockfilePM := DetectFromLockfile(dir)
		if lockfilePM != "" {
			result.PackageManager = lockfilePM
			switch lockfilePM {
			case bun:
				result.DetectionMethod = "bun.lockb"
			case pnpm:
				result.DetectionMethod = "pnpm-lock.yaml"
			case yarn:
				result.DetectionMethod = "yarn.lock"
			case npm:
				result.DetectionMethod = "package-lock.json"
			}
		}
	}

	// If still no package manager, fall back to npm
	if result.PackageManager == "" {
		fmt.Println("No package manager specified, falling back to npm")
		result.PackageManager = npm
		result.DetectionMethod = "default (npm)"
	}

	return result
}
