package main

import (
	"flag"
	"os"

	"github.com/chrusty/go-mono/internal/analyser/build"
	gogit "github.com/chrusty/go-mono/internal/git/go-git"
	"github.com/sirupsen/logrus"
)

var (
	compareBranch = flag.String("branch", "main", "Name of the branch to compare to")
	debug         = flag.Bool("debug", false, "Run in debug mode?")
	buildPackage  = flag.String("package", ".", "Path to the package to analyse (relative to the repo)")
	repoRoot      = flag.String("repo", ".", "Path to the root of the GIT repo")
	trace         = flag.Bool("trace", false, "Run in trace mode?")
)

const (
	version = "0.2.0"
)

func init() {
	flag.Parse()

	// Enable debug logging:
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Enable trace logging:
	if *trace {
		logrus.SetLevel(logrus.TraceLevel)
	}

	// Disable timestamps:
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	// Report config:
	logrus.WithField("debug", *debug).Debug("Flag")
	logrus.WithField("branch", *compareBranch).Debug("Flag")
	logrus.WithField("package", *buildPackage).Debug("Flag")
	logrus.WithField("repo", *repoRoot).Debug("Flag")
	logrus.WithField("trace", *trace).Debug("Flag")
	logrus.WithField("version", version).Debug("go-mono")
}

func main() {

	// Get a Gitter:
	gitter, err := gogit.New(*repoRoot)
	if err != nil {
		logrus.WithError(err).WithField("repo", *repoRoot).Fatalf("Unable to prepare a Gitter")
	}

	// Get a list of changed files from the Gitter:
	changedFiles, err := gitter.Diff(*compareBranch)
	if err != nil {
		logrus.WithError(err).WithField("repo", *repoRoot).WithField("diff", *compareBranch).Fatalf("Unable to list changed files")
	}

	// Prepare a package analyser:
	analyser, err := build.New(*repoRoot)
	if err != nil {
		logrus.WithError(err).Fatalf("Unable to prepare a package analyser")
	}

	// Get a list of imported packages:
	importedPackages, err := analyser.GetPackages(*buildPackage)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to find imported packages")
	}

	// Analyse the changes:
	changesDetected := analyser.AnalyseChanges(changedFiles, importedPackages)

	// If we've found any changes then return a non-zero code:
	if changesDetected {
		os.Exit(1)
	}

	// If we made it this far then we're clean:
	logrus.Info("No changes detected")
}
