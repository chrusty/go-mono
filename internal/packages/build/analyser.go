package build

import (
	"fmt"
)

type Analyser struct {
	rootDir    string
	rootModule string
}

// New rerturns a configured Analyser:
func New(rootDir string) (*Analyser, error) {

	// Make a new analyser:
	analyser := &Analyser{
		rootDir: rootDir,
	}

	// Attempt to get the root module:
	rootModule, err := analyser.GetRootModule()
	if err != nil {
		return nil, err
	}

	// Return a configured analyser:
	analyser.rootModule = rootModule
	return analyser, nil
}

func (a *Analyser) prefixPath(path string) string {
	return fmt.Sprintf("%s/%s", a.rootDir, path)
}
