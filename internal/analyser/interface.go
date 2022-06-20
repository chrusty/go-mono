package analyser

// Analyse go modules / imports, and analyse the changes to see if anything needs to be rebuilt:
type Analyser interface {
	AnalyseChanges(changedFiles []string, importedPackages ImportedPackages) (changesDetected bool)
	GetPackages(buildPackage string) (ImportedPackages, error)
	GetRootModule() (string, error)
}

// ImportedPackages is used to convey a map of imported package names, and their relative import paths:
type ImportedPackages map[string]string
