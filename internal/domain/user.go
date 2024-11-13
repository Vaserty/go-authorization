package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserAggregate struct {
	id                    uuid.UUID `validate:"required"`
	Username              string    `validate:"required,min=2,max=50"`
	Email                 string    `validate:"required,email"`
	isVerified            bool
	verifiedAt            *time.Time
	lastPasswordChangedAt *time.Time
	registeredAt          time.Time `validate:"required"`
	hashedPassword        []byte    `validate:"required"`
	isSuperUser           bool
}

func (u *UserAggregate) GetIdentifier() uuid.UUID {
	return u.id
}

func (u *UserAggregate) Validate() error {
	valdidate := validator.New()
	return valdidate.Struct(u)
}

func (u *UserAggregate) ChangedEmail(newEmail string) error {
	validate := validator.New()

	if err := validate.Var(newEmail, "required,email"); err != nil {
		return err
	}

	u.Email = newEmail
	return nil
}

func (u *UserAggregate) VerifyPassword(password string) bool {
	svc := NewPasswordDomainService()
	return svc.Verify(u.hashedPassword, password)
}

func (u *UserAggregate) setNewPassword(password string) error {
	svc := NewPasswordDomainService()

	hashedPassword, err := svc.Hash(password)

	if err != nil {
		return fmt.Errorf("error during hashed password: %v", err)
	}

	nowUTC := time.Now().UTC()
	u.hashedPassword = hashedPassword
	u.lastPasswordChangedAt = &nowUTC
	return nil
}

func (u *UserAggregate) ChangePassword(password string) error {
	return u.setNewPassword(password)
}

func (u *UserAggregate) ResetPassword() (string, error) {
	svc := NewPasswordDomainService()
	newPassword, err := svc.Generate()

	if err != nil {
		return newPassword, err
	}

	return newPassword, u.setNewPassword(newPassword)
}

func (u *UserAggregate) Verify() error {
	if u.isVerified {
		return errors.New("the verified user could not be verified")
	}

	u.isVerified = true
	now := time.Now().UTC()
	u.verifiedAt = &now
	return nil
}

func newUserAggregate(
	username string,
	email string,
	hashedPassword []byte,
	isSuperUser bool,
	isVerified bool,
) (*UserAggregate, error) {
	user := &UserAggregate{
		id:             uuid.New(),
		Username:       username,
		Email:          email,
		isVerified:     isVerified,
		registeredAt:   time.Now().UTC(),
		hashedPassword: hashedPassword,
		isSuperUser:    isSuperUser,
	}
	validationErr := user.Validate()
	return user, validationErr
}

func NewNormalUserAggregate(
	username string,
	email string,
	hashedPassword []byte,
) (*UserAggregate, error) {
	return newUserAggregate(
		username,
		email,
		hashedPassword,
		false,
		false,
	)
}

func NewSuperUserAggregate(
	username string,
	email string,
	hashedPassword []byte,
) (*UserAggregate, error) {
	return newUserAggregate(
		username,
		email,
		hashedPassword,
		true,
		true,
	)
}
