/* /////////////////////////////////////////////////////////////////////////
 * File:    severity.go
 *
 * Purpose: Defines severity type(s) for Diagnosticism.Go
 *
 * Created: 30th May 2019
 * Updated: 22nd February 2025
 *
 * Home:    https://github.com/synesissoftware/Diagnosticism.Go
 *
 * Copyright (c) 2019-2025, Matthew Wilson and Synesis Information Systems
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

package severity

import (
	"fmt"
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type Severity int

func (severity Severity) String() string {

	return severityTranslator.SeverityToString(severity)
}

type SeverityTranslator interface {
	SeverityToString(severity Severity) string
}

/* /////////////////////////////////////////////////////////////////////////
 * constants
 */

const (
	Unspecified   Severity = 0
	Violation     Severity = 1
	Alert         Severity = 2
	Critical      Severity = 3
	Failure       Severity = 4
	Warning       Severity = 5
	Notice        Severity = 6
	Informational Severity = 7
	Debug0        Severity = 8
	Debug1        Severity = 9
	Debug2        Severity = 10
	Debug3        Severity = 11
	Debug4        Severity = 12
	Debug5        Severity = 13
	Trace         Severity = 14

	Warn Severity = Warning
	Fail Severity = Failure
	Info Severity = Informational
)

/* /////////////////////////////////////////////////////////////////////////
 * private types
 */

type defaultSeverityTranslator struct{}

func (dt defaultSeverityTranslator) SeverityToString(severity Severity) string {

	return TranslateStockSeverity(severity)
}

var severityTranslator SeverityTranslator = new(defaultSeverityTranslator)

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// Obtains the stock string form of a severity.
func TranslateStockSeverity(severity Severity) string {

	switch severity {

	case Violation:

		return "Violation"
	case Alert:

		return "Alert"
	case Critical:

		return "Critical"
	case Failure:

		return "Failure"
	case Warning:

		return "Warning"
	case Notice:

		return "Notice"
	case Informational:

		return "Informational"
	case Debug0:

		return "Debug0"
	case Debug1:

		return "Debug1"
	case Debug2:

		return "Debug2"
	case Debug3:

		return "Debug3"
	case Debug4:

		return "Debug4"
	case Debug5:

		return "Debug5"
	case Trace:

		return "Trace"
	default:

		return fmt.Sprintf("<Severity: %d>", int(severity))
	}
}

/* ///////////////////////////// end of file //////////////////////////// */
