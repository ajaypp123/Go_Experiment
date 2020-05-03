/*
GOPATH
	$GOPATH contains three high-level directories —
		1. $GOPATH/src — Your all code resides here
		2. $GOPATH/pkg — Packages used by Go internally to run itself
		3. $GOPATH/bin — Contains executable binaries like godep, golint

	With the latest release of Go 1.11, Go Modules are introduced which
	let you create your project outside the $GOPATH and improves package management a lot.

Rules
	1. Create single folder for project where $GOPATH is defined
	2. We can not use two same package of different version
	3. All external packages were kept in a vendor folder and pushed
	to the server. Dependency Hell can occur.

Module
	A module is a collection of related Go packages that are versioned together as a single unit.
	Modules record precise dependency requirements and create reproducible builds.

Package
	- go get url // Download third party package

*/

package main

func main() {
}
