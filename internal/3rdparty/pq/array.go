package pq

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"
	"strconv"
	"time"
)

// BoolArray represents a one-dimensional array of the PostgreSQL boolean type.
type BoolArray []bool

// Value implements the driver.Valuer interface.
func (a BoolArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be exactly two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1+2*n)

		for i := 0; i < n; i++ {
			b[2*i] = ','
			if a[i] {
				b[1+2*i] = 't'
			} else {
				b[1+2*i] = 'f'
			}
		}

		b[0] = '{'
		b[2*n] = '}'

		return string(b), nil
	}

	return "{}", nil
}

// ByteaArray represents a one-dimensional array of the PostgreSQL bytea type.
type ByteaArray [][]byte

// Value implements the driver.Valuer interface. It uses the "hex" format which
// is only supported on PostgreSQL 9.0 or newer.
func (a ByteaArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// 3*N bytes of hex formatting, and N-1 bytes of delimiters.
		size := 1 + 6*n
		for _, x := range a {
			size += hex.EncodedLen(len(x))
		}

		b := make([]byte, size)

		for i, s := 0, b; i < n; i++ {
			o := copy(s, `,"\\x`)
			o += hex.Encode(s[o:], a[i])
			s[o] = '"'
			s = s[o+1:]
		}

		b[0] = '{'
		b[size-1] = '}'

		return string(b), nil
	}

	return "{}", nil
}

// Float64Array represents a one-dimensional array of the PostgreSQL double
// precision type.
type Float64Array []float64

// Value implements the driver.Valuer interface.
func (a Float64Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendFloat(b, a[0], 'f', -1, 64)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendFloat(b, a[i], 'f', -1, 64)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// Float32Array represents a one-dimensional array of the PostgreSQL double
// precision type.
type Float32Array []float32

// Value implements the driver.Valuer interface.
func (a Float32Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendFloat(b, float64(a[0]), 'f', -1, 32)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendFloat(b, float64(a[i]), 'f', -1, 32)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// Int64Array represents a one-dimensional array of the PostgreSQL integer types.
type Int64Array []int64

// Value implements the driver.Valuer interface.
func (a Int64Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendInt(b, a[0], 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, a[i], 10)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// Int32Array represents a one-dimensional array of the PostgreSQL integer types.
type Int32Array []int32

// Value implements the driver.Valuer interface.
func (a Int32Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendInt(b, int64(a[0]), 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, int64(a[i]), 10)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// StringArray represents a one-dimensional array of the PostgreSQL character types.
type StringArray []string

// Value implements the driver.Valuer interface.
func (a StringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, []byte(a[0]))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, []byte(a[i]))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

type TimeArray []time.Time

// Value implements the driver.Valuer interface.
func (a TimeArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes, 30 bytes for time string
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 2+25*n+n-1)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, FormatTimestamp(a[0]))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, FormatTimestamp(a[i]))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

func appendArrayQuotedBytes(b, v []byte) []byte {
	b = append(b, '"')
	for {
		i := bytes.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}
