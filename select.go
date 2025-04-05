package main

import "fmt"

func Select(data *PackageJson) string {
	config := InitializeConfig()

	packageManager := config.NPM

	if data.Engines != nil {
		switch {
		case data.Engines.NPM != "":
			packageManager = config.NPM
		case data.Engines.PNPM != "":
			packageManager = config.PNPM
		case data.Engines.Yarn != "":
			packageManager = config.Yarn
		}
	}

	if packageManager == "" {
		fmt.Println("No package manager specified, falling back to npm")
	}

	return packageManager
}
