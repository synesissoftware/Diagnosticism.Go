// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 22nd February 2025
 * Updated: 11th October 2025
 */

package internal

import (
	"fmt"
	"runtime"
)

func getFileLineFunction(depth int, wantFunction bool) (string, int, string, bool) {

	// TODO: refactor in terms of `CallersFrames()`

	pc, file, line, ok := runtime.Caller(depth + 1)

	if ok {

		if wantFunction {

			function := runtime.FuncForPC(pc).Name()

			return file, line, function, true
		} else {

			return file, line, "", true
		}
	} else {
		return "", -1, "", false
	}
}

// Obtains the file information for the calling function.
func File(depth int) string {

	file, _, _, _ := getFileLineFunction(depth+1, false)

	return file
}

// Obtains the file and line information for the calling function.
func FileLine(depth int) string {

	file, line, _, ok := getFileLineFunction(depth+1, false)

	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	} else {
		return ""
	}
}

// Obtains the file, line, and function information for the calling
// function.
func FileLineFunction(depth int) string {

	file, line, function, ok := getFileLineFunction(depth+1, true)

	if ok {
		return fmt.Sprintf("%s:%d:%s", file, line, function)
	} else {
		return ""
	}
}

// Obtains the line function information for the calling function.
func Line(depth int) int {

	_, line, _, ok := getFileLineFunction(depth+1, true)

	if ok {
		return line
	} else {
		return -1
	}
}

// Obtains the line and function information for the calling function.
func LineFunction(depth int) string {

	_, line, function, ok := getFileLineFunction(depth+1, true)

	if ok {
		return fmt.Sprintf("%d:%s", line, function)
	} else {
		return ""
	}
}
