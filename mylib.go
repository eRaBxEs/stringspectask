package mylib

import (
	"strings"
)

// testValidity : took a little bit over 20 mins
func testValidity(in string) (out bool) {
	if len(in) == 0 {
		return out
	}

	if !strings.ContainsAny(in, "-") {
		return out
	}

	for _, r := range in {

		if !strings.ContainsAny("abcdefghijklmnopqrstuvwxyz0123456789", strings.ToLower(string(r))) {
			return out
		}

	}

	out = true

	return out
}
