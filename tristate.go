//
// Copyright 2018 Timothy E. Peoples
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//

// Package tristate provides a custom TriState flag for use with the alternate
// flag package github.com/spf13/pflag. A TriState value may take one of three
// forms: True, False or None and is most useful when you are, for instance,
// filtering records based on a current boolean value -- and you need all
// three possibilities (e.g. True, False and I Don't Care)
//
// There are 8 separate functions provided for defining a TriState flag with
// some combination of the suffixes: "Var", "P", and "FS" each having the
// following meaning:
//
// 		Var: Accepts a TriState pointer instead of returning one.
// 		P:   Also takes a shorthand character for use with a single dash
// 		FS:  Accepts a *FlagSet where the flag should be added
//
// The "Var" and "P" suffixes follow the common pflag convention. The "FS"
// suffix is added to allow the use of alternate FlagSets.
//
package tristate // import "toolman.org/flags/tristate"

import (
	"errors"
	"strings"

	"github.com/spf13/pflag"
)

var (
	// CommandLine is the default FlagSet where flags will be added (unless
	// otherwise specified)
	CommandLine = pflag.CommandLine
)

// Flag defines a tristate.TriState flag with the specified name, default
// value and usage string. The return value is the address of a TriState variable
// the stores the values of the flag.
func Flag(name string, value TriState, usage string) *TriState {
	return FlagFS(CommandLine, name, value, usage)
}

// FlagP is similar to Flag butl also accepts a shorthand letter to be
// used after a single dash.
func FlagP(name, shorthand string, value TriState, usage string) *TriState {
	return FlagPFS(CommandLine, name, shorthand, value, usage)
}

// FlagVar is similar to Flag that also accepts a pointer to TriState
// variable where the flag value should be stored.
func FlagVar(ts *TriState, name string, value TriState, usage string) {
	FlagVarFS(CommandLine, ts, name, value, usage)
}

// FlagVarP is the combination of FlagVar and FlagP.
func FlagVarP(ts *TriState, name, shorthand string, value TriState, usage string) {
	FlagVarPFS(CommandLine, ts, name, shorthand, value, usage)
}

// FlagFS is similar to Flag but accepts a pointer to the FlagSet where
// this flag should be added.
func FlagFS(fs *pflag.FlagSet, name string, value TriState, usage string) *TriState {
	return FlagPFS(fs, name, "", value, usage)
}

// FlagPFS is similar to FlagP but accepts a pointer to the FlagSet
// where this flag should be added.
func FlagPFS(fs *pflag.FlagSet, name, shorthand string, value TriState, usage string) *TriState {
	ts := new(TriState)
	FlagVarPFS(fs, ts, name, shorthand, value, usage)
	return ts
}

// FlagVarFS is similar to FlagVar but accepts a pointer to the FlagSet
// where this flag should be added.
func FlagVarFS(fs *pflag.FlagSet, ts *TriState, name string, value TriState, usage string) {
	FlagVarPFS(fs, ts, name, "", value, usage)
}

// FlagVarPFS is similar to FlagVarP but accepts a pointer to the FlagSet where this flag should be added.
func FlagVarPFS(fs *pflag.FlagSet, ts *TriState, name, shorthand string, value TriState, usage string) {
	fs.VarP(newTriState(value, ts), name, shorthand, usage)
}

// TriState is a TriState Value and may have one of three values: None, False or
// True. Its "zero" value is None.
type TriState int

// The three possible tristate values
const (
	None TriState = iota
	False
	True
)

func newTriState(val TriState, p *TriState) *TriState {
	*p = val
	return (*TriState)(p)
}

// Get implements flag.Getter
func (ts *TriState) Get() interface{} {
	return TriState(*ts)
}

// String contributes to the implementation of pflag.Value
func (ts TriState) String() string {
	var s string
	switch ts {
	case None:
		s = "None"
	case False:
		s = "False"
	case True:
		s = "True"
	}
	return s
}

// Type contributes to the implementation of pflag.Value
func (ts *TriState) Type() string {
	return "TriState"
}

// ErrBadTriStateValue is returned by Set if if cannot parse its input.
var ErrBadTriStateValue = errors.New("bad tristate value")

// Set the tristate.Value by parsing s according to the following rules:
//
//     Value: Strings
//     ------ ----------------------------------------------------------
//     True:  1, t, true, y, yes
//     False: 0, f, false, n, no
//     None:  -1, u, unknown, e, either, b, both, a, all, none, null, nil
//
// Allstring input is case insensitive. Any string not mentioned above will
// return ErrBadTriStateValue.
//
// Set contributes to the implementation of pflag.Value
func (ts *TriState) Set(s string) error {
	switch strings.ToLower(s) {
	case "1", "t", "y", "true", "yes":
		*ts = True
	case "0", "f", "n", "false", "no":
		*ts = False
	case "-1", "u", "e", "b", "a", "unknown", "either", "both", "all", "any", "none", "null", "nil":
		*ts = None
	default:
		return ErrBadTriStateValue
	}
	return nil
}

// Bool returns a pointer to a bool holding the value of ts. If ts is None then
// Bool returns nil.
func (ts *TriState) Bool() *bool {
	var b bool
	switch *ts {
	case None:
		return nil
	case False:
		b = false
	case True:
		b = true
	}
	return &b
}

// Match returns true if ts is not None and matches the boolean value of b.
// If ts is None, b is ignored and noneVal is returned.
func (ts *TriState) Match(b, noneVal bool) bool {
	if bp := ts.Bool(); bp != nil {
		return *bp == b
	}
	return noneVal
}

// IsSet returns false if ts is None, otherwise true.
func (ts *TriState) IsSet() bool {
	return *ts != None
}
