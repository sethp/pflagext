package pflagext

import (
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

func TestChoice(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(fs *FlagSetExt) *string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name: "default value",
			setup: func(fs *FlagSetExt) *string {
				return fs.Choice("choice", "default", "usage")
			},
			want: "default",
		},
		{
			name: "default value (explicit)",
			setup: func(fs *FlagSetExt) *string {
				return fs.Choice("choice", "default", "usage")
			},
			args: []string{"--choice", "default"},
			want: "default",
		},
		{
			name: "option",
			setup: func(fs *FlagSetExt) *string {
				return fs.ChoiceP("choice", "c", "default", "usage", "opt")
			},
			args: []string{"-c", "opt"},
			want: "opt",
		},
		{
			name: "invalid option",
			setup: func(fs *FlagSetExt) *string {
				return fs.ChoiceP("choice", "c", "default", "usage", "opt")
			},
			args:    []string{"-c", "not-an-opt"},
			want:    "default", // We'll still set p to the default, even if we get a parse error
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(tt.name, pflag.ContinueOnError)
			got := tt.setup(fs)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("fs.Parse(%#v) = %v, wantErr %v", tt.args, err, tt.wantErr)
			}
			if *got != tt.want {
				t.Errorf("got = %q, want %q", *got, tt.want)
			}
		})
	}
}

func TestChoice_duplicatedDefault(t *testing.T) {
	var out string
	fs := NewFlagSet("test", pflag.ContinueOnError)
	fs.ChoiceVar(&out, "choice", "opt1", "usage", "opt1")

	err := fs.Parse([]string{"--choice", "something else"})
	if count := strings.Count(err.Error(), "opt1"); count != 1 {
		t.Errorf("wanted error to contain `opt1` exactly once, got %d: %q", count, err.Error())
	}
}
