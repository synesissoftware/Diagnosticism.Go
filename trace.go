/* /////////////////////////////////////////////////////////////////////////
 * File:    trace.go
 *
 * Purpose: Trace API for Diagnosticism.Go
 *
 * Created: 5th March 2019
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
 * - Neither the names of Matthew Wilson and Synesis Information Systems nor
 *   the names of any contributors may be used to endorse or promote
 *   products derived from this software without specific prior written
 *   permission.
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
	severity "github.com/synesissoftware/Diagnosticism.Go/severity"

	"bytes"
	"fmt"
)

var enableTracing bool

func EnableTracing(enable bool) {

	enableTracing = enable
}
func IsTracingEnabled() bool {

	return enableTracing
}

type TraceArgument struct {
	Name     string
	Value    interface{}
	nameOnly bool
}

func (arg TraceArgument) String() string {

	if arg.nameOnly {

		return fmt.Sprintf("%s(%T)", arg.Name, arg.Value)
	} else {

		return fmt.Sprintf("%s(%T)=%v", arg.Name, arg.Value, arg.Value)
	}
}

// [INTERNAL]
//
// Purpose: create a TraceArgument instance
//
// Parameters:
//   - +name+ (string) The name of the argument
//   - +nameOnly+ (bool) Whether only the name (and type) is to be shown
//   - +value+ (interface{}) The value of the argument
//
// Return: (TraceArgument) an instance
//
// Remarks: use of a boolean parameter is valid here because this is an internal method
func makeTraceArgument(name string, nameOnly bool, value interface{}) TraceArgument {

	return TraceArgument{Name: name, Value: value, nameOnly: nameOnly}
}

func Trarg(name string, value interface{}) TraceArgument {

	return makeTraceArgument(name, false, value)
}

func TrargNameOnly(name string, value interface{}) TraceArgument {

	return makeTraceArgument(name, true, value)
}

// Provides named-argument tracing of a function/method, as in:
//
//	 import d "github.com/synesissoftware/Diagnosticism.Go"
//
//		func SomeFunction(x, y int, order string) {
//
//			d.Trace("SomeFunction",
//				d.Trarg("x", x),
//				d.Trarg("y", y),
//				d.TrargNameOnly("order", order),
//			)
//
//			. . . impl. of SomeFunc()
//		}
//
// The first parameter function_name is a string, and the remaining
// parameters are a variable length list of TraceArgument instances, which
// may be created using the Trarg() and TrargNameOnly() functions
func Trace(function_name string, args ...TraceArgument) {

	if !enableTracing {

		return
	}

	var buffer bytes.Buffer

	buffer.WriteString(function_name)
	buffer.WriteString("(")

	for i, arg := range args {

		if i != 0 {

			buffer.WriteString(", ")
		}

		s := arg.String()

		buffer.WriteString(s)
	}

	buffer.WriteString(")")

	log_s(severity.Trace, buffer.String())
}

/* ///////////////////////////// end of file //////////////////////////// */
