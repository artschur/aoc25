package main

import (
	"testing"
)

func TestD1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		len      int
	}{
		{"123456", 56, 2},
		{"9876543210", 98, 2},
		{"111222333", 33, 2},
		{"000111222", 22, 2},
		{"543216789", 89, 2},
	}
	for _, test := range tests {
		result := getMaximumJoltage(test.input, test.len)
		if result != test.expected {
			t.Errorf("getLargestJoltage(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestD2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		len      int
	}{
		{"987654321111111", 987654321111, 12},
		{"811111111111119", 811111111119, 12},
	}
	for _, test := range tests {
		result := getMaximumJoltage(test.input, test.len)
		if result != test.expected {
			t.Errorf("getLargestJoltage(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
