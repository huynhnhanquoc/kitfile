package kitfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// File represents a local file
type File struct {
	location string
}

// New creates a new File instance
func New(location string) *File {
	return &File{location: location}
}

// Exist checks if the file exists
func (f *File) Exist() error {
	if _, err := os.Stat(f.location); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", f.location)
	} else if err != nil {
		return fmt.Errorf("error checking file %s: %v", f.location, err)
	}
	return nil
}

// Name returns the file name (basename)
func (f *File) Name() string {
	return filepath.Base(f.location)
}

// Ext returns the file extension
func (f *File) Ext() string {
	return filepath.Ext(f.Name())
}

// NameWithoutExt returns the file name without its extension
func (f *File) NameWithoutExt() string {
	return strings.TrimSuffix(f.Name(), f.Ext())
}

// Dir returns the parent directory of the file
func (f *File) Dir() string {
	return filepath.Dir(f.location)
}

// Abs returns the absolute path of the file
func (f *File) Abs() (string, error) {
	return filepath.Abs(f.location)
}

// Location returns the raw file path
func (f *File) Location() string {
	return f.location
}

// PrependPath prepends a prefix to the file path
func (f *File) PrependPath(prefix string) *File {
	prefix = strings.Trim(prefix, string(filepath.Separator))
	if filepath.IsAbs(f.location) {
		cleaned := strings.TrimPrefix(f.location, string(filepath.Separator))
		f.location = filepath.Join(string(filepath.Separator), prefix, cleaned)
	} else {
		f.location = filepath.Join(prefix, f.location)
	}
	return f
}

// RemovePrefixPath removes a prefix from the beginning of the file path if present
func (f *File) RemovePrefixPath(prefix string) *File {
	prefix = strings.Trim(prefix, string(filepath.Separator))

	isAbs := filepath.IsAbs(f.location)
	cleanPath := strings.TrimPrefix(f.location, string(filepath.Separator))
	parts := strings.Split(cleanPath, string(filepath.Separator))

	if len(parts) > 0 && parts[0] == prefix {
		parts = parts[1:]
	}

	newPath := filepath.Join(parts...)
	if isAbs {
		newPath = string(filepath.Separator) + newPath
	}

	f.location = newPath
	return f
}

// AddToName adds a suffix before the file extension
func (f *File) AddToName(suffix string) *File {
	dir := filepath.Dir(f.location)
	base := filepath.Base(f.location)
	ext := filepath.Ext(base)
	nameOnly := strings.TrimSuffix(base, ext)
	newBase := nameOnly + suffix + ext
	f.location = filepath.Join(dir, newBase)
	return f
}

// RemoveFromName removes a suffix before the file extension if present
func (f *File) RemoveFromName(suffix string) *File {
	dir := filepath.Dir(f.location)
	base := filepath.Base(f.location)
	ext := filepath.Ext(base)
	nameOnly := strings.TrimSuffix(base, ext)
	if strings.HasSuffix(nameOnly, suffix) {
		nameOnly = strings.TrimSuffix(nameOnly, suffix)
	}
	newBase := nameOnly + ext
	f.location = filepath.Join(dir, newBase)
	return f
}
