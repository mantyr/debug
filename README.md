# Debug

[![Build Status](https://travis-ci.org/mantyr/debug.svg?branch=master)](https://travis-ci.org/mantyr/debug)
[![GoDoc](https://godoc.org/github.com/mantyr/debug?status.png)](http://godoc.org/github.com/mantyr/debug)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This stable version

## Installation

    $ go get github.com/mantyr/debug

## Example

    $ go run ./main.go

    $ go run ./main.go --debug=true

```GO
package main

import (
    "github.com/mantyr/debug"
    "flag"
)

func init() {
    is_debug := flag.Bool("debug", false, "Debug programs")
    flag.Parse()

    debug.SetDebug(*is_debug)
}

func main() {
    // ...

    val := "123"
    debug.Println("log message", val)          // for log.Println("log message", val)
    debug.Printf("log message %q\r\n", val)    // for log.Printf("log message %q\r\n", val)

    // ...
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
