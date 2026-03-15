package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("Diagnosticism v%s\n", ver2go.CalcVersionString(d.VersionMajor, d.VersionMinor, d.VersionPatch, d.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}
