package commonpath

import (
	"errors"
	"strings"
)

// CommonPathWin returns the longest common sub-path given a sequence of path names.
func CommonPathWin(paths []string) (string, error) {
	if len(paths) == 0 {
		return "", errors.New("CommonPathWin() arg is an empty iterable")
	}

	const sep = '\\'
	const altSep = '/'
	const curDir = "."

	components := make([]pathComponents, len(paths))
	for i, p := range paths {
		components[i] = normalize(p, string(altSep), string(sep), curDir)
	}

	drive := components[0].drive
	root := components[0].root

	if drive == "" && strings.HasPrefix(paths[0], `\\`) {
		return "", errors.New("invalid UNC path")
	}

	for _, comp := range components[1:] {
		if comp.drive != drive {
			return "", errors.New("paths don't have the same drive")
		}
		if comp.root != root {
			if drive != "" {
				return "", errors.New("can't mix absolute and relative paths")
			}
			return "", errors.New("can't mix rooted and not-rooted paths")
		}
	}

	commonParts := components[0].parts
	for _, comp := range components[1:] {
		i := 0
		for i < len(commonParts) && i < len(comp.parts) && commonParts[i] == comp.parts[i] {
			i++
		}
		commonParts = commonParts[:i]
	}

	drive = strings.ToLower(drive)

	var result string
	if len(commonParts) > 0 || root != "" {
		result = drive + root + strings.Join(commonParts, string(sep))
	} else {
		result = drive + string(sep)
	}

	if len(commonParts) == 0 && root == "" {
		result = drive + string(sep)
	} else if len(commonParts) == 0 && strings.HasPrefix(drive, `\\`) && root == string(sep) {
		result = drive
	} else {
		result = strings.TrimRight(result, string(sep))
	}

	return result, nil
}

func splitDrive(p string) (string, string) {
	if len(p) >= 2 && p[1] == ':' {
		return p[:2], p[2:]
	}
	if strings.HasPrefix(p, `\\`) {
		idx := strings.Index(p[2:], `\`)
		if idx == -1 {
			return "", ""
		}
		idx += 2
		idx2 := strings.Index(p[idx+1:], `\`)
		if idx2 == -1 {
			return "", ""
		}
		idx2 += idx + 1
		return p[:idx2], p[idx2:]
	}
	return "", p
}

type pathComponents struct {
	drive string
	root  string
	parts []string
}

func normalize(p, altSep, sep, curDir string) pathComponents {
	p = strings.ReplaceAll(p, altSep, sep)
	p = strings.TrimRight(p, sep)
	lower := strings.ToLower(p)
	drive, path := splitDrive(lower)
	root := ""
	if strings.HasPrefix(path, sep) {
		root = sep
		path = strings.TrimPrefix(path, sep)
	}
	parts := strings.Split(path, sep)
	filtered := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" && part != curDir {
			filtered = append(filtered, part)
		}
	}
	return pathComponents{drive, root, filtered}
}
