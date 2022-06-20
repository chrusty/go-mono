package analyser

type Analyser interface {
	GetPackages(buildPackage string) (map[string]string, error)
	GetRootModule() (string, error)
}
