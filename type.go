package helper

import (
    "reflect"
    "strings"
)

// Typeof return the value data type
func Typeof(value interface{}) (typ string, kind string) {
    r := reflect.ValueOf(value)

    return strings.ReplaceAll(r.Type().String(), " ", ""), r.Type().Kind().String()
}

// IsSlice check if value is slice
func IsSlice(value interface{}) bool {
    _, k := Typeof(value)

    return k == "slice"
}

// IsArray Check if the value is an array, the default is non-strict mode
// Non-strict mode type is SLICE, Array, and map will return true.
// Returns True when only the array type
func IsArray(value interface{}, strict ...bool) bool {
    _, k := Typeof(value)
    if len(strict) > 0 && strict[0] {
        return k == "array"
    }

    return k == "slice" || k == "array" || k == "map"
}

// IsPtr Check if value is a pointer type
func IsPtr(value interface{}) bool {
    _, k := Typeof(value)

    return k == "ptr" || k == "unsafe.Pointer"
}
