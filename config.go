package main

type Engines struct {
	NPM  string `json:"npm,omitempty"`
	PNPM string `json:"pnpm,omitempty"`
	Yarn string `json:"yarn,omitempty"`
	Bun  string `json:"bun,omitempty"`
}

type PackageJson struct {
	Engines *Engines `json:"engines"`
	Name    string   `json:"name"`
}

type Config struct {
	PackageJSONName string
	RootDir         string
	NPM             string
	NPMRun          string
	PNPM            string
	Yarn            string
	Bun             string
	GitIgnore       string
}

func InitializeConfig() *Config {

	return &Config{
		PackageJSONName: "package.json",
		RootDir:         ".",
		NPM:             "npm",
		NPMRun:          "run",
		PNPM:            "pnpm",
		Yarn:            "yarn",
		Bun:             "bun",
		GitIgnore:       ".gitignore",
	}
}
