Development Roadmap
===================

v0.1
----

- [x] Read CLI flags
    - [x] Binary package to base dependency tree on
    - [x] GIT branch / commit to diff against
- [x] Build a list of changed files
- [x] Figure out the root module name
- [x] Recursively analyse package dependencies and build a list
- [x] Return non-zero error code if any of the given paths need to be built
- [x] Read GIT repo from any (external) directory
- [x] Analyse packages from any (external) directory


v0.2
----

- [x] Remame the "packages" package to "analyser"
- [x] Bring comparison logic into the analyser
- [x] Stop shelling out to GIT, and implement a Gitter using a GIT library


Backlog
-------
- Analyse go.mod for packages which have changed and are used by this target
