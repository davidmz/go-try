package try

// errBox is a wrapper around the error value.
type errBox struct{ error }

func boxError(err error) error {
	if b, ok := err.(*errBox); ok {
		return b
	}
	return &errBox{err}
}

func UnboxError(val any) (error, bool) {
	if b, ok := val.(*errBox); ok {
		return b.error, ok
	} else {
		return nil, ok
	}
}
