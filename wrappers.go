package try

import "fmt"

type Wrapper func(error) error

func Annotate(format string, args ...any) Wrapper {
	return func(err error) error {
		args = append(args, err)
		return fmt.Errorf(format, args...)
	}
}
