package flags

// String string option value (and default value)
type String struct {
	Value        string
	DefaultValue string
}

// Bool bool option value (and default value)
type Bool struct {
	Value        bool
	DefaultValue bool
	ValueSet     bool
}

// Int integer option value (and default value)
type Int struct {
	Value        int
	DefaultValue int
	ValueSet     bool
}

// Int64 64-bits option integer value (and default value)
type Int64 struct {
	Value        int64
	DefaultValue int64
	ValueSet     bool
}

// Float32 float32 option value (and default value)
type Float32 struct {
	Value        float32
	DefaultValue float32
	ValueSet     bool
}

// Float64 float64 option value (and default value)
type Float64 struct {
	Value        float64
	DefaultValue float64
	ValueSet     bool
}

// Uint unsigned integer option value (and default value)
type Uint struct {
	Value        uint
	DefaultValue uint
	ValueSet     bool
}

// Uint64 64-bits unsigned integer option value (and default value)
type Uint64 struct {
	Value        uint64
	DefaultValue uint64
	ValueSet     bool
}
