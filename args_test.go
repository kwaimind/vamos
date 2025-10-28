package main

import (
	"os"
	"testing"
)

func TestPickFilter_WithFlag(t *testing.T) {
	args := []string{"-f", "filter-value", "other", "args"}
	result, filteredArgs := PickFilter(args)

	if result != "filter-value" {
		t.Errorf("Expected 'filter-value', got '%s'", result)
	}
	if len(filteredArgs) != 2 || filteredArgs[0] != "other" || filteredArgs[1] != "args" {
		t.Errorf("Expected ['other', 'args'], got %v", filteredArgs)
	}
}

func TestPickFilter_WithoutFlag(t *testing.T) {
	args := []string{"--some", "other", "args"}
	result, filteredArgs := PickFilter(args)

	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
	if len(filteredArgs) != 3 {
		t.Errorf("Expected all args to remain, got %v", filteredArgs)
	}
}

func TestPickFilter_FlagWithoutValue(t *testing.T) {
	args := []string{"--other", "value", "-f"}
	result, filteredArgs := PickFilter(args)

	if result != "" {
		t.Errorf("Expected empty string when -f has no value, got '%s'", result)
	}
	if len(filteredArgs) != 2 || filteredArgs[0] != "--other" || filteredArgs[1] != "value" {
		t.Errorf("Expected ['--other', 'value'], got %v", filteredArgs)
	}
}

func TestParseArgs_WithArgs(t *testing.T) {
	// Save original args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"vamos", "test", "--watch"}

	result, err := ParseArgs()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Expected 2 args, got %d", len(result))
	}
	if result[0] != "test" || result[1] != "--watch" {
		t.Errorf("Expected ['test', '--watch'], got %v", result)
	}
}

func TestParseArgs_NoArgs(t *testing.T) {
	// Save original args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"vamos"}

	result, err := ParseArgs()
	if err == nil {
		t.Error("Expected error when no args provided")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}
