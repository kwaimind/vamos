package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSelect_WithNPM(t *testing.T) {
	config := InitializeConfig()
	data := &PackageJson{
		Engines: &Engines{
			NPM: "10.9.2",
		},
	}

	result := Select(data, ".", config)

	if result != "npm" {
		t.Errorf("Expected 'npm', got '%s'", result)
	}
}

func TestSelect_WithPNPM(t *testing.T) {
	config := InitializeConfig()
	data := &PackageJson{
		Engines: &Engines{
			PNPM: ">=8.9.0",
		},
	}

	result := Select(data, ".", config)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm', got '%s'", result)
	}
}

func TestSelect_WithYarn(t *testing.T) {
	config := InitializeConfig()
	data := &PackageJson{
		Engines: &Engines{
			Yarn: "1.22.0",
		},
	}

	result := Select(data, ".", config)

	if result != "yarn" {
		t.Errorf("Expected 'yarn', got '%s'", result)
	}
}

func TestSelect_NoEngines(t *testing.T) {
	config := InitializeConfig()
	data := &PackageJson{}

	result := Select(data, ".", config)

	if result != "npm" {
		t.Errorf("Expected default 'npm', got '%s'", result)
	}
}

func TestSelect_EmptyEngines(t *testing.T) {
	config := InitializeConfig()
	data := &PackageJson{
		Engines: &Engines{},
	}

	result := Select(data, ".", config)

	if result != "npm" {
		t.Errorf("Expected default 'npm', got '%s'", result)
	}
}

func TestSelect_PriorityOrder(t *testing.T) {
	config := InitializeConfig()
	// NPM has highest priority in switch statement
	data := &PackageJson{
		Engines: &Engines{
			NPM:  "10.9.2",
			PNPM: ">=8.9.0",
			Yarn: "1.22.0",
		},
	}

	result := Select(data, ".", config)

	if result != "npm" {
		t.Errorf("Expected 'npm' (highest priority), got '%s'", result)
	}
}

func TestSelect_PNPMOverYarn(t *testing.T) {
	config := InitializeConfig()
	// PNPM has priority over Yarn
	data := &PackageJson{
		Engines: &Engines{
			PNPM: ">=8.9.0",
			Yarn: "1.22.0",
		},
	}

	result := Select(data, ".", config)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm' (priority over yarn), got '%s'", result)
	}
}

func TestSelect_WithBun(t *testing.T) {
	config := InitializeConfig()
	data := &PackageJson{
		Engines: &Engines{
			Bun: "1.0.0",
		},
	}

	result := Select(data, ".", config)

	if result != "bun" {
		t.Errorf("Expected 'bun', got '%s'", result)
	}
}

func TestDetectFromLockfile_BunLockb(t *testing.T) {
	config := InitializeConfig()
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "bun.lockb")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir, config)

	if result != "bun" {
		t.Errorf("Expected 'bun', got '%s'", result)
	}
}

func TestDetectFromLockfile_PnpmLock(t *testing.T) {
	config := InitializeConfig()
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "pnpm-lock.yaml")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir, config)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm', got '%s'", result)
	}
}

func TestDetectFromLockfile_YarnLock(t *testing.T) {
	config := InitializeConfig()
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "yarn.lock")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir, config)

	if result != "yarn" {
		t.Errorf("Expected 'yarn', got '%s'", result)
	}
}

func TestDetectFromLockfile_PackageLock(t *testing.T) {
	config := InitializeConfig()
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "package-lock.json")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	result := DetectFromLockfile(tmpDir, config)

	if result != "npm" {
		t.Errorf("Expected 'npm', got '%s'", result)
	}
}

func TestDetectFromLockfile_NoLockfile(t *testing.T) {
	config := InitializeConfig()
	tmpDir := t.TempDir()

	result := DetectFromLockfile(tmpDir, config)

	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestSelect_LockfileFallback(t *testing.T) {
	config := InitializeConfig()
	tmpDir := t.TempDir()
	lockfilePath := filepath.Join(tmpDir, "pnpm-lock.yaml")
	if err := os.WriteFile(lockfilePath, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	data := &PackageJson{}

	result := Select(data, tmpDir, config)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm' from lockfile detection, got '%s'", result)
	}
}

func TestSelect_EnginesPriorityOverLockfile(t *testing.T) {
	config := InitializeConfig()
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

	result := Select(data, tmpDir, config)

	// Should use engines field (npm), not lockfile (yarn)
	if result != "npm" {
		t.Errorf("Expected 'npm' from engines (priority over lockfile), got '%s'", result)
	}
}
