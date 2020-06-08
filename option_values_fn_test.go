package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewString(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := "default"

	opt := NewString(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := StringValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestStringValueValueSet(t *testing.T) {
	defaultValue := "default"
	value := "value"

	opt := NewString("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := StringValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, value, val)
}

func TestStringString(t *testing.T) {
	value := "value"

	opt := NewString("", EmptyShort, "", "")
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestStringDefaultValueString(t *testing.T) {
	opt := NewString("", EmptyShort, "", "default")
	assert.Equal(t, "default", opt.Value.DefaultValueString())
}

func TestStringIsBoolValue(t *testing.T) {
	opt := NewString("", EmptyShort, "", "")
	assert.False(t, opt.Value.IsBoolValue())
}

func TestNewBool(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := false

	opt := NewBool(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := BoolValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestBoolValueValueSet(t *testing.T) {
	defaultValue := false
	value := "true"

	opt := NewBool("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := BoolValue(opt)
	assert.NoError(t, err)
	assert.True(t, val)
}

func TestBoolString(t *testing.T) {
	value := "true"

	opt := NewBool("", EmptyShort, "", false)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestBoolDefaultValueString(t *testing.T) {
	opt := NewBool("", EmptyShort, "", false)
	assert.Equal(t, "false", opt.Value.DefaultValueString())

	opt = NewBool("", EmptyShort, "", true)
	assert.Equal(t, "true", opt.Value.DefaultValueString())
}

func TestBoolIsBoolValue(t *testing.T) {
	opt := NewBool("", EmptyShort, "", false)
	assert.True(t, opt.Value.IsBoolValue())
}

func TestNewInt(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := 27

	opt := NewInt(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := IntValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestIntValueValueSet(t *testing.T) {
	defaultValue := 0
	value := "42"

	opt := NewInt("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := IntValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestIntString(t *testing.T) {
	value := "42"

	opt := NewInt("", EmptyShort, "", 42)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestIntDefaultValueString(t *testing.T) {
	opt := NewInt("", EmptyShort, "", 42)
	assert.Equal(t, "42", opt.Value.DefaultValueString())
}

func TestIntIsBoolValue(t *testing.T) {
	opt := NewInt("", EmptyShort, "", 0)
	assert.False(t, opt.Value.IsBoolValue())
}

func TestNewInt64(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := int64(27)

	opt := NewInt64(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := Int64Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestInt64ValueValueSet(t *testing.T) {
	defaultValue := int64(0)
	value := "42"

	opt := NewInt64("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := Int64Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, int64(42), val)
}

func TestInt64String(t *testing.T) {
	value := "42"

	opt := NewInt64("", EmptyShort, "", 42)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestInt64DefaultValueString(t *testing.T) {
	opt := NewInt64("", EmptyShort, "", 42)
	assert.Equal(t, "42", opt.Value.DefaultValueString())
}

func TestInt64IsBoolValue(t *testing.T) {
	opt := NewInt64("", EmptyShort, "", 0)
	assert.False(t, opt.Value.IsBoolValue())
}

func TestNewFloat32(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := float32(2.4)

	opt := NewFloat32(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := Float32Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestFloat32ValueValueSet(t *testing.T) {
	defaultValue := float32(0)
	value := "18.600000"

	opt := NewFloat32("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := Float32Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, float32(18.6), val)
}

func TestFloat32String(t *testing.T) {
	value := "18.600000"

	opt := NewFloat32("", EmptyShort, "", 18.6)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestFloat32DefaultValueString(t *testing.T) {
	opt := NewFloat32("", EmptyShort, "", 18.6)
	assert.Equal(t, "18.600000", opt.Value.DefaultValueString())
}

func TestFloat32IsBoolValue(t *testing.T) {
	opt := NewFloat32("", EmptyShort, "", 0)
	assert.False(t, opt.Value.IsBoolValue())
}

func TestNewFloat64(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := float64(2.4)

	opt := NewFloat64(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := Float64Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestFloat64ValueValueSet(t *testing.T) {
	defaultValue := float64(0)
	value := "12.000000"

	opt := NewFloat64("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := Float64Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, float64(12.0), val)
}

func TestFloat64String(t *testing.T) {
	value := "18.600000"

	opt := NewFloat64("", EmptyShort, "", 18.6)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestFloat64DefaultValueString(t *testing.T) {
	opt := NewFloat64("", EmptyShort, "", 18.6)
	assert.Equal(t, "18.600000", opt.Value.DefaultValueString())
}

func TestFloat64IsBoolValue(t *testing.T) {
	opt := NewFloat64("", EmptyShort, "", 0)
	assert.False(t, opt.Value.IsBoolValue())
}

func TestNewUint(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := uint(42)

	opt := NewUint(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := UintValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestUintValueValueSet(t *testing.T) {
	defaultValue := uint(0)
	value := "12"

	opt := NewUint("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := UintValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, uint(12), val)
}

func TestUintString(t *testing.T) {
	value := "18"

	opt := NewUint("", EmptyShort, "", 18)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestUintDefaultValueString(t *testing.T) {
	opt := NewUint("", EmptyShort, "", 18)
	assert.Equal(t, "18", opt.Value.DefaultValueString())
}

func TestUintIsBoolValue(t *testing.T) {
	opt := NewUint("", EmptyShort, "", 0)
	assert.False(t, opt.Value.IsBoolValue())
}

func TestNewUint64(t *testing.T) {
	long := "option"
	short := 'o'
	description := "description"
	defaultValue := uint64(42)

	opt := NewUint64(long, short, description, defaultValue)
	assert.Equal(t, long, opt.Long)
	assert.Equal(t, short, opt.Short)
	assert.Equal(t, description, opt.Description)
	value, err := Uint64Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, defaultValue, value)
}

func TestUint64ValueValueSet(t *testing.T) {
	defaultValue := uint64(0)
	value := "12"

	opt := NewUint64("", EmptyShort, "", defaultValue)
	assert.NoError(t, opt.Value.Set(value))

	val, err := Uint64Value(opt)
	assert.NoError(t, err)
	assert.Equal(t, uint64(12), val)
}

func TestUint64String(t *testing.T) {
	value := "18"

	opt := NewUint64("", EmptyShort, "", 18)
	assert.NoError(t, opt.Value.Set(value))
	assert.Equal(t, value, opt.Value.String())
}

func TestUint64DefaultValueString(t *testing.T) {
	opt := NewUint64("", EmptyShort, "", 18)
	assert.Equal(t, "18", opt.Value.DefaultValueString())
}

func TestUint64IsBoolValue(t *testing.T) {
	opt := NewUint64("", EmptyShort, "", 0)
	assert.False(t, opt.Value.IsBoolValue())
}
