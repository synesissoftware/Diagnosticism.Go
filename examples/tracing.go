
package main

import (

	d "github.com/synesissoftware/Diagnosticism.Go"
)


func SomeFunction(x, y int, order string) {
	d.Trace("SomeFunction",
		d.TrargNameOnly("x", x),
		d.TrargNameOnly("y", y),
		d.Trarg("order", order),
	)
	//. . . impl. of SomeFunc()
}

func main() {

	d.EnableTracing(true)

	SomeFunction(1, 2, "first")
}

