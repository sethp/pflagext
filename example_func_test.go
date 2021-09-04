package pflagext_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sethp/pflagext"
	"github.com/spf13/pflag"
)

func ExampleFunc() {
	fs := pflagext.NewFlagSet("ExampleFunc", pflag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	var config struct {
		Val string `json:"val"`
	}

	// Slightly contrived, but imagine taking a path and using ioutils.ReadFile
	fs.Func("config", "json config", func(s string) error {
		return json.Unmarshal([]byte(s), &config)
	})
	err := fs.Parse([]string{"--config", `{"val": "my-val"}`})
	fmt.Printf("{ Val: %q }, err %#v\n\n", config.Val, err)

	err = fs.Parse([]string{"--config", `{`})
	fmt.Println(err)
	// Output:
	// { Val: "my-val" }, err <nil>
	//
	// invalid argument "{" for "--config" flag: unexpected end of JSON input
}
