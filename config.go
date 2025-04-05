package main

type Engines struct {
	NPM  string `json:"npm,omitempty"`
	PNPM string `json:"pnpm,omitempty"`
	Yarn string `json:"yarn,omitempty"`
}

type PackageJson struct {
	Engines *Engines `json:"engines"`
	Name    string   `json:"name"`
}

type Config struct {
	PackagejsonName string
	RootDir         string
	NPM             string
	NPMRun          string
	PNPM            string
	Yarn            string
}

func InitializeConfig() *Config {

	return &Config{
		PackagejsonName: "package.json",
		RootDir:         ".",
		NPM:             "npm",
		NPMRun:          "run",
		PNPM:            "pnpm",
		Yarn:            "yarn",
	}
}
