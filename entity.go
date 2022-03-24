package user

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Entity ...
type Entity struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Pin      string `json:"-"`

	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Name ...
func (e Entity) Name() string {
	name := []string{e.FirstName}
	if e.LastName != "" {
		name = append(name, e.LastName)
	}
	return strings.Join(name, " ")
}

// HashingPassword ...
func (e *Entity) HashingPassword() error {
	b, err := bcrypt.GenerateFromPassword([]byte(e.Password), 8)
	if err != nil {
		return err
	}
	e.Password = string(b)
	return nil
}

// HashingPin ...
func (e *Entity) HashingPin() error {
	b, err := bcrypt.GenerateFromPassword([]byte(e.Pin), 8)
	if err != nil {
		return err
	}
	e.Pin = string(b)
	return nil
}

// ComparePassword ...
func (e Entity) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
	return err == nil
}

// Token ...
func (e Entity) Token(jwtKey string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": e.ID,
		"exp":     time.Now().Add(duration).Unix(),
	})
	return token.SignedString([]byte(jwtKey))
}
