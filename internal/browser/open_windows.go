//go:build windows

package browser

import . "github.com/git-town/git-town/v24/pkg/prelude"

// defaultBrowserCommand provides the console command to open a custom browser on Windows.
func defaultBrowserCommand() Option[string] {
	// The default browser on Windows is opened via rundll32 in open.go.
	// This function is only used when a custom browser executable is configured.
	return None[string]()
}
