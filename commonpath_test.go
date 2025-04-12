package commonpath_test

import (
	"testing"

	"github.com/tetsuo/commonpath"
)

func TestCommonPathUnix(t *testing.T) {
	tests := []struct {
		name     string
		paths    []string
		expected string
		wantErr  bool
	}{
		{
			name:     "no paths",
			paths:    []string{},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "single path",
			paths:    []string{"/a/b/c"},
			expected: "/a/b/c",
		},
		{
			name:     "single root path",
			paths:    []string{"/"},
			expected: "/",
		},
		{
			name:     "root only",
			paths:    []string{"/a", "/b", "/c"},
			expected: "/",
		},
		{
			name:     "identical",
			paths:    []string{"/x/y/z", "/x/y/z", "/x/y/z"},
			expected: "/x/y/z",
		},
		{
			name:     "nested common",
			paths:    []string{"/a/b/c/d", "/a/b/c/e", "/a/b/c/f/g"},
			expected: "/a/b/c",
		},
		{
			name:     "common /a/b",
			paths:    []string{"/a/b/c", "/a/b/d", "/a/b/e/f"},
			expected: "/a/b",
		},
		{
			name:     "two root paths",
			paths:    []string{"/", "/"},
			expected: "/",
		},
		{
			name:     "mixed no overlap",
			paths:    []string{"/foo", "/bar"},
			expected: "/",
		},
		{
			name:     "relative vs absolute mix",
			paths:    []string{"a/b", "/a/b"},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "all relative paths with common",
			paths:    []string{"x/y/z", "x/y/a", "x/y/b/c"},
			expected: "x/y",
		},
		{
			name:     "all relative paths no common",
			paths:    []string{"foo", "bar", "baz"},
			expected: "",
		},
		{
			name:     "absolute and empty path mix",
			paths:    []string{"/a/b", ""},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "only relative paths",
			paths:    []string{"a", "a/b", "a/b/c"},
			expected: "a",
		},
		{
			name:     "identical relative paths",
			paths:    []string{"foo/bar", "foo/bar"},
			expected: "foo/bar",
		},
		{
			name:     "deep common with trailing slash variants",
			paths:    []string{"/a/b/c/", "/a/b/c/d", "/a/b/c/e/f"},
			expected: "/a/b/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := commonpath.CommonPathUnix(tt.paths)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if got != tt.expected {
				t.Errorf("want %q, got %q", tt.expected, got)
			}
		})
	}
}
func TestCommonPathWin(t *testing.T) {
	tests := []struct {
		name     string
		paths    []string
		expected string
		wantErr  bool
	}{
		{
			name:     "single path root",
			paths:    []string{`C:\`},
			expected: `c:\`,
		},
		{
			name:     "single segment drive only",
			paths:    []string{`C:`},
			expected: `c:\`,
		},
		{
			name:     "common path on same drive",
			paths:    []string{`C:\x\y\z`, `C:\x\y\m`, `C:\x\y\n`},
			expected: `c:\x\y`,
		},
		{
			name:     "case insensitive match",
			paths:    []string{`C:\A\B\C`, `c:\a\b\d`},
			expected: `c:\a\b`,
		},
		{
			name:     "trailing backslash normalized",
			paths:    []string{`C:\foo\`, `C:\foo\bar`},
			expected: `c:\foo`,
		},
		{
			name:    "different drives error",
			paths:   []string{`C:\x`, `D:\y`},
			wantErr: true,
		},
		{
			name:     "UNC path common root",
			paths:    []string{`\\server\share\folder1\sub`, `\\server\share\folder2`},
			expected: `\\server\share`,
		},
		{
			name:     "absolute paths same",
			paths:    []string{`C:\folder\sub`, `C:\folder\sub\child`},
			expected: `c:\folder\sub`,
		},
		{
			name:    "mixed absolute and relative",
			paths:   []string{`C:\folder`, `folder\child`},
			wantErr: true,
		},
		{
			name:    "UNC path mismatch",
			paths:   []string{`\\server1\share\path`, `\\server2\share\path`},
			wantErr: true,
		},
		{
			name:    "empty path input",
			paths:   []string{},
			wantErr: true,
		},
		{
			name:     "identical UNC paths",
			paths:    []string{`\\server\share\dir`, `\\server\share\dir`},
			expected: `\\server\share\dir`,
		},
		{
			name:    "drive absolute vs drive relative",
			paths:   []string{`C:\foo`, `C:bar`},
			wantErr: true,
		},
		{
			name:    "UNC rooted vs non-rooted",
			paths:   []string{`\\server\share\folder`, `\\server\sharefolder`},
			wantErr: true,
		},
		{
			name:    "rooted vs not-rooted paths",
			paths:   []string{`\foo\bar`, `foo\bar`},
			wantErr: true,
		},
		{
			name:    "UNC path missing share",
			paths:   []string{`\\host`},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := commonpath.CommonPathWin(tt.paths)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if got != tt.expected {
				t.Errorf("want %q, got %q", tt.expected, got)
			}
		})
	}
}
