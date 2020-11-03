package service

import "errors"

// ErrInternal is an unexplained error
var ErrInternal = errors.New("Internal Service Error")

// ErrNonexistentUser is returned when the user doesn't exists
var ErrNonexistentUser = errors.New("Nonexistent User Error")
