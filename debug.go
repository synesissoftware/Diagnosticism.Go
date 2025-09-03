// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 22nd February 2025
 * Updated: 3rd September 2025
 */

package diagnosticism

import (
	"github.com/synesissoftware/Diagnosticism.Go/internal"
)

// TODO: refactor in terms of `CallersFrames()`

// Obtains the file information for the calling function.
func File() string {

	return internal.File(1)
}

// Obtains the file and line information for the calling function.
func FileLine() string {

	return internal.FileLine(1)
}

// Obtains the file, line, and function information for the calling
// function.
func FileLineFunction() string {

	return internal.FileLineFunction(1)
}
