package errors

import (
	"github.com/pkg/errors"
)

//BeerErr estructura de error
type BeerErr struct {
	error
}

//WrapUnreacheableBeerErr enmascara error de funcines de thrid's
func WrapUnreacheableBeerErr(err error, format string, args ...interface{}) error {
	return &BeerErr{errors.Wrapf(err, format, args...)}
}

// NewUnreacheableBeerErr returns an error which satisfies IsDataUnreacheable()
func NewUnreacheableBeerErr(format string, args ...interface{}) error {
	return &BeerErr{errors.Errorf(format, args...)}
}

// IsUnreacheableBeerErr reports whether err was created with BeerErr() or
// NewUnreacheable()
func IsUnreacheableBeerErr(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*BeerErr)
	return ok
}
