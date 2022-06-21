package build

import (
	"testing"

	"github.com/chrusty/go-mono/internal/analyser"
	"github.com/stretchr/testify/assert"
)

func TestAnalyser(t *testing.T) {

	// Bring up a new analyser (using this very repo):
	a, err := New("../../..")
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Implements(t, new(analyser.Analyser), a)

	// See if we can prefix a path now:
	assert.Equal(t, "../../../internal/tests", a.prefixPath("internal/tests"))

	// See if we can determine the root module:
	rootModule, err := a.GetRootModule()
	assert.NoError(t, err)
	assert.Equal(t, "github.com/chrusty/go-mono", rootModule)

	// See if we can list packages:
	importedPackages, err := a.GetPackages("internal/analyser/build")
	assert.NoError(t, err)
	assert.Len(t, importedPackages, 1)

	// See if we can detect when there are changes:
	changesDetected := a.AnalyseChanges([]string{"internal/analyser"}, importedPackages)
	assert.False(t, changesDetected)

	// See if we can detect when there are no changes:
	changesDetected = a.AnalyseChanges([]string{}, importedPackages)
	assert.False(t, changesDetected)
}
