package readenv_test

import (
	"fmt"
	"os"

	"github.com/comame/readenv-go"
)

func ExampleRead() {
	os.Setenv("FOO", "BAR")
	os.Setenv("BAZ", "1")

	type Env struct {
		Foo string `env:"FOO"`
		Baz string `env:"BAZ"`
	}

	var env Env
	readenv.Read(&env)

	fmt.Println(env.Foo)
	fmt.Println(env.Baz)

	// Output:
	// BAR
	// 1
}
