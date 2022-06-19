package main

import (
	"flag"

	"github.com/chrusty/go-mono/internal/git/shell"
	"github.com/sirupsen/logrus"
)

var (
	buildPackage  = flag.String("package", ".", "Path to the package to analyse")
	compareCommit = flag.String("diff", "main", "Name of the branch / tag / commit to compare to")
	repoRoot      = flag.String("repo", ".", "Path to the root of the GIT repo")
)

func init() {
	flag.Parse()
	logrus.WithField("package", *buildPackage).Debug("Flag")
	logrus.WithField("diff", *compareCommit).Debug("Flag")
	logrus.WithField("repo", *repoRoot).Debug("Flag")
}

func main() {

	// Get a Gitter:
	gitter, err := shell.New(*repoRoot)
	if err != nil {
		logrus.WithError(err).WithField("repo", *repoRoot).Fatalf("Unable to prepare a Gitter")
	}

	// Get a list of changed files from the Gitter:
	changedFiles, err := gitter.Diff(*compareCommit)
	if err != nil {
		logrus.WithError(err).WithField("repo", *repoRoot).WithField("diff", *compareCommit).Fatalf("Unable to list changed files")
	}

	// Report:
	for _, changedFile := range changedFiles {
		logrus.WithField("filename", changedFile).Info("Changed file")
	}
}
