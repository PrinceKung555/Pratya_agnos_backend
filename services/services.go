package services

import (
	"math/rand"
	"strings"
)

func IsStrong(password string) bool {
	if len(password) < 6 {
		return false
	}

	hasLower, hasUpper, hasDigit := false, false, false
	if strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		hasLower = true
	}
	if strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		hasUpper = true
	}
	if strings.ContainsAny(password, "0123456789") {
		hasDigit = true
	}

	if !(hasLower && hasUpper && hasDigit) {
		return false
	}

	for i := 2; i < len(password); i++ {
		if password[i] == password[i-1] && password[i] == password[i-2] {
			return false
		}
	}

	return true
}

func ActionsNeededToMakeStrong(password string) int {
	if IsStrong(password) {
		return 0
	}

	hasLower, hasUpper, hasDigit := 1, 1, 1
	lowerCharset := "abcdefghijklmnopqrstuvwxyz"
	upperCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitCharset := "0123456789"
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if strings.ContainsAny(password, lowerCharset) {
		hasLower = 0
		charSet = strings.Replace(charSet, string(lowerCharset), "", -1)
	}
	if strings.ContainsAny(password, upperCharset) {
		hasUpper = 0
		charSet = strings.Replace(charSet, string(upperCharset), "", -1)
	}
	if strings.ContainsAny(password, digitCharset) {
		hasDigit = 0
		charSet = strings.Replace(charSet, string(digitCharset), "", -1)
	}

	if charSet == "" {
		charSet = lowerCharset + upperCharset + digitCharset
	}

	charSet = strings.Replace(charSet, string(password), "", -1)

	missingCriteria := hasLower + hasUpper + hasDigit

	if len(password) <= 6 {
		return max(missingCriteria, 6-len(password))
	}

	replacementsNeeded, replacementsPossible := 0, 0
	for i := 2; i < len(password); {
		if password[i] == password[i-1] && password[i] == password[i-2] {
			replacementsNeeded++
			newChar := generateRandomString(charSet)
			password = password[:i] + newChar + password[i+1:]
			i += 2
		} else {
			i++
		}
		replacementsPossible++
	}

	if len(password) <= 20 {
		return max(missingCriteria, replacementsNeeded)
	}

	deletionsNeeded := max(len(password)-20, 0)
	replacementsNeeded = max(replacementsNeeded-deletionsNeeded, 0)

	return deletionsNeeded + max(missingCriteria, replacementsNeeded)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateRandomString(charSet string) string {
	return string(charSet[rand.Intn(len(charSet))])
}
