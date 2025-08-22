# Diagnosticism.Go <!-- omit in toc -->

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/Diagnosticism.Go.svg)](https://github.com/synesissoftware/Diagnosticism.Go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/Diagnosticism.Go)](https://github.com/synesissoftware/Diagnosticism.Go/commits/master)
[![Go](https://github.com/synesissoftware/Diagnosticism.Go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/Diagnosticism.Go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/Diagnosticism.Go)](https://goreportcard.com/report/github.com/synesissoftware/Diagnosticism.Go)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/Diagnosticism.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/Diagnosticism.Go)

Basic diagnostic facilities, for Go


## Introduction

**Diagnosticism** provides low-level diagnostics facilities to support library programming. The first **Diagnosticism** library was a C library with a C++ wrapper. There have been several implementations in other languages. **Diagnosticism.Go** is the
Go version.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
	- [Constants](#constants)
	- [Functions](#functions)
		- [Contingent Reporting](#contingent-reporting)
		- [Debug](#debug)
		- [Logging/Tracing](#loggingtracing)
	- [Interfaces](#interfaces)
	- [Structures](#structures)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development Dependencies](#development-dependencies)
	- [Dependent projects](#dependent-projects)
	- [Related projects](#related-projects)
	- [License](#license)


## Installation

```Go

import diagnosticism "github.com/synesissoftware/Diagnosticism.Go"
```


## Components

**Diagnosticism.Go** provides components in the following categories:

* Contingent Reporting
* Diagnostic Logging
* Tracing

**NOTE**: for the moment, the Diagnostic Logging facilities emit to the standard error stream, via the Contingent Reporting API. In the near future this will be changed to work with more sophisticated logging libraries, including the standard logging facilities and the (as yet to be released) **Pantheios.Go**.


### Constants

No public constants are defined at this time.


### Functions

The following functions are defined:


#### Contingent Reporting

```Go
func Abort(message string)

func Abortf(format string, args ...any)

func ConRep(message string)

func ConRepf(format string, args ...any)

func MirrorToLog(enable bool)

func IsMirroringToLog() bool
```


#### Debug

```Go
// Obtains the file information for the calling function.
func File() string

// Obtains the file and line information for the calling function.
func FileLine() string

// Obtains the file, line, and function information for the calling
// function.
func FileLineFunction() string
```


#### Logging/Tracing

```Go
func SetBackEnd(be *BackEnd)

func GetBackEndHandlerFunc() *BackEnd

func EnableLogging(enable bool)

func IsLoggingEnabled() bool1

func Log(severity severity.Severity, args ...any)

func Logf(severity severity.Severity, format string, args ...any)
```

```Go
func EnableTracing(enable bool)

func IsTracingEnabled() bool

// Creates an argument descriptor that will trace the argument name, type,
// and value.
func Trarg(name string, value any) TraceArgument

// Creates an argument descriptor that will trace the argument name, but not
// type and value.
func TrargNameOnly(name string, value any) TraceArgument

// Creates an argument descriptor that will trace the argument name and
// type, but not value.
func TrargNameTypeOnly(name string, value any) TraceArgument

func TrargTrunc(name string, value any) TraceArgument

// Provides named-argument tracing of a function/method, as in:
//
//	 import d "github.com/synesissoftware/Diagnosticism.Go"
//
//		func SomeFunction(x, y int, order string) {
//
//			d.Trace(d.FileLineFunction(),
//				d.Trarg("x", x),
//				d.Trarg("y", y),
//				d.TrargNameTypeOnly("order", order),
//			)
//
//			. . . impl. of SomeFunc()
//		}
//
// The first parameter `function_name` is a string, and the remaining
// parameters are a variable length list of TraceArgument instances, which
// may be created using the `Trarg()` and `TrargNameOnly()` functions
func Trace(function_name string, args ...TraceArgument)
```

```Go
// Middleware adapter that causes a request to be logged, according to the
// given flags and options
//
// Parameters:
//   - +flags+ (LogRequestFlags) A combination of flags that moderate the behaviour
//   - +options+ Optional arguments (see below)
//
// Options:
//   - * (severity.Severity) The first option of this type is used for before and/or after logging; if none specified, before and/or after logging is done using severity.Informational
func LogRequest(flags LogRequestFlags, options ...any) func(http.Handler) http.Handler
```

```Go
// Obtains the stock string form of a severity.
func TranslateStockSeverity(severity Severity) string
```


### Interfaces

No public interface are defined at this time.


### Structures

No public structures are defined at this time.


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/Diagnosticism.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/Diagnosticism.Go.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/)


#### Development Dependencies

* [**testify**](https://github.com/stretchr/testify/)
* [**STEGoL**](https://github.com/synesissoftware/STEGoL/)


### Dependent projects

* [**libpath.Go**](https://github.com/synesissoftware/libpath.Go/)
* [**recls.Go**](https://github.com/synesissoftware/recls.Go/)


### Related projects

* [**Diagnosticism**](https://github.com/synesissoftware/Diagnosticism/)
* [**Diagnosticism.Python**](https://github.com/synesissoftware/Diagnosticism.Python/)
* [**Diagnosticism.Ruby**](https://github.com/synesissoftware/Diagnosticism.Ruby/)


### License

**Diagnosticism.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->

