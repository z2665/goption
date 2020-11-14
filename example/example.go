package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/z2665/goption"
)

func getSome() goption.Option {
	return goption.Some("hello world")
}
func getNone() goption.Option {
	return goption.None()
}
func getOk() goption.Result {
	return goption.Ok("hello result is ok")
}
func getErr() goption.Result {
	return goption.Err(errors.New("hello result is error"))
}
func MockOpenFile(filename string) (*os.File, error) {
	if filename == "not exists" {
		return nil, errors.New("not exists")

	}
	return &os.File{}, nil
}
func WillreturnNil() *os.File {
	return nil
}
func main() {
	s := getSome()
	if !s.Is_None() {
		fmt.Printf("we get a value %s\n", s.Get().(string))
	}
	//pattern matching style ?
	s.Some(func(v interface{}) {
		fmt.Printf("we get a value %s\n", v.(string))
	}).None(func() {
		fmt.Println("we get a nothing")
	})
	s = getNone()
	if !s.Is_None() {
		fmt.Printf("we get a value %s\n", s.Get().(string))
	}
	s.Some(func(v interface{}) {
		fmt.Printf("we get a value %s\n", v.(string))
	}).None(func() {
		fmt.Println("we get a nothing")
	})
	//Result example
	o := getOk()
	if o.Is_Ok() {
		fmt.Printf("we get a value %s\n", o.Unwrap().(string))
	}
	o.Ok(func(v interface{}) {
		fmt.Printf("we get a value %s\n", v.(string))
	}).Err(func(e error) {
		fmt.Println(e.Error())
	})
	o = getErr()
	o.Ok(func(v interface{}) {
		fmt.Printf("we get a value %s\n", v.(string))
	}).Err(func(e error) {
		fmt.Println(e.Error())
	})
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	// o.Unwrap()
	file2 := goption.ToOption(WillreturnNil())
	if file2.Is_None() {
		fmt.Println("wget none")
	}
	file1 := goption.ToResult(MockOpenFile("a file")).Unwrap().(*os.File)
	fmt.Printf("%v", file1)
	goption.ToResult(MockOpenFile("not exists")).Unwrap()

}
