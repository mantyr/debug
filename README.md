# Debugs

[![Build Status](https://travis-ci.org/mantyr/debugs.svg?branch=master)](https://travis-ci.org/mantyr/debugs)
[![GoDoc](https://godoc.org/github.com/mantyr/debugs?status.png)](http://godoc.org/github.com/mantyr/debugs)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This stable version

## Installation

    $ go get github.com/mantyr/debugs

## Example

    $ go run ./main.go

    $ go run ./main.go --debug=true

```GO
package main

import (
    "github.com/mantyr/debugs"
    "flag"
)

func init() {
    is_debug := flag.Bool("debug", false, "Debug programs")
    flag.Parse()

    debugs.SetDebug(*is_debug)
}

func main() {
    // ...

    val := "123"
    debugs.Println("log message", val)          // for log.Println("log message", val)
    debugs.Printf("log message %q\r\n", val)    // for log.Printf("log message %q\r\n", val)

    // ...
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
