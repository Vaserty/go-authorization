package domain

import (
	"math/rand"
	"unicode"
)

type SignCount struct {
	upperCaseCount int
	lowerCaseCount int
	digitCount     int
}

func charTypeCounts(password string) SignCount {
	var upperCaseCount, lowerCaseCount, digitCount int

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upperCaseCount++
		case unicode.IsLower(char):
			lowerCaseCount++
		case unicode.IsDigit(char):
			digitCount++
		}
	}

	return SignCount{
		upperCaseCount: upperCaseCount,
		lowerCaseCount: lowerCaseCount,
		digitCount:     digitCount,
	}
}

func generateRandomUppercase(numUpper int) string {
	uppercaseCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if numUpper <= 0 {
		return ""
	}

	result := make([]byte, numUpper)

	for i := 0; i < numUpper; i++ {
		index := rand.Intn(len(uppercaseCharset))
		result[i] = uppercaseCharset[index]
	}

	return string(result)
}
