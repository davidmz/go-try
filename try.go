package try

// Check is an advanced version of It that returns Result value. Result allows
// to test/process error before throw.
func Check(err error) *Result[struct{}] {
	return &Result[struct{}]{error: err}
}

// CheckVal is an advanced version of ItVal that returns Result value. Result
// allows to test/process error before throw.
func CheckVal[T any](val T, err error) *Result[T] {
	if err != nil {
		return &Result[T]{error: err}
	}

	return &Result[T]{value: val}
}

// It is a high-level function that just throws error if passed err is not nil.
// It is useful for functions that returns error as a single value.
func It(err error) { Check(err).Wrap() }

// ItVal is a high-level function that throws error if passed err is not nil. If
// err is nil, it returns a passed val. It is useful for functions that returns
// two values: a result and an error.
func ItVal[T any](val T, err error) T {
	return CheckVal(val, err).Wrap()
}

// Throw is an alias for It. It throws err if err is not nil. The 'throw' verb
// is more useful if you know exactly that err is not nil.
func Throw(err error) { It(err) }
