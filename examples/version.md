# Diagnosticism.Go Example - **tracing**

## Summary

Example illustrating use of Trace API.

## Source

```Go
package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
)

func main() {

	d.ConRepf("Diagnosticism v%v", d.VersionString())
}
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
Diagnosticism v0.6.1
```


<!-- ########################### end of file ########################### -->

