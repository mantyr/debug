package debug

import (
    "log"
    "strings"
)

var is_debug bool
var is_level map[string]bool = make(map[string]bool)

// No safe. Do not change the values if you do not believe that non-concurrent access
func SetDebug(status bool) {
    is_debug = status
}

// No safe. Do not change the values if you do not believe that non-concurrent access
func SetLevelString(levels string) {
    for _, level := range strings.Split(levels, " ") {
        if level != "" {
            is_level[level] = true
        }
    }
}

// No safe. Do not change the values if you do not believe that non-concurrent access
func SetLevel(levels ...string) {
    for _, level := range levels {
        is_level[level] = true
    }
}

// No safe. Do not change the values if you do not believe that non-concurrent access
func DeleteLevel(level string) {
    delete(is_level, level)
}

func IsDebug() bool {
    return is_debug
}

func IsLevel(level string) bool {
    if v, ok := is_level[level]; ok && v {
        return true
    }
    return false
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

func LevelPrintln(level string, v ...interface{}) {
    if IsDebug() && ((level != "db" && IsLevel("all")) || IsLevel(level)) {
        log.Println(v...)
    }
}

func LevelPrintf(level string, format string, v ...interface{}) {
    if IsDebug() && ((level != "db" && IsLevel("all")) || IsLevel(level)) {
        log.Printf(format, v...)
    }
}
