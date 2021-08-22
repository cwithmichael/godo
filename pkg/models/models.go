package models

import (
	"errors"
	"time"
)

var (
	//ErrNoRecord is returned with no matching records are found
	ErrNoRecord = errors.New("models: no matching record found")
	//ErrInvalidCredentials is returned when the user provides bad creds
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	//ErrDuplicateEmail is returned whe na user tries to provide an email that's in use
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

//Todo represents a todo object in the app
type Todo struct {
	ID             int
	Title          string
	Content        string
	Created        time.Time
	CompletionDate time.Time
	Completed      bool
	UserID         int
}

//User represents a user of the app
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
