package pflagext

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

func TestFunc(t *testing.T) {
	f := FlagSetExt{pflag.NewFlagSet("test", pflag.ContinueOnError)}
	var ss []string

	f.Func("fn", "func", func(s string) error {
		ss = append(ss, s)
		return nil
	})

	err := f.Parse([]string{"--fn", "1", "--fn=2,3"})
	if err != nil {
		t.Fatal("f.Parse() = ", err)
	}
	if want := []string{"1", "2,3"}; !reflect.DeepEqual(ss, want) {
		t.Errorf("f.Parse() yielded %#v, wanted %#v", ss, want)
	}

	f = FlagSetExt{pflag.NewFlagSet("test", pflag.ContinueOnError)}
	f.Func("fn", "func", func(s string) error {
		return errors.New("test error")
	})

	err = f.Parse([]string{})
	if err != nil {
		t.Error("f.Parse() = ", err)
	}

	err = f.Parse([]string{"--fn="})
	if err == nil {
		t.Fatal("f.Parse() = <nil>, wanted error")
	}
	if !strings.Contains(err.Error(), "test error") {
		t.Errorf("error should contain %q, saw %v", "test error", err)
	}
}
