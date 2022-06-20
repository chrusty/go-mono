Development Roadmap
===================

v0.1
----

- [x] Read CLI flags
    - [x] Binary package to base dependency tree on
    - [x] GIT branch / commit to diff against
- [x] Build a list of changed files
- [ ] Figure out the root module name
- [x] Recursively analyse package dependencies and build a list
- [ ] Return non-zero error code if any of the given paths need to be built


Backlog
-------
- Analyse go.mod for packages which have changed and are used by this target
