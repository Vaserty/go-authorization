package domain

import (
	"errors"
)

type PasswordPolicy struct {
	minLength    int
	minUpperCase int
	minLowerCase int
	minDigits    int
}

func (p PasswordPolicy) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"minLength":    p.minLength,
		"minUpperCase": p.minUpperCase,
		"minLowerCase": p.minLowerCase,
		"minDigits":    p.minDigits,
	}
}

func (p PasswordPolicy) VerifyPolicyCompliance(password string) error {
	if len(password) < p.minLength {
		return errors.New("password is too short")
	}

	passwordCharCounts := charTypeCounts(password)

	if passwordCharCounts.upperCaseCount < p.minUpperCase {
		return errors.New("password must contain more uppercase letters")
	}

	if passwordCharCounts.lowerCaseCount < p.minLowerCase {
		return errors.New("password must contain more lowercase letters")
	}

	if passwordCharCounts.digitCount < p.minDigits {
		return errors.New("password must contain more digits")
	}

	return nil
}

func newPasswordPolicy() PasswordPolicy {
	return PasswordPolicy{
		minLength:    12,
		minUpperCase: 2,
		minLowerCase: 2,
		minDigits:    2,
	}
}
