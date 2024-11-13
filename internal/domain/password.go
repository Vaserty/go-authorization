package domain

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type PasswordValueObject struct {
	value string `validate:"required,min=8"`
}

func (p *PasswordValueObject) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}

	if !govalidator.HasLowerCase(p.value) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	if !govalidator.HasUpperCase(p.value) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	return nil
}

func (p *PasswordValueObject) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.value), bcrypt.DefaultCost)
}

func newPasswordValueObject(plaintextPassword string) (PasswordValueObject, error) {
	password := PasswordValueObject{value: plaintextPassword}
	err := password.Validate()
	return password, err
}
