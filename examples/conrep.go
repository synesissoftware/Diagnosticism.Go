package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
)

func main() {

	d.ConRep("some important information")
	d.ConRepf("some more important information: %s, %d, %f", `abc`, -1, 234.567)

	d.Abort("and we're out!")
}
