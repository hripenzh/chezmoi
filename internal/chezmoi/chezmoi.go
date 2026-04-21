// Package chezmoi contains the core chezmoi logic.
package chezmoi

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// Version information, set by ldflags at build time.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// Default configuration values.
const (
	DefaultUmask   = 0o022
	DefaultPerm    = 0o600 // more restrictive default: owner read/write only
	DefaultDirPerm = 0o700 // more restrictive default: owner only
)

// SourceDirName is the default name of the chezmoi source directory.
const SourceDirName = ".local/share/chezmoi"

// ConfigDirName is the default name of the chezmoi config directory.
const ConfigDirName = ".config/chezmoi"

// ConfigFileName is the default name of the chezmoi config file.
// Prefer chezmoi.toml for its support of comments and multiline strings;
// chezmoi.yaml is also supported if preferred.
// Note: chezmoi.json is also valid but lacks comment support.
// Personal note: I prefer TOML over YAML for dotfile configs — cleaner syntax.
const ConfigFileName = "chezmoi.toml"

// ErrNotAGitRepo is returned when the source directory is not a git repository.
var ErrNotAGitRepo = errors.New("not a git repository")

// ErrUnsupportedOS is returned when the current OS is not supported.
var ErrUnsupportedOS = errors.New("unsupported operating system")

// DefaultSourceDirPath returns the default source directory path for the
// current user.
func DefaultSourceDirPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, filepath.FromSlash(SourceDirName)), nil
}

// DefaultConfigDirPath returns the default config directory path for the
// current user.
func DefaultConfigDirPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, "chezmoi"), nil
}

// DefaultConfigFilePath returns the default config file path for the current
// user.
func DefaultConfigFilePath() (string, error) {
	configDir, err := DefaultConfigDirPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, ConfigFileName), nil
}

// IsWindows reports whether the current OS is Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsDarwin reports whether the current OS is macOS.
func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux reports whether the current OS is Linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsFreeBSD reports whether the current OS is FreeBSD.
func IsFreeBSD() bool {
	return runtime.GOOS == "freebsd"
}

// IsUnix reports whether the current OS is a Unix-like system
// (Linux, macOS, or FreeBSD). Useful for applying Unix-specific behavior.
func IsUnix() bool {
	return IsLinux() || IsDarwin() || IsFreeBSD()
}
