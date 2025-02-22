# Diagnosticism.Go Example - **debugging**

## Summary

Example illustrating use of Debugging API.

## Source

```Go
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
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
~/forks/synesissoftware/Diagnosticism/Diagnosticism.Go/examples/debugging.go:15: here
~/forks/synesissoftware/Diagnosticism/Diagnosticism.Go/examples/debugging.go:16:main.main: here
~/forks/synesissoftware/Diagnosticism/Diagnosticism.Go/examples/debugging.go:9: here
~/forks/synesissoftware/Diagnosticism/Diagnosticism.Go/examples/debugging.go:10:main.f: here
```
