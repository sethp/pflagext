// Package pflagext wraps github.com/spf13/pflag and provides extensions.
package pflagext

import (
	"github.com/spf13/pflag"
)

// FlagSetExt extends pflag.FlagSet with additional command-line options.
//
// Use by either creating a new empty FlagSetExt with `NewFlagSet` or wrapping
// an existing `fs` with `FlagSetExt{fs}`.
type FlagSetExt struct {
	*pflag.FlagSet
}

// NewFlagSet creates a new empty flag set.
func NewFlagSet(name string, errorHandling pflag.ErrorHandling) *FlagSetExt {
	return &FlagSetExt{
		pflag.NewFlagSet(name, errorHandling),
	}
}
