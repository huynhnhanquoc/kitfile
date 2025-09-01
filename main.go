package kitfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// File represents a local file on the filesystem.
type File struct {
	location string
}

// New creates a new File instance with the given location.
// This does not perform validation.
func New(location string) *File {
	return &File{location: location}
}

// NewSafe creates a new File instance with validation and cleaning.
// Returns an error if the location is empty or invalid.
func NewSafe(location string) (*File, error) {
	if location == "" {
		return nil, fmt.Errorf("file location cannot be empty")
	}
	cleaned := filepath.Clean(location)
	if cleaned == "." {
		return nil, fmt.Errorf("invalid file location: %s", location)
	}
	return &File{location: cleaned}, nil
}

// Exist checks if the file exists on disk.
func (f *File) Exist() error {
	if _, err := os.Stat(f.location); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", f.location)
	} else if err != nil {
		return fmt.Errorf("error checking file %s: %v", f.location, err)
	}
	return nil
}

// Name returns the base file name (e.g. "file.txt").
func (f *File) Name() string {
	return filepath.Base(f.location)
}

// Ext returns the file extension (including the dot, e.g. ".txt").
func (f *File) Ext() string {
	return filepath.Ext(f.Name())
}

// NameWithoutExt returns the file name without its extension.
func (f *File) NameWithoutExt() string {
	return strings.TrimSuffix(f.Name(), f.Ext())
}

// Dir returns the parent directory of the file.
func (f *File) Dir() string {
	return filepath.Dir(f.location)
}

// Abs returns the absolute path of the file.
func (f *File) Abs() (string, error) {
	return filepath.Abs(f.location)
}

// Location returns the raw file path.
func (f *File) Location() string {
	return f.location
}

// Read reads the entire content of the file.
// Returns an error if the file does not exist or cannot be read.
func (f *File) Read() ([]byte, error) {
	if err := f.Exist(); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(f.location)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", f.location, err)
	}
	return data, nil
}

// WriteSafe writes the given data into a new file.
// Returns an error if the file already exists (to avoid accidental overwrite).
func (f *File) WriteSafe(data []byte) error {
	if err := f.Exist(); err == nil {
		return fmt.Errorf("file already exists: %s", f.location)
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	file, err := os.OpenFile(f.location, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", f.location, err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("error writing to file %s: %v", f.location, err)
	}
	return nil
}

// Write writes the given data to the file, replacing it if it already exists.
// If the file does not exist, it will be created automatically.
func (f *File) Write(data []byte) error {
	return f.WritePermission(data, 0644)
}

// Write writes the given data to the file, replacing it if it already exists.
// If perm is 0, the default permission 0644 is used.
// Parent directories are created automatically if they do not exist.
func (f *File) WritePermission(data []byte, permission os.FileMode) error {
	if permission == 0 {
		permission = 0644
	}
	dir := filepath.Dir(f.location)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	if err := os.WriteFile(f.location, data, permission); err != nil {
		return fmt.Errorf("error writing to file %s: %v", f.location, err)
	}
	return nil
}

// PrependPath prepends a prefix to the file path.
// Example: "foo/bar.txt" -> PrependPath("root") -> "root/foo/bar.txt".
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

// RemovePrefixPath removes a prefix from the file path if present.
// Example: "root/foo/bar.txt" -> RemovePrefixPath("root") -> "foo/bar.txt".
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

// AddToName appends a suffix before the file extension.
// Example: "file.txt" -> AddToName("_backup") -> "file_backup.txt".
func (f *File) AddToName(suffix string) *File {
	dir := filepath.Dir(f.location)
	base := filepath.Base(f.location)
	ext := filepath.Ext(base)
	nameOnly := strings.TrimSuffix(base, ext)
	newBase := nameOnly + suffix + ext
	f.location = filepath.Join(dir, newBase)
	return f
}

// RemoveFromName removes a given suffix from the filename if present.
// Example: "file_backup.txt" -> RemoveFromName("_backup") -> "file.txt".
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
