package try

// errBox is a wrapper around the error value.
type errBox struct{ error }

func boxError(err error) error {
	if b, ok := err.(*errBox); ok {
		return b
	}
	return &errBox{err}
}

// UnboxError is a low-level function that checks the recover-ed value in panic
// handler, if it contains thrown error. If it does, UnboxError returns error
// and true. You don't need this function until you want to manually recover
// panics, thrown by this library.
func UnboxError(val any) (error, bool) {
	if b, ok := val.(*errBox); ok {
		return b.error, ok
	} else {
		return nil, ok
	}
}
