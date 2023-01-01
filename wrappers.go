package try

import "fmt"

// Wrapper describes function that processes (wraps) error into another form.
type Wrapper func(error) error

// Annotate is a Wrapper that wraps the given error with the custom message
// using fmt.Errorf. The last placeholder of format string must be %w. The error
// passed to wrapper is appended to the end of the args list.
func Annotate(format string, args ...any) Wrapper {
	return func(err error) error {
		args = append(args, err)
		return fmt.Errorf(format, args...)
	}
}
