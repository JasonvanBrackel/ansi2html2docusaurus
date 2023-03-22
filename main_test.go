package main

import "testing"

func TestReplaceSingleStyle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "color example",
			input:    `style="color:#55ffff"`,
			expected: `style={{color: '#55ffff'}}`,
		},
		{
			name:     "background-color example",
			input:    `style="background-color:#aa0"`,
			expected: `style={{background-color: '#aa0'}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceSingleStyle(tt.input)
			if result != tt.expected {
				t.Errorf("ReplaceStyle(%q) returned %q, but expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveExtraSemicolon(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "color example",
			input:    `style={{;color: '#55ffff'}}`,
			expected: `style={{color: '#55ffff'}}`,
		},
		{
			name:     "background-color example",
			input:    `style={{;background-color: '#aa0'}}`,
			expected: `style={{background-color: '#aa0'}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveExtraSemicolon(tt.input)
			if result != tt.expected {
				t.Errorf("ReplaceStyle(%q) returned %q, but expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReplaceMultiStyle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "color and background-color example",
			input:    `style="color:#000;background-color:#aa0"`,
			expected: `style={{color: '#000'}} style={{background-color: '#aa0'}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceMultiStyle(tt.input)
			if result != tt.expected {
				t.Errorf("ReplaceStyle(%q) returned %q, but expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
