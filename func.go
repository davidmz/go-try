package try

// Func takes a function that can throw try.* error and returns a function that
// returns this error in a regular Go style.
func Func(fn func()) func() error {
	return func() (outErr error) {
		defer HandleAs(&outErr)
		fn()
		return
	}
}

// FuncArg takes a function with one argument that can throw try.* error and
// returns a function that returns this error in a regular Go style.
func FuncArg[T any](fn func(x T)) func(x T) error {
	return func(x T) (outErr error) {
		defer HandleAs(&outErr)
		fn(x)
		return
	}
}

// FuncArgOut takes a function with one argument and one return value that can
// throw try.* error and returns a function that returns (value, error) in a
// regular Go style.
func FuncArgOut[T any, U any](fn func(x T) U) func(x T) (U, error) {
	return func(x T) (result U, outErr error) {
		defer HandleAs(&outErr)
		result = fn(x)
		return
	}
}
