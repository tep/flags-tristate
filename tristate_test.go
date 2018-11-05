// Copyright 2018 Timothy E. Peoples

package tristate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

func TestSet(t *testing.T) {
	tests := map[TriState][]string{
		None:  {"-1", "A", "ALL", "ANY", "All", "Any", "B", "BOTH", "Both", "E", "EITHER", "Either", "NIL", "NONE", "NULL", "Nil", "None", "Null", "U", "UNKNOWN", "Unknown", "a", "all", "any", "b", "both", "e", "either", "nil", "none", "null", "u", "unknown"},
		False: {"0", "F", "FALSE", "False", "N", "NO", "No", "f", "false", "n", "no"},
		True:  {"1", "T", "TRUE", "True", "Y", "YES", "Yes", "t", "true", "y", "yes"},
	}

	for want, list := range tests {
		for _, s := range list {
			ts := new(TriState)
			if err := ts.Set(s); err != nil || *ts != want {
				t.Errorf("ts.Set(%q) == %v (*ts=%s); Wanted %v (*ts=%v)", s, err, ts, nil, want)
			}
		}
	}

	ts := new(TriState)
	if err := ts.Set("doug"); err != ErrBadTriStateValue {
		t.Errorf("ts.Set(%q) == %v; Wanted %v", "doug", err, ErrBadTriStateValue)
	}
}

func TestString(t *testing.T) {
	tests := map[TriState]string{
		None:  "None",
		False: "False",
		True:  "True",
	}

	for ts, want := range tests {
		if got := ts.String(); got != want {
			t.Errorf("%#v.String() == %q; wanted %q", ts, got, want)
		}

		p := &ts

		if got := p.String(); got != want {
			t.Errorf("%#v.String() == %q; wanted %q", p, got, want)
		}
	}
}

type flagOptions struct {
	dflt TriState
	args string
	want TriState
	err  error
}

func TestFlags(t *testing.T) {
	var (
		flagopts  []*flagOptions
		flagPopts []*flagOptions
	)

	for _, dflt := range []TriState{None, False, True} {
		flagopts = append(flagopts, &flagOptions{dflt, "", dflt, nil})
		for _, want := range []TriState{None, False, True} {
			flagopts = append(flagopts, &flagOptions{dflt, fmt.Sprintf("--tristate=%s", want.String()), want, nil})
			flagopts = append(flagopts, &flagOptions{dflt, fmt.Sprintf("--tristate %s", want.String()), want, nil})
			flagPopts = append(flagPopts, &flagOptions{dflt, fmt.Sprintf("-t %s", want.String()), want, nil})
		}
	}

	var tests []*testcase
	for _, what := range []string{"Flag", "FlagP", "FlagVar", "FlagVarP", "FlagFS", "FlagPFS", "FlagVarFS", "FlagVarPFS"} {
		for _, fo := range flagopts {
			tests = append(tests, mkTestcase(what, fo.args, fo.dflt, fo.want, fo.err))
		}
		if strings.HasSuffix(what, "P") {
			for _, fo := range flagPopts {
				tests = append(tests, mkTestcase(what, fo.args, fo.dflt, fo.want, fo.err))
			}
		}
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

type testcase struct {
	name  string
	setup func() (*pflag.FlagSet, *TriState)
	args  string
	want  TriState
	err   error
}

func (tc *testcase) test(t *testing.T) {
	defer func(fs *pflag.FlagSet) { CommandLine = fs }(CommandLine)

	fs, got := tc.setup()
	if err := fs.Parse(strings.Fields(tc.args)); err != tc.err || *got != tc.want {
		t.Errorf("%s: got %v (err=%v); wanted %v (err=%v)", t.Name(), *got, err, tc.want, tc.err)
	}
}

func mkTestcase(what, args string, dflt, want TriState, err error) *testcase {
	fs := pflag.NewFlagSet("test:"+what, pflag.ContinueOnError)

	var setup func() (*pflag.FlagSet, *TriState)
	switch what {
	case "Flag":
		setup = func() (*pflag.FlagSet, *TriState) {
			CommandLine = fs
			ts := Flag("tristate", dflt, "tristate flag")
			return fs, ts
		}

	case "FlagP":
		setup = func() (*pflag.FlagSet, *TriState) {
			CommandLine = fs
			ts := FlagP("tristate", "t", dflt, "tristate flag")
			return fs, ts
		}

	case "FlagVar":
		setup = func() (*pflag.FlagSet, *TriState) {
			CommandLine = fs
			var ts TriState
			FlagVar(&ts, "tristate", dflt, "tristate flag")
			return fs, &ts
		}

	case "FlagVarP":
		setup = func() (*pflag.FlagSet, *TriState) {
			CommandLine = fs
			var ts TriState
			FlagVarP(&ts, "tristate", "t", dflt, "tristate flag")
			return fs, &ts
		}

	case "FlagFS":
		setup = func() (*pflag.FlagSet, *TriState) {
			ts := FlagFS(fs, "tristate", dflt, "tristate flag")
			return fs, ts
		}

	case "FlagPFS":
		setup = func() (*pflag.FlagSet, *TriState) {
			ts := FlagPFS(fs, "tristate", "t", dflt, "tristate flag")
			return fs, ts
		}

	case "FlagVarFS":
		setup = func() (*pflag.FlagSet, *TriState) {
			var ts TriState
			FlagVarFS(fs, &ts, "tristate", dflt, "tristate flag")
			return fs, &ts
		}

	case "FlagVarPFS":
		setup = func() (*pflag.FlagSet, *TriState) {
			var ts TriState
			FlagVarPFS(fs, &ts, "tristate", "t", dflt, "tristate flag")
			return fs, &ts
		}
	}

	return &testcase{fmt.Sprintf("F:%s+D:%s+W:%s+A:%s", what, dflt, want, args), setup, args, want, err}
}
