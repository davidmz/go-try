package try_test

import (
	"encoding/json"
	"fmt"

	"github.com/davidmz/go-try"
)

func ExampleIt() {
	defer try.Handle(func(err error) {
		fmt.Println("Oh,", err)
	})

	goodData := []byte(`"Good JSON"`)
	badData := []byte(`Bad JSON`)
	value := ""

	try.It(json.Unmarshal(goodData, &value))
	fmt.Println(value)

	try.It(json.Unmarshal(badData, &value))
	fmt.Println(value)

	// Output:
	// Good JSON
	// Oh, invalid character 'B' looking for beginning of value
}

func ExampleItVal() {
	defer try.Handle(func(err error) {
		fmt.Println("Oh,", err)
	})

	goodValue := "Good value"
	badValue := func() {} // functions cannot be serialized
	var data []byte

	data = try.ItVal(json.Marshal(goodValue))
	fmt.Println(string(data))

	data = try.ItVal(json.Marshal(badValue))
	fmt.Println(string(data))

	// Output:
	// "Good value"
	// Oh, json: unsupported type: func()
}

func ExampleHandleAs() {

	tryUnmarshal := func(data []byte) (result string, outErr error) {
		defer try.HandleAs(&outErr, try.Annotate("tryUnmarshal error: %w"))

		try.It(json.Unmarshal(data, &result))
		return
	}

	_, err := (tryUnmarshal([]byte(`Bad JSON`)))
	fmt.Println(err)

	// Output:
	// tryUnmarshal error: invalid character 'B' looking for beginning of value
}

func ExampleWrap() {
	defer try.Handle(func(err error) { fmt.Println(err) })

	tryUnmarshal := func(data []byte) (result string) {
		defer try.Wrap(try.Annotate("tryUnmarshal error: %w"))

		try.It(json.Unmarshal(data, &result))
		return
	}

	tryUnmarshal([]byte(`Bad JSON`))

	// Output:
	// tryUnmarshal error: invalid character 'B' looking for beginning of value
}
