package helper

import (
    "strings"
    "unicode"

    "github.com/spf13/cast"
)

// StrSizeToBytes converts strings like 1GB or 12 mb into an unsigned integer number of bytes
func StrSizeToBytes(sizeStr string) uint {
    sizeStr = strings.TrimSpace(sizeStr)
    lastChar := len(sizeStr) - 1
    multiplier := uint(1)

    if lastChar > 0 {
        if sizeStr[lastChar] == 'b' || sizeStr[lastChar] == 'B' {
            if lastChar > 1 {
                switch unicode.ToLower(rune(sizeStr[lastChar-1])) {
                case 'k':
                    multiplier = 1 << 10
                    sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
                case 'm':
                    multiplier = 1 << 20
                    sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
                case 'g':
                    multiplier = 1 << 30
                    sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
                default:
                    multiplier = 1
                    sizeStr = strings.TrimSpace(sizeStr[:lastChar])
                }
            }
        }
    }

    size := cast.ToInt(sizeStr)
    if size < 0 {
        size = 0
    }

    return SafeMul(uint(size), multiplier)
}