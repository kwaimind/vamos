package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSelect_WithNPM(t *testing.T) {
	data := &PackageJson{
		Engines: &Engines{
			NPM: "10.9.2",
		},
	}

	result := Select(data, ".")

	if result.PackageManager != "npm" {
		t.Errorf("Expected 'npm', got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "engines.npm in package.json" {
		t.Errorf("Expected 'engines.npm in package.json', got '%s'", result.DetectionMethod)
	}
}

func TestSelect_WithPNPM(t *testing.T) {
	
	data := &PackageJson{
		Engines: &Engines{
			PNPM: ">=8.9.0",
		},
	}

	result := Select(data, ".")

	if result.PackageManager != "pnpm" {
		t.Errorf("Expected 'pnpm', got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "engines.pnpm in package.json" {
		t.Errorf("Expected 'engines.pnpm in package.json', got '%s'", result.DetectionMethod)
	}
}

func TestSelect_WithYarn(t *testing.T) {
	
	data := &PackageJson{
		Engines: &Engines{
			Yarn: "1.22.0",
		},
	}

	result := Select(data, ".")

	if result.PackageManager != "yarn" {
		t.Errorf("Expected 'yarn', got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "engines.yarn in package.json" {
		t.Errorf("Expected 'engines.yarn in package.json', got '%s'", result.DetectionMethod)
	}
}

func TestSelect_NoEngines(t *testing.T) {
	
	data := &PackageJson{}

	result := Select(data, ".")

	if result.PackageManager != "npm" {
		t.Errorf("Expected default 'npm', got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "default (npm)" {
		t.Errorf("Expected 'default (npm)', got '%s'", result.DetectionMethod)
	}
}

func TestSelect_EmptyEngines(t *testing.T) {
	
	data := &PackageJson{
		Engines: &Engines{},
	}

	result := Select(data, ".")

	if result.PackageManager != "npm" {
		t.Errorf("Expected default 'npm', got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "default (npm)" {
		t.Errorf("Expected 'default (npm)', got '%s'", result.DetectionMethod)
	}
}

func TestSelect_PriorityOrder(t *testing.T) {
	
	// NPM has highest priority in switch statement
	data := &PackageJson{
		Engines: &Engines{
			NPM:  "10.9.2",
			PNPM: ">=8.9.0",
			Yarn: "1.22.0",
		},
	}

	result := Select(data, ".")

	if result.PackageManager != "npm" {
		t.Errorf("Expected 'npm' (highest priority), got '%s'", result.PackageManager)
	}
}

func TestSelect_PNPMOverYarn(t *testing.T) {
	
	// PNPM has priority over Yarn
	data := &PackageJson{
		Engines: &Engines{
			PNPM: ">=8.9.0",
			Yarn: "1.22.0",
		},
	}

	result := Select(data, ".")

	if result.PackageManager != "pnpm" {
		t.Errorf("Expected 'pnpm' (priority over yarn), got '%s'", result.PackageManager)
	}
}

func TestSelect_WithBun(t *testing.T) {
	
	data := &PackageJson{
		Engines: &Engines{
			Bun: "1.0.0",
		},
	}

	result := Select(data, ".")

	if result.PackageManager != "bun" {
		t.Errorf("Expected 'bun', got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "engines.bun in package.json" {
		t.Errorf("Expected 'engines.bun in package.json', got '%s'", result.DetectionMethod)
	}
}

func TestDetectFromLockfile_BunLockb(t *testing.T) {
	
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "bun.lockb")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir)

	if result != "bun" {
		t.Errorf("Expected 'bun', got '%s'", result)
	}
}

func TestDetectFromLockfile_PnpmLock(t *testing.T) {
	
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "pnpm-lock.yaml")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm', got '%s'", result)
	}
}

func TestDetectFromLockfile_YarnLock(t *testing.T) {
	
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "yarn.lock")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir)

	if result != "yarn" {
		t.Errorf("Expected 'yarn', got '%s'", result)
	}
}

func TestDetectFromLockfile_PackageLock(t *testing.T) {
	
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "package-lock.json")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir)

	if result != "npm" {
		t.Errorf("Expected 'npm', got '%s'", result)
	}
}

func TestDetectFromLockfile_NoLockfile(t *testing.T) {
	
	tmpDir := t.TempDir()

	result := DetectFromLockfile(tmpDir)

	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestSelect_LockfileFallback(t *testing.T) {
	
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "pnpm-lock.yaml")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	data := &PackageJson{}

	result := Select(data, tmpDir)

	if result.PackageManager != "pnpm" {
		t.Errorf("Expected 'pnpm' from lockfile detection, got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "pnpm-lock.yaml" {
		t.Errorf("Expected 'pnpm-lock.yaml', got '%s'", result.DetectionMethod)
	}
}

func TestSelect_EnginesPriorityOverLockfile(t *testing.T) {
	
	tmpDir := t.TempDir()
	// Create a yarn.lock file
	lockfilePath := filepath.Join(tmpDir, "yarn.lock")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	// But engines field specifies npm
	data := &PackageJson{
		Engines: &Engines{
			NPM: "10.0.0",
		},
	}

	result := Select(data, tmpDir)

	// Should use engines field (npm), not lockfile (yarn)
	if result.PackageManager != "npm" {
		t.Errorf("Expected 'npm' from engines (priority over lockfile), got '%s'", result.PackageManager)
	}
	if result.DetectionMethod != "engines.npm in package.json" {
		t.Errorf("Expected 'engines.npm in package.json', got '%s'", result.DetectionMethod)
	}
}
