// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 30th May 2019
 * Updated: 26th August 2025
 */

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

// Flags for controlling back-end behaviour/features.
type BackEndFlag int

const (
	NoPrefix          BackEndFlag = 1
	NoPrefixSeparator BackEndFlag = 2
	NoTime            BackEndFlag = 4
)

// Type describing an entry to be processed by the logging back-end.
type BackEndEntry struct {
	// The severity of the log statement.
	Severity severity.Severity
	// The time at which the log statement was consumed.
	Time time.Time
	// The statement message.
	Message string
}

// The BackEndHandlerFunc is called when a log statement is to be emitted.
type BackEndHandlerFunc func(be *BackEnd, bee *BackEndEntry)

// Backend log handler.
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

// Obtains the current backend handler function.
func GetBackEndHandlerFunc() *BackEnd {

	return activeBE
}

var enableLogging bool

// Sets whether logging is enabled.
func EnableLogging(enable bool) {

	enableLogging = enable
}

// Indicates whether logging is enabled.
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

// Logs the arguments at the given severity.
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

// Logs the formatted arguments at the given severity.
func Logf(severity severity.Severity, format string, args ...any) {

	if !enableLogging {

		return
	}

	s := fmt.Sprintf(format, args...)

	log_s(severity, s)
}

/* ///////////////////////////// end of file //////////////////////////// */
