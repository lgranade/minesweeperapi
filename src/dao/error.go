package dao

import (
	"github.com/lib/pq"
)

// GetPQError returns postgresql native error
func GetPQError(err error) (*pq.Error, bool) {
	pqError, ok := err.(*pq.Error)
	return pqError, ok
}

// IsPQIntegrityViolationError returns true if the error is because of one of the following:
// "23000": "integrity_constraint_violation",
// "23001": "restrict_violation",
// "23502": "not_null_violation",
// "23503": "foreign_key_violation",
// "23505": "unique_violation",
// "23514": "check_violation",
// "23P01": "exclusion_violation",
func IsPQIntegrityViolationError(err error) bool {
	if pqError, ok := GetPQError(err); ok {
		return pqError.Code.Class() == "23"
	}
	return false
}
