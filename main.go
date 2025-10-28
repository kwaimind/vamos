package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4"))

func main() {
	fmt.Println(style.Render("⚡️ vamos..."))

	args := ParseArgs()
	config := InitializeConfig()

	name, filteredArgs := PickFilter(args)

	packagePath := config.PackagejsonName

	if name != "" {
		ignoreFile := config.GitIgnore
		ignoreFiles, err := SetupIgnore(ignoreFile)
		if err != nil {
			// If .gitignore doesn't exist, just use empty ignore list
			ignoreFiles = []string{}
		}
		nextPackagePath := FindPackageJson(config.RootDir, name, ignoreFiles)
		packagePath = nextPackagePath
	}

	file, err := os.Open(packagePath)
	if err != nil {
		rootError := err.Error()
		errMsg := strings.Split(rootError, " : ")[1]
		reason := fmt.Sprintf("❌ Had problems finding a package.json. %s.", CapitalizeFirst(errMsg))
		fmt.Println(style.Render(reason))
		return
	}
	defer file.Close()

	data, err := ParseJson(file)
	if err != nil {
		t := err.Error()
		fmt.Println(t)
		reason := fmt.Sprintf("👆 This is probably related to your %s script and not vamos.", "packageManager")
		fmt.Println(style.Render(reason))
		return
	}

	packageManager := Select(data)

	nextargs := filteredArgs

	if packageManager == config.NPM {
		nextargs = append([]string{config.NPMRun}, filteredArgs...)
	}

	cmd := exec.Command(packageManager, nextargs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr := cmd.Run()
	if cmdErr != nil {
		reason := fmt.Sprintf("👆 This is probably related to your %s script and not vamos.", packageManager)
		fmt.Println(style.Render(reason))
	}
}
