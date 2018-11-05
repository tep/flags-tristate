
[![GoDoc](https://godoc.org/toolman.org/flags/tristate?status.svg)](https://godoc.org/toolman.org/flags/tristate)  [![Go Report Card](https://goreportcard.com/badge/toolman.org/flags/tristate)](https://goreportcard.com/report/toolman.org/flags/tristate) [![Build Status](https://travis-ci.org/tep/flags-tristate.svg?branch=master)](https://travis-ci.org/tep/flags-tristate)


# tristate
`import "toolman.org/flags/tristate"`

* [Install](#pkg-install)
* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-install">Install</a>

```sh
    go get toolman.org/flags/tristate
```

## <a name="pkg-overview">Overview</a>
Package tristate provides a custom TriState flag for use with the alternate
flag package github.com/spf13/pflag. A TriState value may take one of three
forms: True, False or None and is most useful when you are, for instance,
filtering records based on a current boolean value -- and you need all
three possibilities (e.g. True, False and I Don't Care)

There are 8 separate functions provided for defining a TriState flag with
some combination of the suffixes: "Var", "P", and "FS" each having the
following meaning:


	Var: Accepts a TriState pointer instead of returning one.
	P:   Also takes a shorthand character for use with a single dash
	FS:  Accepts a *FlagSet where the flag should be added

The "Var" and "P" suffixes follow the common pflag convention. The "FS"
suffix is added to allow the use of alternate FlagSets.




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func FlagVar(ts *TriState, name string, value TriState, usage string)](#FlagVar)
* [func FlagVarFS(fs *pflag.FlagSet, ts *TriState, name string, value TriState, usage string)](#FlagVarFS)
* [func FlagVarP(ts *TriState, name, shorthand string, value TriState, usage string)](#FlagVarP)
* [func FlagVarPFS(fs *pflag.FlagSet, ts *TriState, name, shorthand string, value TriState, usage string)](#FlagVarPFS)
* [type TriState](#TriState)
  * [func Flag(name string, value TriState, usage string) *TriState](#Flag)
  * [func FlagFS(fs *pflag.FlagSet, name string, value TriState, usage string) *TriState](#FlagFS)
  * [func FlagP(name, shorthand string, value TriState, usage string) *TriState](#FlagP)
  * [func FlagPFS(fs *pflag.FlagSet, name, shorthand string, value TriState, usage string) *TriState](#FlagPFS)
  * [func (ts *TriState) Bool() *bool](#TriState.Bool)
  * [func (ts *TriState) Get() interface{}](#TriState.Get)
  * [func (ts *TriState) Set(s string) error](#TriState.Set)
  * [func (ts TriState) String() string](#TriState.String)
  * [func (ts *TriState) Type() string](#TriState.Type)


#### <a name="pkg-files">Package files</a>
[tristate.go](/src/toolman.org/flags/tristate/tristate.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    // CommandLine is the default FlagSet where flags will be added (unless
    // otherwise specified)
    CommandLine = pflag.CommandLine
)
```
``` go
var ErrBadTriStateValue = errors.New("bad tristate value")
```
ErrBadTriStateValue is returned by Set if if cannot parse its input.



## <a name="FlagVar">func</a> [FlagVar](/src/target/tristate.go?s=2927:2996#L60)
``` go
func FlagVar(ts *TriState, name string, value TriState, usage string)
```
FlagVar is similar to Flag that also accepts a pointer to TriState
variable where the flag value should be stored.



## <a name="FlagVarFS">func</a> [FlagVarFS](/src/target/tristate.go?s=3881:3971#L85)
``` go
func FlagVarFS(fs *pflag.FlagSet, ts *TriState, name string, value TriState, usage string)
```
FlagVarFS is similar to FlagVar but accepts a pointer to the FlagSet
where this flag should be added.



## <a name="FlagVarP">func</a> [FlagVarP](/src/target/tristate.go?s=3103:3184#L65)
``` go
func FlagVarP(ts *TriState, name, shorthand string, value TriState, usage string)
```
FlagVarP is the combination of FlagVar and FlagP.



## <a name="FlagVarPFS">func</a> [FlagVarPFS](/src/target/tristate.go?s=4128:4230#L90)
``` go
func FlagVarPFS(fs *pflag.FlagSet, ts *TriState, name, shorthand string, value TriState, usage string)
```
FlagVarPFS is similar to FlagVarP but accepts a pointer to the FlagSet where this flag should be added.




## <a name="TriState">type</a> [TriState](/src/target/tristate.go?s=4409:4426#L96)
``` go
type TriState int
```
TriState is a TriState Value and may have one of three values: None, False or
True. Its "zero" value is None.


``` go
const (
    None TriState = iota
    False
    True
)
```
The three possible tristate values







### <a name="Flag">func</a> [Flag](/src/target/tristate.go?s=2450:2512#L48)
``` go
func Flag(name string, value TriState, usage string) *TriState
```
Flag defines a tristate.TriState flag with the specified name, default
value and usage string. The return value is the address of a TriState variable
the stores the values of the flag.


### <a name="FlagFS">func</a> [FlagFS](/src/target/tristate.go?s=3352:3435#L71)
``` go
func FlagFS(fs *pflag.FlagSet, name string, value TriState, usage string) *TriState
```
FlagFS is similar to Flag but accepts a pointer to the FlagSet where
this flag should be added.


### <a name="FlagP">func</a> [FlagP](/src/target/tristate.go?s=2666:2740#L54)
``` go
func FlagP(name, shorthand string, value TriState, usage string) *TriState
```
FlagP is similar to Flag butl also accepts a shorthand letter to be
used after a single dash.


### <a name="FlagPFS">func</a> [FlagPFS](/src/target/tristate.go?s=3589:3684#L77)
``` go
func FlagPFS(fs *pflag.FlagSet, name, shorthand string, value TriState, usage string) *TriState
```
FlagPFS is similar to FlagP but accepts a pointer to the FlagSet
where this flag should be added.





### <a name="TriState.Bool">func</a> (\*TriState) [Bool](/src/target/tristate.go?s=6081:6113#L165)
``` go
func (ts *TriState) Bool() *bool
```
Bool returns a pointer to a bool holding the value of ts. If ts is None then
Bool returns nil.




### <a name="TriState.Get">func</a> (\*TriState) [Get](/src/target/tristate.go?s=4634:4671#L111)
``` go
func (ts *TriState) Get() interface{}
```
Get implements flag.Getter




### <a name="TriState.Set">func</a> (\*TriState) [Set](/src/target/tristate.go?s=5642:5681#L149)
``` go
func (ts *TriState) Set(s string) error
```
Set the tristate.Value by parsing s according to the following rules:


	Value: Strings
	------ ----------------------------------------------------------
	True:  1, t, true, y, yes
	False: 0, f, false, n, no
	None:  -1, u, unknown, e, either, b, both, a, all, none, null, nil

Allstring input is case insensitive. Any string not mentioned above will
return ErrBadTriStateValue.

Set contributes to the implementation of pflag.Value




### <a name="TriState.String">func</a> (TriState) [String](/src/target/tristate.go?s=4758:4792#L116)
``` go
func (ts TriState) String() string
```
String contributes to the implementation of pflag.Value




### <a name="TriState.Type">func</a> (\*TriState) [Type](/src/target/tristate.go?s=4972:5005#L130)
``` go
func (ts *TriState) Type() string
```
Type contributes to the implementation of pflag.Value
