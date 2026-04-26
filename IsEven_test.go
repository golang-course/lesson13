package main

import "testing"

func TestIsEven(t *testing.T) {
	var testCases = []struct {
		description string
		input       int
		want        string
	}{
		{"Positive test on 2", 2, "Yes"},
		{"Positice test on 3", 3, "No"},
		{"Positive test on 0", 0, "Yes"},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.want {
				t.Errorf("Incorrect test, want: %s, but got %s", tt.want, result)
			}
		})
	}
}
