package echo_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/echo"
)

func ExampleEcho_multiple() {
	// echo "one" "two" "three"
	yup.MustRun(
		Echo("one", "two", "three"),
	)
	// Output:
	// one two three
}

