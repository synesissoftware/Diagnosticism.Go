
package diagnosticism_test

import (

	d "github.com/synesissoftware/Diagnosticism.Go"
	sev "github.com/synesissoftware/Diagnosticism.Go/severity"
	stegol "github.com/synesissoftware/STEGoL"

	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_custom_BackEnd_no_time(t *testing.T) {

	buf := new(bytes.Buffer)

	d.SetBackEnd(&d.BackEnd {

		Flags		:	d.NoTime,
		HandlerFunc	:	func(be *d.BackEnd, bee *d.BackEndEntry) {

			fmt.Fprintf(buf, "%s : %s\n", bee.Severity, bee.Message)
		},
	})
	d.EnableLogging(true)

	d.Log(sev.Notice, "message-1")

	stegol.CheckStringEqual(t, "Notice : message-1\n", buf.String())

	buf.Truncate(0)

	stegol.CheckStringEqual(t, "", buf.String())

	d.Log(sev.Notice, "message-1")
	d.Log(sev.Warning, "message-2")

	stegol.CheckStringEqual(t, "Notice : message-1\nWarning : message-2\n", buf.String())
}

func Test_custom_BackEnd_no_time_and_upcase_severities(t *testing.T) {

	buf := new(bytes.Buffer)

	d.SetBackEnd(&d.BackEnd {

		Flags		:	d.NoTime,
		HandlerFunc	:	func(be *d.BackEnd, bee *d.BackEndEntry) {

			fmt.Fprintf(buf, "%s : %s\n", strings.ToUpper(bee.Severity.String()), bee.Message)
		},
	})
	d.EnableLogging(true)

	d.Log(sev.Notice, "message-1")

	stegol.CheckStringEqual(t, "NOTICE : message-1\n", buf.String())

	buf.Truncate(0)

	stegol.CheckStringEqual(t, "", buf.String())

	d.Log(sev.Notice, "message-1")
	d.Log(sev.Warning, "message-2")

	stegol.CheckStringEqual(t, "NOTICE : message-1\nWARNING : message-2\n", buf.String())
}

