package main 

import (
	"fmt"

	"github.com/z2665/goption"
)

func getSome() goption.Option {
	return goption.Some("hello world")
}
func getNone() goption.Option {
	return goption.None()
}

func main() {
	s:= getSome()
	if !s.Is_None(){
		fmt.Printf("we get a value %s\n",s.Get().(string))
	}
	//pattern matching style ?
	s.Some(func(v interface{}){
		fmt.Printf("we get a value %s\n",v.(string))
	}).None(func(){
		fmt.Println("we get a nothing")
	})
	s= getNone()
	if !s.Is_None(){
		fmt.Printf("we get a value %s\n",s.Get().(string))
	}
	s.Some(func(v interface{}){
		fmt.Printf("we get a value %s\n",v.(string))
	}).None(func(){
		fmt.Println("we get a nothing")
	})
}