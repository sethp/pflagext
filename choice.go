package pflagext

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
)

type choiceValue struct {
	val *string

	options []string
}

var errChoice = errors.New("invalid selection")

func newChoiceValue(val string, p *string, options ...string) choiceValue {
	*p = val

	found := false
	for _, opt := range options {
		if opt == val {
			found = true
			break
		}
	}

	if !found {
		opts := make([]string, 0, len(options)+1)
		opts = append(opts, val)
		opts = append(opts, options...)
		options = opts
	}

	return choiceValue{
		val:     p,
		options: options,
	}
}

func (c choiceValue) Set(s string) error {
	for _, opt := range c.options {
		if s == opt {
			*c.val = s
			return nil
		}
	}
	return fmt.Errorf("%w: %q is not one of %v", errChoice, s, c.options)
}

func (c choiceValue) Type() string {
	return "choice"
}

func (c choiceValue) String() string {
	return *c.val
}

// ChoiceVar defines a string flag with specified name, default value, and usage string.
// It only permits selection from the available options + default.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSetExt) ChoiceVar(p *string, name, value, usage string, options ...string) {
	f.VarP(newChoiceValue(value, p, options...), name, "", usage)
}

// ChoiceVarP is like ChoiceVar, but accepts a shorthand letter that can be used after a single dash.
// It only permits selection from the available options + default.
func (f *FlagSetExt) ChoiceVarP(p *string, name, shorthand, value, usage string, options ...string) {
	f.VarP(newChoiceValue(value, p, options...), name, shorthand, usage)
}

// ChoiceVar defines a string flag with specified name, default value, and usage string.
// It only permits selection from the available options + default.
// The argument p points to a string variable in which to store the value of the flag.
func ChoiceVar(p *string, name, value, usage string, options ...string) {
	pflag.CommandLine.VarP(newChoiceValue(value, p, options...), name, "", usage)
}

// ChoiceVarP is like ChoiceVar, but accepts a shorthand letter that can be used after a single dash.
// It only permits selection from the available options + default.
func ChoiceVarP(p *string, name, shorthand, value, usage string, options ...string) {
	pflag.CommandLine.VarP(newChoiceValue(value, p, options...), name, shorthand, usage)
}

// Choice defines a string flag with specified name, default value, and usage string.
// It only permits selection from the available options + default.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSetExt) Choice(name, value, usage string, options ...string) *string {
	p := new(string)
	f.ChoiceVarP(p, name, "", value, usage, options...)
	return p
}

// ChoiceP is like Choice, but accepts a shorthand letter that can be used after a single dash.
// It only permits selection from the available options + default.
func (f *FlagSetExt) ChoiceP(name, shorthand, value, usage string, options ...string) *string {
	p := new(string)
	f.ChoiceVarP(p, name, shorthand, value, usage, options...)
	return p
}

// Choice defines a Choice flag with specified name, default value, and usage string.
// The return value is the address of a Choice variable that stores the value of the flag.
func Choice(name, value, usage string, options ...string) *string {
	return (&FlagSetExt{pflag.CommandLine}).ChoiceP(name, "", value, usage, options...)
}

// ChoiceP is like Choice, but accepts a shorthand letter that can be used after a single dash.
func ChoiceP(name, shorthand, value, usage string, options ...string) *string {
	return (&FlagSetExt{pflag.CommandLine}).ChoiceP(name, shorthand, value, usage, options...)
}
