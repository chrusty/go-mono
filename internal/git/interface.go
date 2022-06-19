package git

type Gitter interface {
	Diff(commit string) ([]string, error) // Get a list of files which differ from the given commit/tag/branch
}
