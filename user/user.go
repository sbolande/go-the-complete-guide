package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName string
	BirthDate string
	CreatedAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewUser(fName, lName, birthDate string) (*User, error) {
	if fName == "" || lName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	return &User{
		FirstName: fName,
		LastName: lName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}