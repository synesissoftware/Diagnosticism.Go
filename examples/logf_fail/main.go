package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
	sev "github.com/synesissoftware/Diagnosticism.Go/severity"
)

func main() {

	d.EnableLogging(true)

	format := "i=%d"

	d.Logf(sev.Informational, format, 10)

	d.Logf(sev.Informational, format, "10")
}
