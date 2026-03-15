package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
)

func SomeFunction(x, y int, order, alphabet string) {
	d.Trace("SomeFunction",
		d.TrargNameOnly("x", x),
		d.TrargNameTypeOnly("y", y),
		d.Trarg("order", order),
		d.TrargTrunc("alphabet", alphabet),
	)
	//. . . impl. of SomeFunc()
}

func main() {

	d.EnableTracing(true)

	SomeFunction(1, 2, "first", "abcdefghijklmnopqrstuvwxyz")
}
