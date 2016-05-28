package debugs

import (
    "log"
)

var is_debug bool

// No safe. Do not change the values if you do not believe that non-concurrent access
func SetDebug(status bool) {
    is_debug = status
}

func IsDebug() bool {
    return is_debug
}

func Println(v ...interface{}) {
    if IsDebug() {
        log.Println(v...)
    }
}

func Printf(format string, v ...interface{}) {
    if IsDebug() {
        log.Printf(format, v...)
    }
}

