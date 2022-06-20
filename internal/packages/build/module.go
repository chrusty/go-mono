package build

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"golang.org/x/mod/modfile"
)

func GetModule(repoRoot string) (string, error) {
	var modUndefined = "undefined"

	// Open go.mod:
	goModFilename := fmt.Sprintf("%s/go.mod", repoRoot)
	goModFileContents, err := ioutil.ReadFile(goModFilename)
	if err != nil {
		return modUndefined, fmt.Errorf("Unable to open go.mod: %w", err)
	}
	logrus.Tracef("Read %d bytes from %s", len(goModFileContents), goModFilename)

	// Use it to inspect the root module:
	goMod, err := modfile.Parse(goModFilename, goModFileContents, nil)
	if err != nil {
		return modUndefined, fmt.Errorf("Unable to analyse repo root package in %s: %w", repoRoot, err)
	}

	return goMod.Module.Mod.String(), nil
}
