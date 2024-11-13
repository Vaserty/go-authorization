package config

type PasswordPolicy struct {
    minLength       int
    minUppercase    int
    minLowercase    int
    minDigits       int
}

type signCount struct {
	UpperCaseCount int
	lowerCaseCount int
	digitCount int
}

func (p PasswordPolicy) ValidatePassword(password string) error {
    if len(password) < p.minLength {
        return errors.New("password is too short")
    }

    var upperCaseCount, lowerCaseCount, digitCount int

    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            uppercaseCount++
        case unicode.IsLower(char):
            lowercaseCount++
        case unicode.IsDigit(char):
            digitCount++
    	}
	}

    if uppercaseCount < p.minUppercase {
        return errors.New("password must contain more uppercase letters")
    }

    if lowercaseCount < p.minLowercase {
        return errors.New("password must contain more lowercase letters")
    }

    if digitCount < p.minDigits {
        return errors.New("password must contain more digits")
    }