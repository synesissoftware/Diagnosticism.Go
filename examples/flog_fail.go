
package main

import (

	d "github.com/synesissoftware/diagnosticism.Go"
	severity "github.com/synesissoftware/diagnosticism.Go/severity"
)


func main() {

	d.EnableLogging(true)

	d.LogF(severity.Informational, "i=%d", 10)

	d.LogF(severity.Informational, "i=%d", "10")
}

