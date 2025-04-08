package main

import (
	"testing"
)

func TestPickFilter_WithFlag(t *testing.T) {
	args := []string{"-f", "filter-value", "other", "args"}
	result := PickFilter(args)

	if result != "filter-value" {
		t.Errorf("Expected 'filter-value', got '%s'", result)
	}
}

func TestPickFilter_WithoutFlag(t *testing.T) {
	args := []string{"--some", "other", "args"}
	result := PickFilter(args)

	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestPickFilter_FlagWithoutValue(t *testing.T) {
	args := []string{"--other", "value", "-f"}
	result := PickFilter(args)

	if result != "" {
		t.Errorf("Expected empty string when -f has no value, got '%s'", result)
	}
}
