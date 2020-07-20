# Diagnosticism.Go Example - **conrep_with_log_mirroring**

## Summary

Example illustrating use of Contingent Report API with log mirroring.

## Source

``` Go
package main

import (

	d "github.com/synesissoftware/Diagnosticism.Go"
)

func main() {

	d.EnableLogging(true)

	d.MirrorToLog(true)

	d.ConRep("some important information")
	d.ConRepf("some more important information: %s, %d, %f", `abc`, -1, 234.567)

	d.Abort("and we're out!")
}
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
2019/07/08 08:28:44 Failure : some important information
some important information
2019/07/08 08:28:44 Failure : some more important information: abc, -1, 234.567000
some more important information: abc, -1, 234.567000
2019/07/08 08:28:44 Alert : conrep_with_log_mirroring: and we're out!

conrep_with_log_mirroring: and we're out!
```

with an exit code of 1 (as a result of the call to ``Abort()``)
