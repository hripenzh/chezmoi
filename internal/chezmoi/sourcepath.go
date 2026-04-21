package chezmoi

import (
	"path/filepath"
	"strings"
)

// SourceRelPath represents a relative path within the source directory.
type SourceRelPath struct {
	relPath RelPath
	isDir   bool
}

// RelPath represents a relative path.
type RelPath struct {
	path string
}

// NewRelPath returns a new RelPath for the given path string.
func NewRelPath(path string) RelPath {
	return RelPath{path: filepath.ToSlash(path)}
}

// String returns the string representation of the RelPath.
func (r RelPath) String() string {
	return r.path
}

// Base returns the last element of the path.
func (r RelPath) Base() string {
	return filepath.Base(r.path)
}

// Dir returns all but the last element of the path.
func (r RelPath) Dir() RelPath {
	return NewRelPath(filepath.Dir(r.path))
}

// Join returns a new RelPath with the given elements appended.
func (r RelPath) Join(elems ...RelPath) RelPath {
	strs := make([]string, 0, len(elems)+1)
	strs = append(strs, r.path)
	for _, e := range elems {
		strs = append(strs, e.path)
	}
	return NewRelPath(filepath.Join(strs...))
}

// HasPrefix returns true if the path starts with the given prefix followed by
// a path separator, or if the path equals the prefix exactly. This ensures
// we don't get false positives when one path is a prefix of another path's
// component name (e.g. "foo/bar" should not HasPrefix "foo/b").
func (r RelPath) HasPrefix(prefix RelPath) bool {
	if prefix.path == "" {
		return true
	}
	if r.path == prefix.path {
		return true
	}
	return strings.HasPrefix(r.path, prefix.path+"/")
}

// Empty returns true if the path is empty.
func (r RelPath) Empty() bool {
	return r.path == ""
}

// TrimPrefix returns a new RelPath with the given prefix removed.
// If the path does not have the prefix, the original path is returned unchanged.
func (r RelPath) TrimPrefix(prefix RelPath) RelPath {
	if !r.HasPrefix(prefix) {
		return r
	}
	if r.path == prefix.path {
		return NewRelPath("")
	}
	return NewRelPath(strings.TrimPrefix(r.path, prefix.path+"/"))
}

// NewSourceRelPath returns a new SourceRelPath.
func NewSourceRelPath(relPath RelPath, isDir bool) SourceRelPath {
	return SourceRelPath{
		relPath: relPath,
		isDir:   isDir,
	}
}

// RelPath returns the underlying RelPath.
func (s SourceRelPath) RelPath() RelPath {
	return s.relPath
}

// IsDir returns true if this source path represents a directory.
func (s SourceRelPath) IsDir() bool {
	return s.isDir
}

// String returns the string representation of the SourceRelPath.
func (s SourceRelPath) String() string {
	return s.relPath.String()
}

// AbsPath represents an absolute path.
type AbsPath struct {
	path string
}

// NewAbsPath returns a new AbsPath for the given path string.
func NewAbsPath(path string) AbsPath {
	return AbsPath{path: filepath.ToSlash(filepath.Clean(path))}
}

// String returns the string representation of the AbsPath.
func (a AbsPath) String() string {
	return a.path
}

// Base returns the last element of the path.
func (a AbsPath) Base() string {
	return filepath.Base(a.path)
}

// Dir returns all but the last element of the path.
func (a AbsPath) Dir() AbsPath {
	return NewAbsPath(filepath.Dir(a.path))
}

// Join appends the given RelPath elements to the AbsPath.
func (a AbsPath) Join(relPaths ...RelPath) AbsPath {
	strs := make([]string, 0, len(relPaths)+1)
	strs = append(strs, a.path)
	for _, r := range relPaths {
		strs = append(strs, r.path)
	}
	return NewAbsPath(filepath.Join(strs...))
}
