// chezmoi - A dotfile manager for multiple machines.
// Fork of twpayne/chezmoi with additional enhancements.
package main

import (
	"fmt"
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
		fmt.Fprintf(os.Stderr, "chezmoi: error: %v\n", err)
		// Use exit code 2 to distinguish chezmoi errors from other non-zero exits
		os.Exit(2)
	}
}
