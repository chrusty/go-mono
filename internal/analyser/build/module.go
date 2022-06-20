package build

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"golang.org/x/mod/modfile"
)

func (a *Analyser) GetRootModule() (string, error) {
	var modUndefined = "undefined"

	// Open go.mod:
	goModFilename := fmt.Sprintf("%s/go.mod", a.rootDir)
	goModFileContents, err := ioutil.ReadFile(goModFilename)
	if err != nil {
		return modUndefined, fmt.Errorf("Unable to open go.mod: %w", err)
	}
	logrus.Tracef("Read %d bytes from %s", len(goModFileContents), goModFilename)

	// Use it to inspect the root module:
	goMod, err := modfile.Parse(goModFilename, goModFileContents, nil)
	if err != nil {
		return modUndefined, fmt.Errorf("Unable to analyse repo root package in %s: %w", a.rootDir, err)
	}

	logrus.WithField("root_module", goMod.Module.Mod.String()).Debug("Read root module name")
	return goMod.Module.Mod.String(), nil
}
