/* /////////////////////////////////////////////////////////////////////////
 * File:        loghandler.go
 *
 * Purpose:     LogHandler for diagnosticism.Go
 *
 * Created:     1st June 2019
 * Updated:     8th July 2019
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

	"bytes"
	"fmt"
	"net/http"
)

type LogRequestFlags int

const (

	// Do not log before the request
	LogRequest_NotBefore		LogRequestFlags	=	1 << iota
	// Do not log after the request
	LogRequest_NotAfter			LogRequestFlags	=	1 << iota
	// Do not include a BEFORE/AFTER prefix
	LogRequest_NoWhenLabel		LogRequestFlags	=	1 << iota

	// Do not include the method
	LogRequest_NotMethod		LogRequestFlags	=	1 << iota
	// Do not include the URL
	LogRequest_NotURL			LogRequestFlags	=	1 << iota
	// Include the protocol
	LogRequest_Protocol			LogRequestFlags	=	1 << iota
)

const (

	defaultLogRequestFlags		LogRequestFlags	=	0
	elementLogRequestWhenMask	LogRequestFlags	=	LogRequest_NotBefore | LogRequest_NotAfter
	elementLogRequestTypeMask	LogRequestFlags	=	LogRequest_NotMethod | LogRequest_NotURL | LogRequest_Protocol
)

const (

	AfterPrefix		=	"AFTER "
	BeforePrefix	=	"BEFORE "
)

type logStringFunc func(LogRequestFlags, *http.Request) string

var logStringFunctions = map[LogRequestFlags]logStringFunc {

	0 : logString_M_U_p,
	LogRequest_NotMethod : logString_m_U_p,
	LogRequest_NotURL : logString_M_u_p,
	LogRequest_Protocol : logString_M_U_P,
	LogRequest_NotMethod | LogRequest_NotURL : logString_X_X_X,
	LogRequest_NotMethod | LogRequest_Protocol : logString_X_X_X,
	LogRequest_NotURL | LogRequest_Protocol : logString_X_X_X,
	LogRequest_NotMethod | LogRequest_NotURL | LogRequest_Protocol : logString_X_X_X,
}

func logString_M_U_p(flags LogRequestFlags, req *http.Request) string {

	return fmt.Sprintf("%s %s", req.Method, req.URL)
}
func logString_M_U_P(flags LogRequestFlags, req *http.Request) string {

	return fmt.Sprintf("%s %s %s", req.Proto, req.Method, req.URL)
}
func logString_m_U_p(flags LogRequestFlags, req *http.Request) string {

	return req.URL.String()
}
func logString_M_u_p(flags LogRequestFlags, req *http.Request) string {

	return req.Method
}
func logString_X_X_X(flags LogRequestFlags, req *http.Request) string {

	var buff bytes.Buffer

	if 0 != (LogRequest_Protocol & flags ) {

		buff.WriteString(req.Proto)
	}

	if 0 == (LogRequest_NotMethod & flags) {

		if 0 != buff.Len() {

			buff.WriteRune(' ')
		}

		buff.WriteString(req.Method)
	}

	if 0 == (LogRequest_NotURL & flags) {

		if 0 != buff.Len() {

			buff.WriteRune(' ')
		}

		buff.WriteString(req.URL.String())
	}

	return buff.String()
}


func parseSeverityFromArgs(options ...interface{}) (severity.Severity) {

	for _, option := range(options) {

		switch v := option.(type) {

		case severity.Severity:

			return v
		}
	}

	return severity.Informational
}

// Middleware adapter that causes a request to be logged, according to the
// given flags and options
//
// Parameters:
//  - +flags+ (LogRequestFlags) A combination of flags that moderate the behaviour
//  - +options+ Optional arguments (see below)
//
// Options:
//  - * (severity.Severity) The first option of this type is used for before and/or after logging; if none specified, before and/or after logging is done using severity.Informational
func LogRequest(flags LogRequestFlags, options ...interface{}) (func(http.Handler) (http.Handler)) {

	sev := parseSeverityFromArgs(options)

	logDisplayFlags := flags & elementLogRequestTypeMask

	var afterPrefix string
	var beforePrefix string

	switch elementLogRequestWhenMask & flags {

	case 0:

		afterPrefix = AfterPrefix
		beforePrefix = BeforePrefix
		break
	case LogRequest_NotAfter:

		if 0 == (LogRequest_NoWhenLabel & flags) {

			beforePrefix = BeforePrefix
		}
		break
	case LogRequest_NotBefore:

		if 0 == (LogRequest_NoWhenLabel & flags) {

			afterPrefix = AfterPrefix
		}
		break
	case LogRequest_NotBefore | LogRequest_NotAfter:

		break
	}

	return func(h http.Handler) (http.Handler) {

		return http.HandlerFunc(func (resp http.ResponseWriter, req *http.Request) {

			logMsg := logStringFunctions[logDisplayFlags](logDisplayFlags, req)

			if 0 == (LogRequest_NotBefore & flags) {

				Log(sev, beforePrefix, logMsg)
			}
			if 0 == (LogRequest_NotAfter & flags) {

				defer Log(sev, afterPrefix, logMsg)
			}

			h.ServeHTTP(resp, req)
		})
	}
}

/* ///////////////////////////// end of file //////////////////////////// */


