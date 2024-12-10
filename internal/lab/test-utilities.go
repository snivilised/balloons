package lab

import (
	"os/exec"
	"path/filepath"
	"strings"
)

func Path(parent, relative string) string {
	segments := strings.Split(relative, "/")
	return filepath.Join(append([]string{parent}, segments...)...)
}

// Repo gets the path of the repo with relative joined on
func Repo(relative string) string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, _ := cmd.Output()

	repo := strings.TrimSpace(string(output))

	return combine(repo, relative)
}

// combine creates a path from the parent combined with the relative path. The relative
// path is a file system path so should only contain forward slashes, not the standard
// file path separator as denoted by filepath.Separator, typically used when interacting
// with the local file system. Do not use trailing "/".
func combine(parent, relative string) string {
	if relative == "" {
		return parent
	}

	return parent + "/" + relative
}
