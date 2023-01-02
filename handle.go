package try

// Handle catches thrown errors, wraps them by wrappers (use Annotate here, if
// you need) and passed them to the fn function. Use Handle with defer (see
// examples of It/ItVal).
func Handle(fn func(error), ws ...Wrapper) {
	handle(recover(), fn, ws...)
}

// HandleAs catches thrown errors, wraps them by wrappers (use Annotate here, if
// you need) and assigns them to the targetError. It is useful for functions
// that returns error. See example for use case.
func HandleAs(targetError *error, ws ...Wrapper) {
	handle(recover(), func(err error) { *targetError = err }, ws...)
}

func handle(pnc interface{}, fn func(error), ws ...Wrapper) {
	if pnc == nil {
		// none
	} else if err, ok := UnboxError(pnc); ok {
		for _, w := range ws {
			err = w(err)
		}
		fn(err)
	} else {
		panic(pnc)
	}
}
