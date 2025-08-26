// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 30th May 2019
 * Updated: 27th August 2025
 */

package severity

import (
	"fmt"
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

// API Severity level.
type Severity int

// Default string form of the stock severity levels.
func (severity Severity) String() string {

	return severityTranslator.SeverityToString(severity)
}

// SeverityTranslator is implemented by a type to customise the translation
// of [Severity] into string form.
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

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// Obtains the stock string form of a given [Severity].
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
