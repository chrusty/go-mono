package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Gitter struct {
	repoRootDirectory string
}

// Return a new Gitter, configured to point to a repo at a given location:
func New(repoRootDirectory string) (*Gitter, error) {

	// Check that the directory exists:
	directoryInfo, err := os.Stat(repoRootDirectory)
	if os.IsNotExist(err) {
		return nil, err
	}

	// Make sure it is actually a directory:
	if !directoryInfo.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", repoRootDirectory)
	}

	// Return a new Gitter:
	return &Gitter{
		repoRootDirectory: repoRootDirectory,
	}, nil
}

// Get a list of files which differ from the given commit/tag/branch:
func (g *Gitter) Diff(commit string) ([]string, error) {
	var changedFiles = []string{}

	// Make a GIT command (git diff <commit> --name-only):
	cmd := exec.Command("git", "diff", commit, "--name-only")

	// Get the stdout:
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Split up the output and return the list of changed files:
	splitOutputLines := strings.Split(string(output), "\n")
	for _, splitOutputLine := range splitOutputLines {
		if len(splitOutputLine) > 0 {
			changedFiles = append(changedFiles, splitOutputLine)
		}
	}

	return changedFiles, nil
}
