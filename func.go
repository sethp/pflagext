package pflagext

import "github.com/spf13/pflag"

type funcValue func(string) error

func (f funcValue) Set(s string) error {
	return f(s)
}

func (f funcValue) Type() string {
	return "func"
}

func (f funcValue) String() string {
	return ""
}

// Func defines a func flag with specified name and usage string.
func (f *FlagSetExt) Func(name, usage string, fn func(string) error) {
	f.FuncP(name, "", usage, fn)
}

// FuncP is like Func, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSetExt) FuncP(name, shorthand, usage string, fn func(string) error) {
	f.VarP(funcValue(fn), name, shorthand, usage)
}

// Func defines a func flag with specified name and usage string.
func Func(name, usage string, fn func(string) error) {
	pflag.CommandLine.Var(funcValue(fn), name, usage)
}

// FuncP is like Func, but accepts a shorthand letter that can be used after a single dash.
func FuncP(name, shorthand, usage string, fn func(string) error) {
	pflag.CommandLine.VarP(funcValue(fn), name, shorthand, usage)
}
