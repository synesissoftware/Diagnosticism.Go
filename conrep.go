/* /////////////////////////////////////////////////////////////////////////
 * File:        conrep.go
 *
 * Purpose:     Contingent report API for diagnosticism.Go
 *
 * Created:     31st May 2019
 * Updated:     1st June 2019
 *
 * Home:        https://github.com/synesissoftware/diagnosticism.Go
 *
 * Copyright (c) 2019, Matthew Wilson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson and Synesis Software nor the names
 *   of any contributors may be used to endorse or promote products derived
 *   from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */


package diagnosticism

import (

	severity "github.com/synesissoftware/diagnosticism.Go/severity"

	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/* /////////////////////////////////////////////////////////////////////////
 * globals
 */

var mirroringToLog bool

var contingentReportWriter	io.Writer	=	os.Stderr

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

// TODO: separate this out using proper build constraints
func osIsWindows() bool {

	if ';' != os.PathListSeparator {

		return false
	}

	if '\\' != os.PathSeparator {

		return false
	}

	return true
}

func getProgramName() string {

	arg0 := os.Args[0]

	base := filepath.Base(arg0)

	if osIsWindows() {

		ext := filepath.Ext(base)

		switch strings.ToLower(ext) {

		case ".exe":

			base = base[0:len(base) - 4]
		}
	}

	return base
}

func conRepAddEol(severity severity.Severity, w io.Writer, msg string) {

	defer fmt.Fprintf(w, "%s\n", msg)

	if mirroringToLog {

		Log(severity, msg)
	}
}

func conRepWithEol(severity severity.Severity, w io.Writer, msg string) {

	defer io.WriteString(w, msg)

	if mirroringToLog {

		Log(severity, msg)
	}
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

func MirrorToLog(enable bool) {

	mirroringToLog = enable
}
func IsMirroringToLog() bool {

	return mirroringToLog
}

func ConRep(message string) {

	conRepAddEol(severity.Failure, contingentReportWriter, message)
}

func ConRepF(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	conRepAddEol(severity.Failure, contingentReportWriter, msg)
}

func Abort(message string) {

	msg := fmt.Sprintf("%s: %s\n", getProgramName(), message)

	conRepWithEol(severity.Alert, contingentReportWriter, msg)

	os.Exit(1)
}

func AbortF(format string, args ...interface{}) {

	message := fmt.Sprintf(format, args...)

	msg := fmt.Sprintf("%s: %s\n", getProgramName(), message)

	conRepWithEol(severity.Alert, contingentReportWriter, msg)

	os.Exit(1)
}

/* ///////////////////////////// end of file //////////////////////////// */


