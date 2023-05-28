package clean

import "strings"

const MaxLength = 4096

func reject(s string, maxLength int) bool {
	if maxLength > 0 && len(s) > maxLength {
		return true
	}

	if strings.Contains(s, "${") || strings.Contains(s, "ldap://") {
		return true
	}

	return false
}
