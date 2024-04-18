package try_test

import (
	"errors"
	"fmt"

	"github.com/davidmz/go-try"
)

func ExampleAnnotate() {
	fmt.Println(try.Annotate("Annotated %w")(errors.New("some error")))

	// Output:
	// Annotated some error
}

func ExampleAnnotate_nil() {
	fmt.Println(try.Annotate("Annotated %w")(nil))

	// Output:
	// <nil>
}
