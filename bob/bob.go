// Package bob provides a function to handle different 
// string manipulations and branching 
package bob

import (
	"strings"
)

// Hey receives a string and returns another given some rules
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	var (
		isLower    = false
		isUpper    = false
		isQuestion = false
	)

	if strings.ToUpper(remark) == remark {
		isUpper = true
	}

	if strings.ToLower(remark) == remark {
		isLower = true
	}

	if remark[len(remark)-1] == '?' {
		isQuestion = true
	}

	if isUpper && isLower {
		if isQuestion {
			return "Sure."
		}
		return "Whatever."
	}

	if isUpper {
		if isQuestion {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
