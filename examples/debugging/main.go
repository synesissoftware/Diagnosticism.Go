package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
)

func f() {

	d.ConRepf("%s: here", d.FileLine())
	d.ConRepf("%s(): here", d.FileLineFunction())
	d.ConRepf("%s(): here", d.Function())
}

func main() {

	d.ConRepf("%s: here", d.FileLine())
	d.ConRepf("%s(): here", d.FileLineFunction())
	d.ConRepf("%s(): here", d.Function())

	f()
}
