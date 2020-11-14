package goption

import "reflect"

// Option is  interface to be the type of function return
type Option interface {
	Is_None() bool
	Some(func(interface{})) Option
	None(func()) Option
	Get() interface{}
}

type vSome struct {
	v interface{}
}

func (v *vSome) Is_None() bool {
	return false
}
func (v *vSome) Some(f func(interface{})) Option {
	f(v.v)
	return v
}
func (v *vSome) None(func()) Option {
	return v
}
func (v *vSome) Get() interface{} {
	return v.v
}

//Some return a type vSome
func Some(v interface{}) Option {
	return &vSome{v: v}
}

type vNone struct{}

func (v *vNone) Is_None() bool {
	return true
}
func (v *vNone) Some(f func(interface{})) Option {
	return v
}
func (v *vNone) None(f func()) Option {
	f()
	return v
}
func (v *vNone) Get() interface{} {
	panic("none not support get you should use is_none to check")
}

//None return a type vNone
func None() Option {
	return &vNone{}
}

//ToOption conver golang nil to option
// if function return nil then this function return None
func ToOption(v interface{}) Option {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		return None()
	}
	return Some(v)
}
