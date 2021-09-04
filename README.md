# pflagext

[![Go Report Card](https://goreportcard.com/badge/github.com/sethp/pflagext)](https://goreportcard.com/report/github.com/sethp/pflagext) [![Coverage Status](https://coveralls.io/repos/github/sethp/pflagext/badge.svg?branch=main)](https://coveralls.io/github/sethp/pflagext?branch=main) [![Go Reference](https://pkg.go.dev/badge/github.com/sethp/pflagext.svg)](https://pkg.go.dev/github.com/sethp/pflagext)

Extensions to https://github.com/spf13/pflag

## Installation

Use the standard go get command:

```
go get github.com/sethp/pflagext
```

## Usage

Either mix pflagext package-level calls with

```golang
var (
    myVar = myStruct{}
    myStr = "default"
)

func init() {
    pflag.StringVar(&myStr, "strflag", myStr, "help")
    pflagext.Func("structflag", "help", func(s string) (err error) {
        myVar, err = myStructFromString(s)
        return
    })
}
```

Or wrap a `*pflag.FlagSet` and use the bound methods:

```golang
var (
    myVar = myStruct{}
    myStr = "default"
)

func init() {
    fs := pflagext.FlagSetExt{pflag.CommandLine}
    fs.StringVar(&myStr, "strflag", myStr, "help")
    fs.Func("structflag", "help", func(s string) (err error) {
        myVar, err = myStructFromString(s)
        return
    })
}
```
