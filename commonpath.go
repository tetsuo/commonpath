// Package commonpath provides utilities for computing the longest common sub-path across path sequences.
package commonpath

import "runtime"

// CommonPath returns the longest common sub-path given a sequence of path names.
// Dispatches to the platform-specific implementation.
func CommonPath(paths []string) (string, error) {
	if runtime.GOOS == "windows" {
		return CommonPathWin(paths)
	}
	return CommonPathUnix(paths)
}
