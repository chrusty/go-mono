package build

import (
	"fmt"

	"github.com/sirupsen/logrus"
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
	rootModule, err := analyser.GetModule()
	if err != nil {
		return nil, err
	}

	// Return a configured analyser:
	analyser.rootModule = rootModule
	logrus.WithField("root_module", analyser.rootModule).Info("Read root module name")
	return analyser, nil
}

func (a *Analyser) prefixPath(path string) string {
	return fmt.Sprintf("%s/%s", a.rootDir, path)
}
