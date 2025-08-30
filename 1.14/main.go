package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("int: %d\n", val)
	case string:
		fmt.Printf("string: %s\n", val)
	case bool:
		fmt.Printf("bool: %t\n", val)
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Printf("unknown type: %v\n", reflect.TypeOf(v))
		}
	}
}

func main() {
	var a int = 42
	var b string = "hello"
	var c bool = true
	ch := make(chan int)

	detectType(a)
	detectType(b)
	detectType(c)
	detectType(ch)
	detectType(3.14)
}
