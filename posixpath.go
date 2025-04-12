package commonpath

import (
	"fmt"
	"strings"
)

// CommonPathUnix returns the longest common sub-path given a sequence of path names.
func CommonPathUnix(paths []string) (string, error) {
	n := len(paths)

	if n == 0 {
		return "", fmt.Errorf("CommonPathUnix() arg is an empty sequence")
	}

	const sep = "/"
	const curdir = "."

	isAbs := strings.HasPrefix(paths[0], sep)
	for _, p := range paths {
		if strings.HasPrefix(p, sep) != isAbs {
			return "", fmt.Errorf("can't mix absolute and relative paths")
		}
	}

	splitPaths := make([][]string, n)
	for i, p := range paths {
		parts := strings.Split(p, sep)
		filtered := make([]string, 0, len(parts))
		for _, part := range parts {
			if part != "" && part != curdir {
				filtered = append(filtered, part)
			}
		}
		splitPaths[i] = filtered
	}

	s1, s2 := minPath(splitPaths), maxPath(splitPaths)

	common := s1
	for i := range s1 {
		if i >= len(s2) || s1[i] != s2[i] {
			common = s1[:i]
			break
		}
	}

	prefix := ""
	if isAbs {
		prefix = sep
	}
	return prefix + strings.Join(common, sep), nil
}

func minPath(paths [][]string) []string {
	min := paths[0]
	for _, p := range paths[1:] {
		if lessPath(p, min) {
			min = p
		}
	}
	return min
}

func maxPath(paths [][]string) []string {
	max := paths[0]
	for _, p := range paths[1:] {
		if lessPath(max, p) {
			max = p
		}
	}
	return max
}

func lessPath(a, b []string) bool {
	aN := len(a)
	bN := len(b)
	n := min(bN, aN)
	for i := range n {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return aN < bN
}
