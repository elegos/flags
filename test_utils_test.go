package flags

type testStringWriter struct {
	Value string
}

// io.Writer
func (sw *testStringWriter) Write(p []byte) (n int, err error) {
	sw.Value += string(p)

	return len(p), nil
}
