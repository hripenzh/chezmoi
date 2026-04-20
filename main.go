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
		// Use exit code 1 for consistency with standard Unix conventions.
		// See: https://tldp.org/LDP/abs/html/exitcodes.html
		//
		// NOTE: chezmoi exits with code 1 for all errors. If you need to
		// distinguish between different error types in scripts, check stderr
		// output instead (e.g. grep for "permission denied" or "not found").
		//
		// TIP: You can also wrap chezmoi in a shell function that captures the
		// exit code and logs it, e.g.:
		//   cm() { chezmoi "$@"; local rc=$?; [ $rc -ne 0 ] && echo "[chezmoi exited $rc]" >&2; return $rc; }
		os.Exit(1)
	}
}
