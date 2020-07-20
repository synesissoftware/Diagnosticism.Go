/* /////////////////////////////////////////////////////////////////////////
 * File:        log.go
 *
 * Purpose:     Log API for Diagnosticism.Go
 *
 * Created:     30th May 2019
 * Updated:     20th July 2020
 *
 * Home:        https://github.com/synesissoftware/Diagnosticism.Go
 *
 * Copyright (c) 2019-2020, Matthew Wilson
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
	"log"
)

var enableLogging bool

func EnableLogging(enable bool) {

	enableLogging = enable
}
func IsLoggingEnabled() bool {

	return enableLogging
}

func log_s(severity severity.Severity, s string) {

	log.Println(severity.String() + " : " + s)
}

func Log(severity severity.Severity, args ...interface{}) {

	if !enableLogging {

		return
	}

	var buffer bytes.Buffer

	for _, arg := range(args) {

		s := fmt.Sprintf("%v", arg)

		buffer.WriteString(s)
	}

	s := buffer.String()

	log_s(severity, s)
}

func Logf(severity severity.Severity, format string, args ...interface{}) {

	if !enableLogging {

		return
	}

	s := fmt.Sprintf(format, args...)

	log_s(severity, s)
}

/* ///////////////////////////// end of file //////////////////////////// */


