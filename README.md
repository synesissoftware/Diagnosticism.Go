# diagnosticism.Go
Basic diagnostic facilities, for Go

## Introduction

**diagnosticism** provides low-level diagnostics facilities to support library programming. The first **diagnosticism** library was a C library with a C++ wrapper. There have been several implementations in other languages. **diagnosticism.Go** is the
Go version.

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Components](#components)
4. [Examples](#examples)
5. [Project Information](#project-information)

## Introduction

T.B.C.

## Installation

```Go

import diagnosticism "github.com/synesissoftware/diagnosticism.Go"
```

## Components

**diagnosticism.Python** provides components in the following categories:

* Contingent Reporting
* Diagnostic Logging
* Tracing

**NOTE**: for the moment, the Diagnostic Logging facilities emit to the standard error stream, via the Contingent Reporting API. In the near future this will be changed to work with more sophisticated logging libraries, including the standard logging facilities and the (as yet to be release) **Pantheios.Go**.

## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).

## Project Information

### Where to get help

[GitHub Page](https://github.com/synesissoftware/diagnosticism.Go "GitHub Page")

### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/diagnosticism.Go.

### Dependencies

None

### Dependent projects

* [**libpath.Go**](https://github.com/synesissoftware/libpath.Go/)
* [**recls.Go**](https://github.com/synesissoftware/recls.Go/)

### Related projects

* [**diagnosticism**](https://github.com/synesissoftware/diagnosticism/)
* [**diagnosticism.Python**](https://github.com/synesissoftware/diagnosticism.Python/)
* [**diagnosticism.Ruby**](https://github.com/synesissoftware/diagnosticism.Ruby/)

### License

**diagnosticism.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.
