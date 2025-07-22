package mediasort

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestExtensionCaseInsensitive(t *testing.T) {
	config := Config{
		Extensions: "mp4,mkv",
	}

	fs := &fsSort{
		Config:    config,
		validExts: map[string]bool{},
	}

	for _, e := range strings.Split(config.Extensions, ",") {
		fs.validExts["."+strings.ToLower(e)] = true
	}

	expectedExts := map[string]bool{
		".mp4": true,
		".mkv": true,
	}

	for ext, shouldExist := range expectedExts {
		if fs.validExts[ext] != shouldExist {
			t.Errorf("Expected validExts[%s] to be %v, got %v", ext, shouldExist, fs.validExts[ext])
		}
	}

	unexpectedExts := []string{".MP4", ".Mp4", ".mP4", ".MKV", ".Mkv"}
	for _, ext := range unexpectedExts {
		if fs.validExts[ext] {
			t.Errorf("Expected validExts[%s] to be false, got true", ext)
		}
	}

	testCases := []struct {
		filename string
		expected bool
	}{
		{"movie.MP4", true},
		{"show.mp4", true},
		{"file.MKV", true},
		{"document.TXT", false},
	}

	for _, tc := range testCases {
		ext := strings.ToLower(filepath.Ext(tc.filename))
		accepted := fs.validExts[ext]
		if accepted != tc.expected {
			t.Errorf("File %s: expected %v, got %v (extension: %s)", tc.filename, tc.expected, accepted, ext)
		}
	}
}

func TestExtensionCaseInsensitiveWithMixedCaseConfig(t *testing.T) {
	config := Config{
		Extensions: "MP4,Mkv,AVI",
	}

	fs := &fsSort{
		Config:    config,
		validExts: map[string]bool{},
	}

	for _, e := range strings.Split(config.Extensions, ",") {
		fs.validExts["."+strings.ToLower(e)] = true
	}

	expectedExts := map[string]bool{
		".mp4": true,
		".mkv": true,
		".avi": true,
	}

	for ext, shouldExist := range expectedExts {
		if fs.validExts[ext] != shouldExist {
			t.Errorf("Expected validExts[%s] to be %v, got %v", ext, shouldExist, fs.validExts[ext])
		}
	}

	unexpectedExts := []string{".MP4", ".Mkv", ".AVI"}
	for _, ext := range unexpectedExts {
		if fs.validExts[ext] {
			t.Errorf("Expected validExts[%s] to be false, got true", ext)
		}
	}
}
