// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const (
	appBin      = "gostock.exe"
	packagePath = "github.com/tkido/gostock"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Run

// Clean clean up
func Clean() {
	fmt.Println("Clean...")
	os.Remove(appBin)
}

// Build build a bin file
func Build() error {
	mg.Deps(Clean)
	fmt.Println("Build...")
	return sh.RunV("go", "build", "-o", appBin, "-v")
}

// Run execute app
func Run() error {
	mg.Deps(Build)
	fmt.Println("Run...")
	return sh.RunV("./"+appBin, "-t", "rss")
}

// Release releases an exe file
func Release() error {
	mg.Deps(Test)
	return Build()
}

// Test tests all packages under this package
func Test() error {
	packages := []string{
		"edinet",
		"kaiji",
		"my",
		"page",
		"patrol",
		"rireki",
		"rss",
		"spider",
		"statistics",
		"tdnet",
		"ufo",
		"xbrl",
	}
	for _, p := range packages {
		err := sh.RunV("go", "test", "-v", packagePath+"/"+p)
		if err != nil {
			return err
		}
	}
	return nil
}
