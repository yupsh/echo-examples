package echo_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/echo"
)

func ExampleEcho_basic() {
	// echo "Hello World"
	gloo.MustRun(
		Echo("Hello World"),
	)
	// Output:
	// Hello World
}
