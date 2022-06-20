package build

import (
	"path"

	"github.com/chrusty/go-mono/internal/analyser"
	"github.com/sirupsen/logrus"
)

func (a *Analyser) AnalyseChanges(changedFiles []string, importedPackages analyser.ImportedPackages) (changesDetected bool) {

	// Report on the imports we found:
	for importedPackage, relativeImport := range importedPackages {

		// Compare to changed files:
		for _, changedFile := range changedFiles {
			if path.Dir(changedFile) == relativeImport {
				logrus.WithField("package", importedPackage).Warn("Import has changed")
				changesDetected = true
			}
		}
	}

	return
}
