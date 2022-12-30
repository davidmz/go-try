package try

func Handle(fn func(error), ws ...Wrapper) {
	if pnc := recover(); pnc == nil {
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

func HandleAs(targetError *error, ws ...Wrapper) {
	Handle(func(err error) { *targetError = err })
}
