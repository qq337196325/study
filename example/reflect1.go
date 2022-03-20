package main

import (
	"fmt"
	"os"
	"reflect"
)

type Test struct {
	Name string
}

func main() {
	//aa := Test{}
	aa := os.Stdout
	t := reflect.TypeOf(aa) // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)

	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"

	s := reflect.ValueOf(3) // a reflect.Value
	x := s.Interface()      // an interface{}
	i := x.(int)            // an int
	fmt.Printf("%d\n", i)   // "3"
	//fmt.Println(s)
}
