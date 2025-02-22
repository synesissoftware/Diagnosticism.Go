// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 22nd February 2025
 * Updated: 22nd February 2025
 */

package diagnosticism

import (
	"fmt"
	"runtime"
)

// TODO: refactor in terms of `CallersFrames()`

// Obtains the file and line information for the calling function.
func FileLine() string {

	_, file, line, ok := runtime.Caller(1)

	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	} else {
		return ""
	}
}

// Obtains the file, line, and function information for the calling
// function.
func FileLineFunction() string {

	pc, file, line, ok := runtime.Caller(1)
	function := runtime.FuncForPC(pc).Name()

	if ok {
		return fmt.Sprintf("%s:%d:%s", file, line, function)
	} else {
		return ""
	}
}
