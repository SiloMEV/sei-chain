package cli

import (
	"strconv"
)

// ParseUint64 parses a string into a uint64
func ParseUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}
