package build

import (
	"fmt"
	"go/build"
	"strings"

	"github.com/sirupsen/logrus"
)

// GetPackages find all local packages under a directory:
func GetPackages(rootModule, buildPackage string) (map[string]string, error) {

	// Keep packages in here:
	packages := make(map[string]string)

	// Get the package info:
	buildContext, err := build.ImportDir(buildPackage, build.ImportComment)
	if err != nil {
		return nil, fmt.Errorf("Unable to analyse go package in %s: %w", buildPackage, err)
	}

	// Add the packages:
	for _, importedPackage := range buildContext.Imports {

		logrus.Tracef("Found import: %s", importedPackage)

		// Only proceed if this package is within our monorepo and hasn't already been added:
		if _, exists := packages[importedPackage]; !exists && strings.Contains(importedPackage, rootModule) {

			// Figure out the relative path to the imported package:
			relativePackagePath := strings.Replace(importedPackage, rootModule, "", -1)[1:]
			packages[importedPackage] = relativePackagePath

			// Recurse:
			logrus.Tracef("Recursing %s", relativePackagePath)
			recursedPackages, err := GetPackages(rootModule, relativePackagePath)
			if err != nil {
				return nil, fmt.Errorf("Unable to recurse go package in %s: %w", relativePackagePath, err)
			}

			// Combine the new packages with our existing ones:
			for k, v := range recursedPackages {
				packages[k] = v
			}
		}
	}

	return packages, nil
}
