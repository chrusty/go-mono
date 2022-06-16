go-mono
=======

Toolkit for Golang monorepos


How it works
------------

- Gets a list of files that have changed
    - From the main branch
    - From the previous commit
    - From a named branch / commit / tag
- For a given list of files / directories (these are the buildable binaries provided by CLI)
    - Analyse their dependency trees
    - Figure out if any of the dependencies are affected by the files in the list of changes
    - Returns a list of "main" packages that need to be built
