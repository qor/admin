package admin

import (
	"fmt"
	"unicode/utf8"
)

// EncodePath returns the safe encoding of the given module path.
// It fails if the module path is invalid.
func EncodePath(path string) (encoding string, err error) {
	haveUpper := false
	for _, r := range path {
		if r == '!' || r >= utf8.RuneSelf {
			// This should be disallowed by CheckPath, but diagnose anyway.
			// The correctness of the encoding loop below depends on it.
			return "", fmt.Errorf("internal error: inconsistency in EncodePath")
		}
		if 'A' <= r && r <= 'Z' {
			haveUpper = true
		}
	}

	if !haveUpper {
		return path, nil
	}

	var buf []byte
	for _, r := range path {
		if 'A' <= r && r <= 'Z' {
			buf = append(buf, '!', byte(r+'a'-'A'))
		} else {
			buf = append(buf, byte(r))
		}
	}
	return string(buf), nil
}
