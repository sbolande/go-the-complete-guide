package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(fName, lName, birthDate string) (*User, error) {
	if fName == "" || lName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	return &User{
		firstName: fName,
		lastName: lName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}