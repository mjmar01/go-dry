package options

import (
	"github.com/mjmar01/go-dry/pkg/options"
	. "github.com/stretchr/testify/assert"
	"testing"
)

var err error

func TestOptions(t *testing.T) {
	assert := New(t)
	type SomeOption struct {
		StringA  string    `default:"abc"`
		StringB  string    `default:"abc"`
		Int      int       `default:"5"`
		Negative int       `default:"-5"`
		Uint     uint16    `default:"0"`
		Float    float32   `default:"1.5"`
		Complex  complex64 `default:"1+1i"`
	}

	opts := &SomeOption{
		Int:     10,
		StringB: "def",
	}

	opts, err = options.Defaults(opts)

	assert.Nil(err)
	assert.Equal(10, opts.Int)
	assert.Equal("abc", opts.StringA)
	assert.Equal("def", opts.StringB)
	assert.Equal(-5, opts.Negative)
	assert.Equal(uint16(0), opts.Uint)
	assert.Equal(float32(1.5), opts.Float)
	assert.Equal(complex64(1+1i), opts.Complex)
}

func TestNil(t *testing.T) {
	assert := New(t)

	type SomeOption struct {
		Option int `default:"5"`
	}

	var opts *SomeOption

	opts, err = options.Defaults(opts)

	assert.Nil(err)
	assert.Equal(5, opts.Option)
}

func TestExportErr(t *testing.T) {
	assert := New(t)

	type SomeOption struct {
		OptionA string `default:"abc"`
		optionB string `default:"abc"`
	}

	opts := &SomeOption{}

	opts, err = options.Defaults(opts)

	assert.IsType(options.ErrNotExported{}, err)
	assert.Equal("field 'optionB' of Option must be exported", err.Error())
}

func TestTagErr(t *testing.T) {
	assert := New(t)

	type SomeOption struct {
		OptionA string `default:"abc"`
		OptionB string
	}

	opts := &SomeOption{}

	opts, err = options.Defaults(opts)

	assert.IsType(options.ErrNoDefault{}, err)
	assert.Equal("field 'OptionB' has no 'default' tag specified", err.Error())
}

func TestConvertErr(t *testing.T) {
	assert := New(t)

	type SomeOption[T any] struct {
		Option T `default:"abc"`
	}

	_, err = options.Defaults(&SomeOption[int]{})
	assert.IsType(options.ErrConvert{}, err)
	assert.Equal("default value 'abc' can not be converted to int", err.Error())

	_, err = options.Defaults(&SomeOption[uint]{})
	assert.IsType(options.ErrConvert{}, err)
	assert.Equal("default value 'abc' can not be converted to uint", err.Error())

	_, err = options.Defaults(&SomeOption[float32]{})
	assert.IsType(options.ErrConvert{}, err)
	assert.Equal("default value 'abc' can not be converted to float", err.Error())

	_, err = options.Defaults(&SomeOption[complex64]{})
	assert.IsType(options.ErrConvert{}, err)
	assert.Equal("default value 'abc' can not be converted to complex", err.Error())

	_, err = options.Defaults(&SomeOption[bool]{})
	assert.IsType(options.ErrConvert{}, err)
	assert.Equal("default value 'abc' can not be converted to bool", err.Error())
}
