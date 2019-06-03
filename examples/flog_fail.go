
package main

import (

	d "github.com/synesissoftware/diagnosticism.Go"
	severity "github.com/synesissoftware/diagnosticism.Go/severity"
)


func main() {

	d.EnableLogging(true)

	d.Logf(severity.Informational, "i=%d", 10)

	d.Logf(severity.Informational, "i=%d", "10")
}

