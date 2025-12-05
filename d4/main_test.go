package main

import "testing"

func TestValidatePosition(t *testing.T) {
	tests := []struct {
		lines       [][]string
		x, y        int
		maxAdjacent int
		isValid     bool
	}{
		{
			lines: [][]string{
				{".", "@", "."},
				{"@", "@", "."},
				{".", "@", "."},
			},
			x:           1,
			y:           1,
			maxAdjacent: 4,
			isValid:     true,
		},
		{
			lines: [][]string{
				{".", "@", "."},
				{".", "@", "@"},
				{".", "@", "."},
			},
			x:           1,
			y:           1,
			maxAdjacent: 4,
			isValid:     true,
		},
		{
			lines: [][]string{
				{".", "@", "."},
				{"@", "@", "@"},
				{".", "@", "."},
			},
			x:           1,
			y:           1,
			maxAdjacent: 4,
			isValid:     false,
		},
	}
	for _, test := range tests {
		result := validatePosition(test.lines, test.x, test.y, test.maxAdjacent)
		if result != test.isValid {
			t.Errorf("validatePosition(%v, %d, %d, %d) = %v; want %v", test.lines, test.x, test.y, test.maxAdjacent, result, test.isValid)
		}
	}

}
