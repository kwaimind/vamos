package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

var version = "0.4.0"

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
	verbose := cmd.Bool("verbose")
	filter := cmd.String("filter")
	args := cmd.Args().Slice()


	if len(args) == 0 {
		fmt.Println("❌ please provide a command to run")
		return nil
	}

	fmt.Println("⚡️ vamos...")

	packagePath := packageJSONName

	// Filter logic
	if filter != "" {
		if verbose {
			fmt.Printf("🔍 Searching for package: %s\n", filter)
		}
		ignoreFiles, err := SetupIgnore(gitIgnore)
		if err != nil {
			// If .gitignore doesn't exist, just use empty ignore list
			ignoreFiles = []string{}
		}
		nextPackagePath, err := FindPackageJson(rootDir, filter, ignoreFiles)
		if err != nil {
			reason := fmt.Sprintf("❌ Error searching for package: %s", err.Error())
			fmt.Println(reason)
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
		fmt.Println(reason)
		return nil
	}
	defer file.Close()

	data, err := ParseJson(file)
	if err != nil {
		t := err.Error()
		fmt.Println(t)
		reason := fmt.Sprintf("👆 This is probably related to your %s script and not vamos.", "packageManager")
		fmt.Println(reason)
		return nil
	}

	// Detect package manager
	packageDir := filepath.Dir(packagePath)
	selection := Select(data, packageDir)

	if verbose {
		fmt.Printf("🚀 Using %s (detected from %s)\n", selection.PackageManager, selection.DetectionMethod)
	}

	// Build command
	nextargs := args

	// Strip "run" if it's the first argument (e.g., "vamos run test" -> "test")
	if len(args) > 1 && args[0] == npmRun {
		nextargs = args[1:]
	}

	isProtected := false
	for _, cmd := range protectedCommands {
		if nextargs[0] == cmd {
			isProtected = true
			break
		}
	}

	// For npm, add "run" prefix unless it's a protected command
	if selection.PackageManager == npm && !isProtected {
		nextargs = append([]string{npmRun}, nextargs...)
	}

	// Execute
	execCmd := exec.Command(selection.PackageManager, nextargs...)
	execCmd.Dir = packageDir
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	cmdErr := execCmd.Run()
	if cmdErr != nil {
		reason := fmt.Sprintf("👆 This is probably related to your %s script and not vamos.", selection.PackageManager)
		fmt.Println(reason)
	}

	return nil
}
