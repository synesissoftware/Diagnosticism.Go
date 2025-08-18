/* /////////////////////////////////////////////////////////////////////////
 * File:    log.go
 *
 * Purpose: Log API for Diagnosticism.Go
 *
 * Created: 30th May 2019
 * Updated: 18th August 2025
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
	"log"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Flags for controlling back-end behaviour/features
type BackEndFlag int

const (
	NoPrefix          BackEndFlag = 1
	NoPrefixSeparator BackEndFlag = 2
	NoTime            BackEndFlag = 4
)

type BackEndEntry struct {
	Severity severity.Severity
	Time     time.Time
	Message  string
}

// The BackEndHandlerFunc is called when a log statement is to be emitted
type BackEndHandlerFunc func(be *BackEnd, bee *BackEndEntry)

type BackEnd struct {

	// Flags that control the back-end behaviour/features
	Flags BackEndFlag
	// The back-end handler function. May not be nil
	HandlerFunc BackEndHandlerFunc
	// The string to be used as a separator. If the empty string, then the
	// default separator - " : " - is used. If no separator is desired, the
	// NoPrefixSeparator flag must be specified
	PrefixSeparator string
}

func defaultBackEndHandlerFunc(be *BackEnd, bee *BackEndEntry) {

	log.Println(bee.Severity.String() + " : " + bee.Message)
}

var defaultBackEnd = BackEnd{

	Flags:           0,
	HandlerFunc:     defaultBackEndHandlerFunc,
	PrefixSeparator: " : ",
}

var activeBE *BackEnd = &defaultBackEnd

func check_atomically() bool {

	var be BackEndHandlerFunc
	var ui uintptr

	if unsafe.Sizeof(be) == unsafe.Sizeof(ui) {

		return true
	}

	return false
}

var atomicallyBE = check_atomically()
var mxBE = &sync.Mutex{}

func SetBackEnd(be *BackEnd) (r *BackEnd) {

	if r == nil {

		r = &defaultBackEnd
	}

	if atomicallyBE {

		var pp_old = (*unsafe.Pointer)(unsafe.Pointer(&activeBE))
		var p_new = *(*unsafe.Pointer)(unsafe.Pointer(&be))

		r0 := atomic.SwapPointer(pp_old, p_new)

		r = (*BackEnd)(r0)

	} else {

		mxBE.Lock()
		defer mxBE.Unlock()

		r = activeBE

		activeBE = be
	}

	return
}
func GetBackEndHandlerFunc() *BackEnd {

	return activeBE
}

var enableLogging bool

func EnableLogging(enable bool) {

	enableLogging = enable
}
func IsLoggingEnabled() bool {

	return enableLogging
}

func log_s(severity severity.Severity, message string) {

	be := activeBE

	bee := BackEndEntry{

		Severity: severity,
		Message:  message,
	}

	if 0 == (NoTime & be.Flags) {

		bee.Time = time.Now()
	}

	be.HandlerFunc(be, &bee)
}

func Log(severity severity.Severity, args ...any) {

	if !enableLogging {

		return
	}

	var buffer bytes.Buffer

	for _, arg := range args {

		s := fmt.Sprintf("%v", arg)

		buffer.WriteString(s)
	}

	s := buffer.String()

	log_s(severity, s)
}

func Logf(severity severity.Severity, format string, args ...any) {

	if !enableLogging {

		return
	}

	s := fmt.Sprintf(format, args...)

	log_s(severity, s)
}

/* ///////////////////////////// end of file //////////////////////////// */
