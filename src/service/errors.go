package service

import "errors"

// ErrInternal is an unexplained error
var ErrInternal = errors.New("Internal Service Error")

// ErrNonexistentUser is returned when the user doesn't exists
var ErrNonexistentUser = errors.New("Nonexistent User Error")

// ErrNonexistentGame is returned when the requested game doesn't exist
var ErrNonexistentGame = errors.New("Nonexistent Game Error")

// ErrForbidden is returned when command is not allowed to execute
var ErrForbidden = errors.New("Command not allowed or the resource belongs to someone else")

// ErrOutsideBoardBoundaries is returned when command outside boundaries
var ErrOutsideBoardBoundaries = errors.New("Cell outside board")
