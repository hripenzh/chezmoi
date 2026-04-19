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
