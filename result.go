package try

import "errors"

type Result[T any] struct {
	value T
	error error
}

func (r *Result[T]) Val() T     { return r.value }
func (r *Result[T]) Err() error { return r.error }

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

func (r *Result[T]) Allow(errs ...error) (T, error) {
	r1 := r.AllowOr(errs...)
	return r1.value, r.error
}

func (r *Result[T]) Annotate(format string, args ...any) T {
	return r.Wrap(Annotate(format, args...))
}

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

func (r *Result[T]) Wrap(ws ...Wrapper) T {
	r1 := r.WrapOr(ws...)
	if r1.error != nil {
		panic(boxError(r.error))
	}
	return r1.value
}
