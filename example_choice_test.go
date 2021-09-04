package pflagext_test

import (
	"fmt"
	"os"

	"github.com/sethp/pflagext"
	"github.com/spf13/pflag"
)

func ExampleChoice() {
	fs := pflagext.NewFlagSet("ExampleChoice", pflag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	choice := fs.Choice("choice", "default", "usage", "freedom-to")
	err := fs.Parse([]string{"--choice", "freedom-to"})
	fmt.Printf("choice %q, err %#v\n\n", *choice, err)

	err = fs.Parse([]string{"--choice", "freedom-from"})
	fmt.Println(err)
	// Output:
	// choice "freedom-to", err <nil>
	//
	// invalid argument "freedom-from" for "--choice" flag: invalid selection: "freedom-from" is not one of [default freedom-to]
}
