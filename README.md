# Diagnosticism.Go
Basic diagnostic facilities, for Go

## Introduction

**Diagnosticism** provides low-level diagnostics facilities to support library programming. The first **Diagnosticism** library was a C library with a C++ wrapper. There have been several implementations in other languages. **Diagnosticism.Go** is the
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

import diagnosticism "github.com/synesissoftware/Diagnosticism.Go"
```

## Components

**Diagnosticism.Go** provides components in the following categories:

* Contingent Reporting
* Diagnostic Logging
* Tracing

**NOTE**: for the moment, the Diagnostic Logging facilities emit to the standard error stream, via the Contingent Reporting API. In the near future this will be changed to work with more sophisticated logging libraries, including the standard logging facilities and the (as yet to be release) **Pantheios.Go**.

## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).

## Project Information

### Where to get help

[GitHub Page](https://github.com/synesissoftware/Diagnosticism.Go "GitHub Page")

### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/Diagnosticism.Go.

### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go.Go/)

### Dependent projects

* [**libpath.Go**](https://github.com/synesissoftware/libpath.Go/)
* [**recls.Go**](https://github.com/synesissoftware/recls.Go/)

### Related projects

* [**Diagnosticism**](https://github.com/synesissoftware/Diagnosticism/)
* [**Diagnosticism.Python**](https://github.com/synesissoftware/Diagnosticism.Python/)
* [**Diagnosticism.Ruby**](https://github.com/synesissoftware/Diagnosticism.Ruby/)

### License

**Diagnosticism.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.
