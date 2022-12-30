package try

func Check(err error) *Result[struct{}] {
	return &Result[struct{}]{error: err}
}

func CheckVal[T any](val T, err error) *Result[T] {
	if err != nil {
		return &Result[T]{error: err}
	}

	return &Result[T]{value: val}
}

func It(err error) { Check(err).Wrap() }

func ItVal[T any](val T, err error) T {
	return CheckVal(val, err).Wrap()
}

func Throw(err error) { It(err) }
