
package main

import (

	d "github.com/synesissoftware/diagnosticism.Go"
)

func main() {

	d.EnableLogging(true)

	d.MirrorToLog(true)

	d.ConRep("some important information")
	d.ConRepF("some more important information: %s, %d, %f", `abc`, -1, 234.567)

	d.Abort("and we're out!")
}

