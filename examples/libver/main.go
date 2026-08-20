package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("Diagnosticism v%s\n", d.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
