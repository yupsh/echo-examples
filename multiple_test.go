package echo_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/echo"
)

func ExampleEcho_multiple() {
	// echo "one" "two" "three"
	gloo.MustRun(
		Echo("one", "two", "three"),
	)
	// Output:
	// one two three
}
