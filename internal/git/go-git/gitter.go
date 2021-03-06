package gogit

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/sirupsen/logrus"
)

type Gitter struct {
	headCommit        *object.Commit
	repoRootDirectory string
	repo              *git.Repository
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

	// Open our repo:
	repo, err := git.PlainOpen(repoRootDirectory)
	if err != nil {
		return nil, err
	}

	// Look up the reference for the repo head:
	headReference, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("Unable to retieve the target commit reference: %w", err)
	}

	// Get the head commit (so we can diff against it later):
	headCommit, err := repo.CommitObject(headReference.Hash())
	if err != nil {
		return nil, fmt.Errorf("Unable to retieve the head commit: %w", err)
	}
	logrus.WithField("hash", headCommit.Hash).Debug("Found the head reference")

	// Return a new Gitter:
	return &Gitter{
		repoRootDirectory: repoRootDirectory,
		repo:              repo,
		headCommit:        headCommit,
	}, nil
}

// Get a list of files which differ from the given commit/tag/branch:
func (g *Gitter) Diff(compareBranch string) ([]string, error) {
	var changedFiles = []string{}

	// Look up the reference for our comparison branch:
	compareBranchReference, err := g.repo.Reference(plumbing.NewBranchReferenceName(compareBranch), false)
	if err != nil {
		return nil, fmt.Errorf("Unable to retieve the branch commit reference: %w", err)
	}

	// Get the source commit (so we can diff against it later):
	compareBranchObject, err := g.repo.CommitObject(compareBranchReference.Hash())
	if err != nil {
		return nil, fmt.Errorf("Unable to retieve the branch commit: %w", err)
	}

	// Calculate the patch (diff) between the source and target commits:
	patch, err := g.headCommit.Patch(compareBranchObject)
	if err != nil {
		return nil, fmt.Errorf("Unable to calculate patch between branch and head: %w", err)
	}

	// Add the changes to our list:
	for _, filePatch := range patch.FilePatches() {
		fromFile, _ := filePatch.Files()
		logrus.WithField("filename", fromFile.Path()).Debug("Changed file detected")
		changedFiles = append(changedFiles, fromFile.Path())
	}

	return changedFiles, nil
}
