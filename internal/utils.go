package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

func Pointer[T any](v T) *T {
	return &v
}

// WriteFile writes data to a file with the given name and permissions.
// The file is written to a temporary file in the same directory as the target file
// and then renamed to the target file to avoid partial writes in case of a crash.
func WriteFile(name string, data []byte, perm fs.FileMode) error {
	dir := filepath.Dir(name)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpFile, err := os.CreateTemp(dir, "tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	if err := tmpFile.Chmod(perm); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed to set permissions on temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	if err := os.Rename(tmpFile.Name(), name); err != nil {
		return fmt.Errorf("failed to rename temp file to target file: %w", err)
	}

	return nil
}

// GetFileMatches returns the path of the first file in a dir matching the regex or "" if no file
// match was found.
func GetFileMatches(path string, re *regexp.Regexp) ([]string, error) {
	var matches []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return matches, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		match := re.FindString(entry.Name())
		if match == "" {
			continue
		}
		matches = append(matches, match)
	}
	return matches, nil
}

// Serializes a map of service arguments in the format "argument=value" to file.
func SerializeArgumentFile(arguments map[string]string, path string, headerComment string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to write argument file %s: %w", path, err)
	}
	defer file.Close()

	if headerComment != "" {
		file.WriteString(headerComment)
	}

	// Order the argument keys alphabetically to make the output deterministic
	keys := make([]string, 0)
	for k := range arguments {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		file.WriteString(fmt.Sprintf("%s=%s\n", k, arguments[k]))
	}

	return nil
}
