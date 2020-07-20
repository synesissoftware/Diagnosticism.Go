
package main

import (

	d "github.com/synesissoftware/Diagnosticism.Go"
	severity "github.com/synesissoftware/Diagnosticism.Go/severity"
)


func main() {

	d.EnableLogging(true)

	d.Logf(severity.Informational, "i=%d", 10)

	d.Logf(severity.Informational, "i=%d", "10")
}

