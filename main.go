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
		// Exit with code 1 on error, but print a hint to run with --verbose for more details.
		fmt.Fprintf(os.Stderr, "chezmoi: run with --verbose for more details\n")
		os.Exit(1)
	}
}
