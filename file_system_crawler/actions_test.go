package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      []string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", []string{""}, 0, false},
		{"FilterExtensionMatch", "testdata/dir.log", []string{".log"}, 0, false},
		{"FilterExtensionNoMatch", "testdata/dir.log", []string{".sh"}, 0, true},
		{"FilterExtensionSizeMatch", "testdata/dir.log", []string{".log"}, 10, false},
		{"FilterExtensionSizeNoMatch", "testdata/dir.log", []string{".log"}, 20, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatal(err)
			}
			f := filterOut(tc.file, tc.ext, tc.minSize, DateTime{} ,info)

			if f != tc.expected {
				t.Errorf("Expected '%t', got '%t' instead\n", tc.expected, f)
			}
		})
	}
}
