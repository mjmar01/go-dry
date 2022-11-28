package options

import "fmt"

const (
	errNotExported = "field '%s' of Option must be exported"
	errNoDefault   = "field '%s' has no '%s' tag specified"
	errCantSet     = "field '%s' can not be set"

	errConvert  = "default value '%s' can not be converted to %s"
	errOverflow = "default value overflows values data size"
)

// ErrNotExported is the error that a field of the option structs type is not exported
type ErrNotExported struct {
	field string
}

func (e ErrNotExported) Error() string {
	return fmt.Sprintf(errNotExported, e.field)
}

// ErrNoDefault is the error that a field of the option structs type has no default tag
type ErrNoDefault struct {
	field string
	key   string
}

func (e ErrNoDefault) Error() string {
	return fmt.Sprintf(errNoDefault, e.field, e.key)
}

// ErrCantSet is the error that a field of the option structs type can not be set by reflect
type ErrCantSet struct {
	field string
}

func (e ErrCantSet) Error() string {
	return fmt.Sprintf(errCantSet, e.field)
}

// ErrConvert is the error that the specified default value can not be converted by strconv
type ErrConvert struct {
	// StrConvErr is the original error returned by strconv
	StrConvErr error

	value string
	kind  string
}

func (e ErrConvert) Error() string {
	return fmt.Sprintf(errConvert, e.value, e.kind)
}

// ErrOverflow is the error that the specified default value overflows the data size of the field type
type ErrOverflow struct {
}

func (e ErrOverflow) Error() string {
	return errOverflow
}
