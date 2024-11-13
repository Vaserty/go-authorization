package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserAggregate struct {
	id                    uuid.UUID `validate:"required"`
	Username              string    `validate:"required,min=2,max=50"`
	Email                 string    `validate:"required,email"`
	lastPasswordChangedAt *time.Time
	registeredAt          time.Time `validate:"required"`
	hashedPassword        []byte    `validate:"required"`
}

func (u *UserAggregate) GetIdentifier() uuid.UUID {
	return u.id
}

func (u *UserAggregate) ChangedEmail(newEmail string) error {
	validate := validator.New()
	if err := validate.Var(newEmail, "required,email"); err != nil {
		return err
	}
	u.Email = newEmail
	return nil
}

func (u *UserAggregate) VerifyPassword(plaintextPassword string) bool {
	_, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	return err == nil
}

func (u *UserAggregate) ChangedPassword(plaintextPassword string) error {
	newPassword, err := newPasswordValueObject(plaintextPassword)

	if err != nil {
		return err
	}

	hashedPassword, err := newPassword.Hash()

	if err != nil {
		return fmt.Errorf("error during hashed password: %v", err)
	}

	nowUTC := time.Now().UTC()

	u.hashedPassword = hashedPassword
	u.lastPasswordChangedAt = &nowUTC

	return nil
}

func NewUserAggregate(Username string, Email string, Password PasswordValueObject) {
	return

}
