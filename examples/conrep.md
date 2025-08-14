# Diagnosticism.Go Example - **conrep**

## Summary

Example illustrating use of Contingent Report API.

## Source

``` Go
package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
)

func main() {

	d.ConRep("some important information")
	d.ConRepf("some more important information: %s, %d, %f", `abc`, -1, 234.567)

	d.Abort("and we're out!")
}
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
some important information
some more important information: abc, -1, 234.567000
conrep: and we're out!
```

with an exit code of 1 (as a result of the call to ``Abort()``)


<!-- ########################### end of file ########################### -->

