package flags

import (
	"fmt"
	"regexp"
	"strconv"
)

const trueStr = "true"
const falseStr = "false"

// NewString create a string option
func NewString(long string, short rune, description string, defaultValue string) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &String{DefaultValue: defaultValue},
	}
}

// StringValue return the value of a String option
func StringValue(option *Option) (string, error) {
	if value, ok := option.Value.(*String); ok {
		if value.Value == "" {
			return value.DefaultValue, nil
		}

		return value.Value, nil
	}

	return "", fmt.Errorf("Not a string option")
}

// Set set the value
func (val *String) Set(value string) error {
	val.Value = value

	return nil
}

// String representation of the value
func (val *String) String() string {
	v := val.Value
	if v == "" {
		v = val.DefaultValue
	}

	return v
}

// DefaultValueString string representation of the default value
func (val *String) DefaultValueString() string {
	return val.DefaultValue
}

// IsBoolValue check if value is boolean
func (val *String) IsBoolValue() bool {
	return false
}

// NewBool create a bool option
func NewBool(long string, short rune, description string, defaultValue bool) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Bool{DefaultValue: defaultValue},
	}
}

// BoolValue return the value of a Bool option
func BoolValue(option *Option) (bool, error) {
	if value, ok := option.Value.(*Bool); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return false, fmt.Errorf("Not a boolean option")
}

// Set set the value
func (val *Bool) Set(value string) error {
	rx := regexp.MustCompile("(?i)true")
	val.Value = value == "" || rx.MatchString(value)
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Bool) String() string {
	value := falseStr

	if val.ValueSet {
		if val.Value {
			value = trueStr
		}
	} else {
		if val.DefaultValue {
			value = trueStr
		}
	}

	return value
}

// DefaultValueString string representation of the default value
func (val *Bool) DefaultValueString() string {
	defaultValue := falseStr
	if val.DefaultValue {
		defaultValue = trueStr
	}

	return defaultValue
}

// IsBoolValue check if value is boolean
func (val *Bool) IsBoolValue() bool {
	return true
}

// NewInt create an int option
func NewInt(long string, short rune, description string, defaultValue int) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Int{DefaultValue: defaultValue},
	}
}

// IntValue return the value of an Int option
func IntValue(option *Option) (int, error) {
	if value, ok := option.Value.(*Int); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return 0, fmt.Errorf("Not an int option")
}

// Set set the value
func (val *Int) Set(value string) error {
	intVal, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return err
	}

	val.Value = int(intVal)
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Int) String() string {
	if val.ValueSet {
		return fmt.Sprintf("%d", val.Value)
	}

	return fmt.Sprintf("%d", val.DefaultValue)
}

// DefaultValueString string representation of the default value
func (val *Int) DefaultValueString() string {
	return fmt.Sprintf("%d", val.DefaultValue)
}

// IsBoolValue check if value is boolean
func (val *Int) IsBoolValue() bool {
	return false
}

// NewInt64 create an Int64 option
func NewInt64(long string, short rune, description string, defaultValue int64) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Int64{DefaultValue: defaultValue},
	}
}

// Int64Value return the value of an Int option
func Int64Value(option *Option) (int64, error) {
	if value, ok := option.Value.(*Int64); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return 0, fmt.Errorf("Not an int64 option")
}

// Set set the value
func (val *Int64) Set(value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}

	val.Value = intVal
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Int64) String() string {
	if val.ValueSet {
		return fmt.Sprintf("%d", val.Value)
	}

	return fmt.Sprintf("%d", val.DefaultValue)
}

// DefaultValueString string representation of the default value
func (val *Int64) DefaultValueString() string {
	return fmt.Sprintf("%d", val.DefaultValue)
}

// IsBoolValue check if value is boolean
func (val *Int64) IsBoolValue() bool {
	return false
}

// NewFloat32 create an Float32 option
func NewFloat32(long string, short rune, description string, defaultValue float32) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Float32{DefaultValue: defaultValue},
	}
}

// Float32Value return the value of a Float32 option
func Float32Value(option *Option) (float32, error) {
	if value, ok := option.Value.(*Float32); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return 0, fmt.Errorf("Not a float32 option")
}

// Set set the value
func (val *Float32) Set(value string) error {
	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}

	val.Value = float32(floatVal)
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Float32) String() string {
	if val.ValueSet {
		return fmt.Sprintf("%f", val.Value)
	}

	return fmt.Sprintf("%f", val.DefaultValue)
}

// DefaultValueString string representation of the default value
func (val *Float32) DefaultValueString() string {
	return fmt.Sprintf("%f", val.DefaultValue)
}

// IsBoolValue check if value is boolean
func (val *Float32) IsBoolValue() bool {
	return false
}

// NewFloat64 create an Float64 option
func NewFloat64(long string, short rune, description string, defaultValue float64) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Float64{DefaultValue: defaultValue},
	}
}

// Float64Value return the value of a Float64 option
func Float64Value(option *Option) (float64, error) {
	if value, ok := option.Value.(*Float64); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return 0, fmt.Errorf("Not a float64 option")
}

// Set set the value
func (val *Float64) Set(value string) error {
	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}

	val.Value = floatVal
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Float64) String() string {
	if val.ValueSet {
		return fmt.Sprintf("%f", val.Value)
	}

	return fmt.Sprintf("%f", val.DefaultValue)
}

// DefaultValueString string representation of the default value
func (val *Float64) DefaultValueString() string {
	return fmt.Sprintf("%f", val.DefaultValue)
}

// IsBoolValue check if value is boolean
func (val *Float64) IsBoolValue() bool {
	return false
}

// NewUint create an Float64 option
func NewUint(long string, short rune, description string, defaultValue uint) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Uint{DefaultValue: defaultValue},
	}
}

// UintValue return the value of a Float64 option
func UintValue(option *Option) (uint, error) {
	if value, ok := option.Value.(*Uint); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return 0, fmt.Errorf("Not an uint option")
}

// Set set the value
func (val *Uint) Set(value string) error {
	uintVal, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return err
	}

	val.Value = uint(uintVal)
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Uint) String() string {
	if val.ValueSet {
		return fmt.Sprintf("%d", val.Value)
	}

	return fmt.Sprintf("%d", val.DefaultValue)
}

// DefaultValueString string representation of the default value
func (val *Uint) DefaultValueString() string {
	return fmt.Sprintf("%d", val.DefaultValue)
}

// IsBoolValue check if value is boolean
func (val *Uint) IsBoolValue() bool {
	return false
}

// NewUint64 create an Float64 option
func NewUint64(long string, short rune, description string, defaultValue uint64) *Option {
	return &Option{
		Long:        long,
		Short:       short,
		Description: description,
		Value:       &Uint64{DefaultValue: defaultValue},
	}
}

// Uint64Value return the value of a Float64 option
func Uint64Value(option *Option) (uint64, error) {
	if value, ok := option.Value.(*Uint64); ok {
		if value.ValueSet {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return 0, fmt.Errorf("Not an uint64 option")
}

// Set set the value
func (val *Uint64) Set(value string) error {
	uintVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	val.Value = uintVal
	val.ValueSet = true

	return nil
}

// String representation of the value
func (val *Uint64) String() string {
	if val.ValueSet {
		return fmt.Sprintf("%d", val.Value)
	}

	return fmt.Sprintf("%d", val.DefaultValue)
}

// DefaultValueString string representation of the default value
func (val *Uint64) DefaultValueString() string {
	return fmt.Sprintf("%d", val.DefaultValue)
}

// IsBoolValue check if value is boolean
func (val *Uint64) IsBoolValue() bool {
	return false
}
