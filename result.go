package try

import "errors"

// Result is a result of function response checking (see Check/CheckVal
// functions). It contains the return value and error.
type Result[T any] struct {
	value T
	error error
}

// Val returns value contained in Result.
func (r *Result[T]) Val() T { return r.value }

// Err returns error contained in Result.
func (r *Result[T]) Err() error { return r.error }

// Allow checks if the Result's error is (as in errors.Is) any of the given
// errs. If it is, the Result is returned, else the error is thrown.
func (r *Result[T]) AllowOr(errs ...error) *Result[T] {
	if r.error != nil {
		for _, err := range errs {
			if !errors.Is(r.error, err) {
				r.Wrap()
			}
		}
	}
	return r
}

// Allow checks if the Result's error is (as in errors.Is) any of the given
// errs. If it is, the Result value and error is returned, else the error is
// thrown.
func (r *Result[T]) Allow(errs ...error) (T, error) {
	r1 := r.AllowOr(errs...)
	return r1.value, r.error
}

// Annotate wraps the Result's error with the custom message using fmt.Errorf
// and throws it.
//
// The last placeholder of format string must be %w. The error passed to wrapper
// is appended to the end of the args list.
//
// r.Annotate(...) is equivalent of r.Wrap(Annotate(...))
func (r *Result[T]) Annotate(format string, args ...any) T {
	return r.Wrap(Annotate(format, args...))
}

// WrapOr wraps Result's error (if it isn't nil) by the given wrappers, and
// returns new Result with the same value and wrapped error.
func (r *Result[T]) WrapOr(ws ...Wrapper) *Result[T] {
	if r.error == nil || len(ws) == 0 {
		return r
	}
	err := r.error
	for _, w := range ws {
		err = w(err)
	}
	return &Result[T]{r.value, err}
}

// Wrap wraps (and then throws) Result's error by the given wrappers. If error
// is nil, it returns Result value instead.
func (r *Result[T]) Wrap(ws ...Wrapper) T {
	r1 := r.WrapOr(ws...)
	if r1.error != nil {
		panic(boxError(r.error))
	}
	return r1.value
}
