package chezmoi

import (
	"io/fs"
	"time"
)

// A SourceStateEntryType is the type of a source state entry.
type SourceStateEntryType int

const (
	SourceStateEntryTypeAbsent SourceStateEntryType = iota
	SourceStateEntryTypeDir
	SourceStateEntryTypeFile
	SourceStateEntryTypeSymlink
)

// A ChezmoiStat contains information about a file or directory, similar to
// fs.FileInfo but with additional chezmoi-specific fields.
type ChezmoiStat struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	isDir   bool
}

// NewChezmoiStat returns a new ChezmoiStat from an fs.FileInfo.
// Returns a zero-value ChezmoiStat if info is nil, which will have
// SourceStateEntryTypeAbsent as its Type().
func NewChezmoiStat(info fs.FileInfo) ChezmoiStat {
	if info == nil {
		return ChezmoiStat{}
	}
	return ChezmoiStat{
		name:    info.Name(),
		size:    info.Size(),
		mode:    info.Mode(),
		modTime: info.ModTime(),
		isDir:   info.IsDir(),
	}
}

// Name returns the name of the file.
func (s ChezmoiStat) Name() string {
	return s.name
}

// Size returns the size of the file in bytes.
func (s ChezmoiStat) Size() int64 {
	return s.size
}

// Mode returns the file mode bits.
func (s ChezmoiStat) Mode() fs.FileMode {
	return s.mode
}

// ModTime returns the modification time of the file.
func (s ChezmoiStat) ModTime() time.Time {
	return s.modTime
}

// IsDir reports whether the file is a directory.
func (s ChezmoiStat) IsDir() bool {
	return s.isDir
}

// IsRegular reports whether the file is a regular file.
func (s ChezmoiStat) IsRegular() bool {
	return s.mode.IsRegular()
}

// IsSymlink reports whether the file is a symbolic link.
func (s ChezmoiStat) IsSymlink() bool {
	return s.mode&fs.ModeSymlink != 0
}

// Type returns the SourceStateEntryType for this stat.
// Note: isDir is checked first because on some systems a directory entry
// may also have other mode bits set that could cause ambiguity.
// Note: device files, named pipes, and other special files fall through to
// Absent since chezmoi does not manage them.
func (s ChezmoiStat) Type() SourceStateEntryType {
	switch {
	case s.isDir:
		return SourceStateEntryTypeDir
	case s.mode&fs.ModeSymlink != 0:
		return SourceStateEntryTypeSymlink
	case s.mode.IsRegular():
		return SourceStateEntryTypeFile
	default:
		return SourceStateEntryTypeAbsent
	}
}

// Perm returns the Unix permission bits of the file.
func (s ChezmoiStat) Perm() fs.FileMode {
	return s.mode.Perm()
}
