package try_test

import (
	"encoding/json"
	"fmt"

	"github.com/davidmz/try"
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
