package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"
)

var version = "0.1.0"

func main() {
	stopOn := 1
	cmd := &cli.Command{
		Name:        "vamos",
		Usage:       "Smart package manager runner for monorepos",
		UsageText:   "vamos [options] <command> [args...]",
		Version:     version,
		StopOnNthArg: &stopOn,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "filter",
				Aliases: []string{"f"},
				Usage:   "Filter by package name in monorepo",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "Show which package manager was selected and why",
			},
		},
		Action: run,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, cmd *cli.Command) error {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4"))

	verbose := cmd.Bool("verbose")
	filter := cmd.String("filter")
	args := cmd.Args().Slice()


	if len(args) == 0 {
		fmt.Println(style.Render("❌ please provide a command to run"))
		return nil
	}

	fmt.Println(style.Render("⚡️ vamos..."))

	config := InitializeConfig()
	packagePath := config.PackageJSONName

	// Filter logic
	if filter != "" {
		if verbose {
			fmt.Printf("🔍 Searching for package: %s\n", filter)
		}
		ignoreFiles, err := SetupIgnore(config.GitIgnore)
		if err != nil {
			// If .gitignore doesn't exist, just use empty ignore list
			ignoreFiles = []string{}
		}
		nextPackagePath, err := FindPackageJson(config.RootDir, filter, ignoreFiles, config)
		if err != nil {
			reason := fmt.Sprintf("❌ Error searching for package: %s", err.Error())
			fmt.Println(style.Render(reason))
			return nil
		}
		packagePath = nextPackagePath
		if verbose {
			fmt.Printf("📦 Found package at: %s\n", packagePath)
		}
	}

	// Parse package.json
	file, err := os.Open(packagePath)
	if err != nil {
		rootError := err.Error()
		errMsg := strings.Split(rootError, " : ")[1]
		reason := fmt.Sprintf("❌ Had problems finding a package.json. %s.", CapitalizeFirst(errMsg))
		fmt.Println(style.Render(reason))
		return nil
	}
	defer file.Close()

	data, err := ParseJson(file)
	if err != nil {
		t := err.Error()
		fmt.Println(t)
		reason := fmt.Sprintf("👆 This is probably related to your %s script and not vamos.", "packageManager")
		fmt.Println(style.Render(reason))
		return nil
	}


	// Detect package manager
	packageDir := filepath.Dir(packagePath)
	packageManager := Select(data, packageDir, config)

	if verbose {
		// Determine detection method
		detectionMethod := "default (npm)"
		if data.Engines != nil {
			switch {
			case data.Engines.NPM != "":
				detectionMethod = "engines.npm in package.json"
			case data.Engines.PNPM != "":
				detectionMethod = "engines.pnpm in package.json"
			case data.Engines.Yarn != "":
				detectionMethod = "engines.yarn in package.json"
			case data.Engines.Bun != "":
				detectionMethod = "engines.bun in package.json"
			}
		} else {
			// Check if lockfile detection was used
			lockfileDetected := DetectFromLockfile(packageDir, config)
			if lockfileDetected != "" {
				switch lockfileDetected {
				case config.Bun:
					detectionMethod = "bun.lockb"
				case config.PNPM:
					detectionMethod = "pnpm-lock.yaml"
				case config.Yarn:
					detectionMethod = "yarn.lock"
				case config.NPM:
					detectionMethod = "package-lock.json"
				}
			}
		}
		fmt.Printf("🚀 Using %s (detected from %s)\n", packageManager, detectionMethod)
	}

	// Build command
	nextargs := args
	if packageManager == config.NPM {
		nextargs = append([]string{config.NPMRun}, args...)
	}

	// Execute
	execCmd := exec.Command(packageManager, nextargs...)
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	cmdErr := execCmd.Run()
	if cmdErr != nil {
		reason := fmt.Sprintf("👆 This is probably related to your %s script and not vamos.", packageManager)
		fmt.Println(style.Render(reason))
	}

	return nil
}
