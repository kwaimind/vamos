package main

import (
	"testing"
)

func TestSelect_WithNPM(t *testing.T) {
	data := &PackageJson{
		Engines: &Engines{
			NPM: "10.9.2",
		},
	}

	result := Select(data)

	if result != "npm" {
		t.Errorf("Expected 'npm', got '%s'", result)
	}
}

func TestSelect_WithPNPM(t *testing.T) {
	data := &PackageJson{
		Engines: &Engines{
			PNPM: ">=8.9.0",
		},
	}

	result := Select(data)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm', got '%s'", result)
	}
}

func TestSelect_WithYarn(t *testing.T) {
	data := &PackageJson{
		Engines: &Engines{
			Yarn: "1.22.0",
		},
	}

	result := Select(data)

	if result != "yarn" {
		t.Errorf("Expected 'yarn', got '%s'", result)
	}
}

func TestSelect_NoEngines(t *testing.T) {
	data := &PackageJson{}

	result := Select(data)

	if result != "npm" {
		t.Errorf("Expected default 'npm', got '%s'", result)
	}
}

func TestSelect_EmptyEngines(t *testing.T) {
	data := &PackageJson{
		Engines: &Engines{},
	}

	result := Select(data)

	if result != "npm" {
		t.Errorf("Expected default 'npm', got '%s'", result)
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

	result := Select(data)

	if result != "npm" {
		t.Errorf("Expected 'npm' (highest priority), got '%s'", result)
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

	result := Select(data)

	if result != "pnpm" {
		t.Errorf("Expected 'pnpm' (priority over yarn), got '%s'", result)
	}
}
