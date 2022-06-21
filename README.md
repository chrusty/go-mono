go-mono
=======

Go-Mono is a toolkit for Golang monorepos. The aim is to provide a simple CLI tool which can tell you if a given Go package needs to be rebuilt, based on a dependency tree.


Why build 100 services when only 3 of them have actually changed?
-----------------------------------------------------------------

The aim for v1 is to be able to base these decisions entirely on packages within a single repo. This is useful in scenarios where you have a monorepo consisting of a bunch of shared libraries, and applications which are built on top of them. The decision to rebuilt is based on a dependency analysis using Go's `build` package to analyse imports, and follow them recursively. Go's `go.mod` doesn't help us here, as these imports are all "local" to the repo.

This functionality will be extended in v2, which will tell you if a package needs to be rebuilt based on _relevant_ dependencies within `go.mod`. For example, just because go.mod has changed doesn't mean that you have to rebuild everything in your repo - only things that actually _use_ the changed dependencies. This will work for _external_ dependencies.


Usage
-----

A couple of examples.


### Example 1

- Root of repo is the current working directory
- Package to check is the root of the repo (".")
- Diff against the default branch ("main")
- This command returns a non-zero code (1), letting you know that this package needs to be rebuilt
- It lists the single dependency that has changed

```
# go-mono
WARN Import has changed                            package=github.com/chrusty/go-mono/internal/git/shell
```


### Example 2

- Repo in some other directory
- Package to check is in a subdirectory
- Diff against a branch called "master"
- This command returns a non-zero code (1), letting you know that this package needs to be rebuilt
- It lists the two dependencies which have changed

```
# go-mono -diff=master -repo=./go/src/github.com/chrusty/protoc-gen-jsonschema -package=cmd/protoc-gen-jsonschema
WARN Import has changed                            package=github.com/chrusty/protoc-gen-jsonschema/internal/converter
WARN Import has changed                            package=github.com/chrusty/protoc-gen-jsonschema/internal/protos
```


### Example 3

- Nothing has changed

```
# go-mono
INFO No changes detected

```


Flags
-----

```
go-mono --help
Usage of go-mono:
  -debug
    	Run in debug mode?
  -diff string
    	Name of the branch to compare to (default "main")
  -package string
    	Path to the package to analyse (relative to the repo) (default ".")
  -repo string
    	Path to the root of the GIT repo (default ".")
  -trace
    	Run in trace mode?
```


How it works
------------

- Gets a list of files that have changed from a named branch
- For the Go package in the specified path
    - Analyse the dependency trees
    - Figure out if any of the local dependencies (imported from the same monorepo) are affected by the files in the list of changes
    - Returns a non-zero code if any of the dependencies have changed


Installation
------------

`go install github.com/chrusty/go-mono@latest`
