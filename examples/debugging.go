package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
)

func f() {

	d.ConRepf("%s: here", d.FileLine())
	d.ConRepf("%s: here", d.FileLineFunction())
}

func main() {

	d.ConRepf("%s: here", d.FileLine())
	d.ConRepf("%s: here", d.FileLineFunction())

	f()
}
