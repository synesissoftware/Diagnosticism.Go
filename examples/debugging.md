# Diagnosticism.Go Example - **debugging**

## Summary

Example illustrating use of Debugging API functions `FileLine()`, `FileLineFunction()`, `Function()`.

## Source

```Go
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
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
~/synesissoftware/freelibs/Diagnosticism/Diagnosticism.Go/examples/debugging.go:16: here
~/synesissoftware/freelibs/Diagnosticism/Diagnosticism.Go/examples/debugging.go:17:main.main(): here
main.main(): here
~/synesissoftware/freelibs/Diagnosticism/Diagnosticism.Go/examples/debugging.go:9: here
~/synesissoftware/freelibs/Diagnosticism/Diagnosticism.Go/examples/debugging.go:10:main.f(): here
main.f(): here
```


<!-- ########################### end of file ########################### -->

