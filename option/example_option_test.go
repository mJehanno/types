package option_test

import (
	"fmt"

	"github.com/mJehanno/types/option"
)

type Config struct {
	RootPath option.Option[string]
}

func GetConfigFromEnv() Config {
	return Config{RootPath: option.NewNone[string]()}
}

func ExampleOption() {
	config := GetConfigFromEnv()

	fmt.Println(config.RootPath.UnwrapOrDefault("~/.config/example"))
	// Output: "~/.config/example"
}
