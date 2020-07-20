# Diagnosticism.Go Example - **tracing**

## Summary

Example illustrating use of Trace API.

## Source

```Go
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
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
2019/07/08 08:25:05 Trace : SomeFunction(x(int), y(int), order(string)=first)
```
