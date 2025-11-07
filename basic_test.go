package echo_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/echo"
)

func ExampleEcho_basic() {
	// echo "Hello World"
	yup.MustRun(
		Echo("Hello World"),
	)
	// Output:
	// Hello World
}

