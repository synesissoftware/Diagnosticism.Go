# Diagnosticism.Go Example - **conrep_with_custom_backend**

## Summary

Example illustrating use of Contingent Report API with custom logging back-end

## Source

``` Go
package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
	sev "github.com/synesissoftware/Diagnosticism.Go/severity"

	"fmt"
	"os"
)

var BackEndWithUpcaseSeverityToStderr = d.BackEnd {

	Flags		:	d.NoTime,
	HandlerFunc	:	func(be *d.BackEnd, bee* d.BackEndEntry) {

		fmt.Fprintf(os.Stderr, "custom (to stderr): [%s] %s\n", bee.Severity, bee.Message)
	},
}

func main() {

	d.EnableLogging(true)

	d.Log(sev.Info, "one message")

	d.SetBackEnd(&d.BackEnd {

		Flags		:	d.NoTime,
		HandlerFunc	:	func(be *d.BackEnd, bee* d.BackEndEntry) {

			fmt.Println("custom: " + bee.Severity.String() + " : " + bee.Message)
		},
	})

	d.Log(sev.Info, "a second message")

	d.SetBackEnd(&BackEndWithUpcaseSeverityToStderr)

	d.Log(sev.Info, "a third message")
}
```

## Execution

When executed, it gives output (to the standard error stream) along the lines of

```
2020/07/23 10:11:15 Informational : one message
custom: Informational : a second message
custom (to stderr): [Informational] a third message
```


<!-- ########################### end of file ########################### -->

