package goption

// Result is  interface to be the type of function return
type Result interface {
	Ok(func(interface{})) Result
	Err(func(error)) Result
	Is_Ok() bool
	Unwrap() interface{}
}

type vOk struct {
	v interface{}
}

//Ok returns the type vOk
func Ok(v interface{}) Result {
	return &vOk{v: v}
}
func (v *vOk) Ok(f func(interface{})) Result {
	f(v.v)
	return v
}
func (v *vOk) Err(f func(error)) Result {
	return v
}

func (v *vOk) Is_Ok() bool {
	return true
}
func (v *vOk) Unwrap() interface{} {
	return v.v
}

type vErr struct {
	v error
}

//Err returns the type vErr
func Err(v error) Result {
	return &vErr{v: v}
}
func (v *vErr) Ok(f func(interface{})) Result {
	return v
}
func (v *vErr) Err(f func(error)) Result {
	f(v.v)
	return v
}
func (v *vErr) Is_Ok() bool {
	return false
}
func (v *vErr) Unwrap() interface{} {
	panic(v.v)
}

//ToResult conver golang function return to result
func ToResult(i interface{}, err error) Result {
	if err != nil {
		return Err(err)
	}
	return Ok(i)
}
