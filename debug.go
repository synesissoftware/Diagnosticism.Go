// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 22nd February 2025
 * Updated: 1st March 2026
 */

package diagnosticism

import (
	"errors"

	"github.com/synesissoftware/Diagnosticism.Go/internal"
)

var (
	errDepthCannotBeNegative = errors.New("depth cannot be negative")
)

// Obtains the file information for the calling context.
func File() string {

	return internal.File(1)
}

// Obtains the file and line information for the calling context.
func FileLine() string {

	return internal.FileLine(1)
}

// Obtains the file, line, and function information for the calling
// context.
func FileLineFunction() string {

	return internal.FileLineFunction(1)
}

// Obtains the function information for the calling context.
func Function() string {

	return internal.Function(1)
}

// Obtains the line information for the calling context.
func Line() int {

	return internal.Line(1)
}

// Obtains the line and function information for the calling context.
func LineFunction() string {

	return internal.LineFunction(1)
}

// Obtains the file information for the calling context at the given depth.
func GetFileFor(depth int) (string, error) {

	if depth < 0 {
		return "", errDepthCannotBeNegative
	}

	return internal.File(1 + depth), nil
}

// Obtains the file and line information for the calling context at the
// given depth.
func GetFileLineFor(depth int) (string, error) {

	if depth < 0 {
		return "", errDepthCannotBeNegative
	}

	return internal.FileLine(1 + depth), nil
}

// Obtains the file, line, and function information for the calling
// context at the given depth.
func GetFileLineFunctionFor(depth int) (string, error) {

	if depth < 0 {
		return "", errDepthCannotBeNegative
	}

	return internal.FileLineFunction(1 + depth), nil
}

// Obtains the function information for the calling context at the given
// depth.
func GetFunctionFor(depth int) (string, error) {

	if depth < 0 {
		return "", errDepthCannotBeNegative
	}

	return internal.Function(1 + depth), nil
}

// Obtains the line information for the calling context at the given depth.
func GetLineFor(depth int) (int, error) {

	if depth < 0 {
		return 0, errDepthCannotBeNegative
	}

	return internal.Line(1 + depth), nil
}
