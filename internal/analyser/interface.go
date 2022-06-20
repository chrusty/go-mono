package analyser

type Analyser interface {
	GetPackages(buildPackage string) (ImportedPackages, error)
	GetRootModule() (string, error)
}

// ImportedPackages is used to convey a map of imported package names, and their relative import paths:
type ImportedPackages map[string]string
