package options

import (
	"errors"
	"reflect"
	"strconv"
)

const defaultKey = "default"

// Defaults fills a given structs fields with specified default values unless given explicitly.
// Supports types: string, int, uint, float, complex and bool
//
// Example:
func Defaults[O any](option *O) (*O, error) {
	if option == nil {
		option = new(O)
	}

	optionType := reflect.TypeOf(*option)
	optionValue := reflect.ValueOf(option).Elem()
	for i := 0; i < optionType.NumField(); i++ {
		field := optionType.Field(i)
		if !field.IsExported() {
			return nil, ErrNotExported{field: field.Name}
		}

		defString, ok := field.Tag.Lookup(defaultKey)
		if !ok {
			return nil, ErrNoDefault{
				field: field.Name,
				key:   defaultKey,
			}
		}

		value := optionValue.FieldByName(field.Name)
		if !value.CanSet() {
			return nil, ErrCantSet{field: field.Name}
		}
		switch value.Kind() {
		case reflect.String:
			if value.String() != "" {
				break
			}
			value.SetString(defString)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if value.Int() != 0 {
				break
			}
			defInt, err := strconv.ParseInt(defString, 0, 64)
			if err != nil {
				return nil, ErrConvert{
					StrConvErr: err,
					value:      defString,
					kind:       "int",
				}
			}
			if value.OverflowInt(defInt) {
				return nil, ErrOverflow{}
			}
			value.SetInt(defInt)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if value.Uint() != 0 {
				break
			}
			defUint, err := strconv.ParseUint(defString, 0, 64)
			if err != nil {
				return nil, ErrConvert{
					StrConvErr: err,
					value:      defString,
					kind:       "uint",
				}
			}
			if value.OverflowUint(defUint) {
				return nil, errors.New(errOverflow)
			}
			value.SetUint(defUint)
		case reflect.Float32, reflect.Float64:
			if value.Float() != 0 {
				break
			}
			defFloat, err := strconv.ParseFloat(defString, 64)
			if err != nil {
				return nil, ErrConvert{
					StrConvErr: err,
					value:      defString,
					kind:       "float",
				}
			}
			if value.OverflowFloat(defFloat) {
				return nil, ErrOverflow{}
			}
			value.SetFloat(defFloat)
		case reflect.Complex64, reflect.Complex128:
			if value.Complex() != 0+0i {
				break
			}
			defComplex, err := strconv.ParseComplex(defString, 128)
			if err != nil {
				return nil, ErrConvert{
					StrConvErr: err,
					value:      defString,
					kind:       "complex",
				}
			}
			if value.OverflowComplex(defComplex) {
				return nil, ErrOverflow{}
			}
			value.SetComplex(defComplex)
		case reflect.Bool:
			defBool, err := strconv.ParseBool(defString)
			if err != nil {
				return nil, ErrConvert{
					StrConvErr: err,
					value:      defString,
					kind:       "bool",
				}
			}
			value.SetBool(defBool)
		}
	}

	return option, nil
}
