# diagnosticism.Go Example - **flog_fail**

## Summary

Example illustrating use of Log API with ``Logf()`` errors (in arguments).

## Source

``` Go
package main

import (

	d "github.com/synesissoftware/diagnosticism.Go"
	severity "github.com/synesissoftware/diagnosticism.Go/severity"
)


func main() {

	d.EnableLogging(true)

	d.Logf(severity.Informational, "i=%d", 10)

	d.Logf(severity.Informational, "i=%d", "10")
}
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
2019/07/08 08:29:44 Informational : i=10
2019/07/08 08:29:44 Informational : i=%!d(string=10)
```
