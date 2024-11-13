package domain

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

type PasswordDomainService struct {
	passwordPolicy PasswordPolicy
}

func (p *PasswordDomainService) Validate(password string) error {
	if err := p.passwordPolicy.VerifyPolicyCompliance(password); err != nil {
		return fmt.Errorf("the password does not meet the required policy: %v", err)
	}
	return nil
}

func (p *PasswordDomainService) Generate() (string, error) {
	plaintextPassword, err := password.Generate(
		p.passwordPolicy.minLength,
		p.passwordPolicy.minDigits,
		1,
		true,
		false,
	)

	if err != nil {
		return plaintextPassword,
			fmt.Errorf("error during generate new password: %v", err)
	}

	newPassword := plaintextPassword + generateRandomUppercase(p.passwordPolicy.minUpperCase)
	validateErr := p.Validate(newPassword)
	return newPassword, validateErr
}

func (p *PasswordDomainService) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (p *PasswordDomainService) Verify(hashedPassword []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}

func NewPasswordDomainService() *PasswordDomainService {
	return &PasswordDomainService{passwordPolicy: newPasswordPolicy()}
}
