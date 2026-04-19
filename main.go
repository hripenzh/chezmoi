// chezmoi - A dotfile manager for multiple machines.
// Fork of twpayne/chezmoi with additional enhancements.
package main

import (
	"os"

	"github.com/twpayne/chezmoi/v2/internal/cmd"
)

func main() {
	if err := cmd.Main(cmd.VersionInfo{
		Version: "dev",
		Commit:  "none",
		Date:    "unknown",
		BuiltBy: "source",
	}, os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
