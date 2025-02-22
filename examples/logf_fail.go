package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
	sev "github.com/synesissoftware/Diagnosticism.Go/severity"
)

func main() {

	d.EnableLogging(true)

	d.Logf(sev.Informational, "i=%d", 10)

	d.Logf(sev.Informational, "i=%d", "10")
}
