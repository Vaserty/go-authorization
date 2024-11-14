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
	counts := SignCount{}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			counts.upperCaseCount++
		case unicode.IsLower(char):
			counts.lowerCaseCount++
		case unicode.IsDigit(char):
			counts.digitCount++
		}
	}

	return counts
}

func generateRandomUppercase(upperCaseCount int) string {
	uppercaseCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if upperCaseCount <= 0 {
		return ""
	}

	result := make([]byte, upperCaseCount)

	for i := 0; i < upperCaseCount; i++ {
		index := rand.Intn(len(uppercaseCharset))
		result[i] = uppercaseCharset[index]
	}

	return string(result)
}
