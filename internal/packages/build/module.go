package build

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func GetModule(repoRoot string) (string, error) {
	var modUndefined = "undefined"

	// Open go.mod:
	goModFilename := fmt.Sprintf("%s/go.mod", repoRoot)
	goModFile, err := os.Open(goModFilename)
	if err != nil {
		return modUndefined, fmt.Errorf("Unable to open go.mod: %w", err)
	}

	// Read it:
	var goModFileContents []byte
	if _, err := goModFile.Read(goModFileContents); err != nil {
		return modUndefined, fmt.Errorf("Unable to read go.mod: %w", err)
	}

	// Use it to inspect the root module:
	goMod, err := modfile.Parse(goModFilename, goModFileContents, nil)
	if err != nil {
		return modUndefined, fmt.Errorf("Unable to analyse repo root package in %s: %w", repoRoot, err)
	}

	return goMod.Module.Mod.String(), nil

}
