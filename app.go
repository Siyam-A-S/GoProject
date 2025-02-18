package main

import (
	"errors"
	"strings"
)

var (
	ErrNameRequired  = errors.New("first name and last name are required")
	ErrEmailRequired = errors.New("email address is required")
	ErrInavlidEmail  = errors.New("email address is invalid")
)

type User struct {
	id        int
	firstName string
	lastName  string
	emails    []Email
}

func NewUser(firstName string, lastName string, emailAddress string) (User, error) {
	if firstName == "" && lastName == "" {
		return User{}, ErrNameRequired
	}

	email, err := NewEmail(emailAddress, true)
	if err != nil {
		return User{}, err
	}

	return User{
		firstName: firstName,
		lastName:  lastName,
		emails:    []Email{email},
	}, nil
}

type Email struct {
	address string
	primary bool
}

func NewEmail(address string, primary bool) (Email, error) {
	if address == "" {
		return Email{}, ErrEmailRequired
	}
	if !strings.Contains(address, "@") {
		return Email{}, ErrInavlidEmail
	}
	return Email{
		address: address,
		primary: primary,
	}, nil
}
