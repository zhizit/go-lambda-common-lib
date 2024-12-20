package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNullBool(t *testing.T) {
	t.Run("nil でない場合、sql.NullBool が有効な値で返されること", func(t *testing.T) {
		b := true
		result := ToNullBool(&b)
		assert.True(t, result.Valid)
		assert.Equal(t, b, result.Bool)
	})

	t.Run("nil の場合、sql.NullBool が無効な値で返されること", func(t *testing.T) {
		var b *bool
		result := ToNullBool(b)
		assert.False(t, result.Valid)
	})
}

func TestToNullInt64(t *testing.T) {
	t.Run("nil でない場合、sql.NullInt64 が有効な値で返されること", func(t *testing.T) {
		i := int64(42)
		result := ToNullInt64(&i)
		assert.True(t, result.Valid)
		assert.Equal(t, i, result.Int64)
	})

	t.Run("nil の場合、sql.NullInt64 が無効な値で返されること", func(t *testing.T) {
		var i *int64
		result := ToNullInt64(i)
		assert.False(t, result.Valid)
	})
}

func TestToNullFloat64(t *testing.T) {
	t.Run("nil でない場合、sql.NullFloat64 が有効な値で返されること", func(t *testing.T) {
		f := 42.42
		result := ToNullFloat64(&f)
		assert.True(t, result.Valid)
		assert.Equal(t, f, result.Float64)
	})

	t.Run("nil の場合、sql.NullFloat64 が無効な値で返されること", func(t *testing.T) {
		var f *float64
		result := ToNullFloat64(f)
		assert.False(t, result.Valid)
	})
}

func TestToNullString(t *testing.T) {
	t.Run("nil でない場合、sql.NullString が有効な値で返されること", func(t *testing.T) {
		s := "test string"
		result := ToNullString(&s)
		assert.True(t, result.Valid)
		assert.Equal(t, s, result.String)
	})

	t.Run("nil の場合、sql.NullString が無効な値で返されること", func(t *testing.T) {
		var s *string
		result := ToNullString(s)
		assert.False(t, result.Valid)
	})
}
