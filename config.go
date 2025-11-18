package main

const (
	packageJSONName = "package.json"
	rootDir         = "."
	npm             = "npm"
	npmRun          = "run"
	pnpm            = "pnpm"
	yarn            = "yarn"
	bun             = "bun"
	gitIgnore       = ".gitignore"
)

var protectedCommands = []string{"install", "i"}

type Engines struct {
	NPM  string `json:"npm,omitempty"`
	PNPM string `json:"pnpm,omitempty"`
	Yarn string `json:"yarn,omitempty"`
	Bun  string `json:"bun,omitempty"`
}

type PackageJson struct {
	Engines *Engines          `json:"engines"`
	Name    string            `json:"name"`
	Scripts map[string]string `json:"scripts,omitempty"`
}
