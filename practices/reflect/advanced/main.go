package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	ptrV := reflect.ValueOf(&x)

	ptrVElem := ptrV.Elem() // ptrVElem refers to the variable x

	// check the address ability of x's address
	fmt.Printf("Pointer address [%v], value [%v]\n", ptrV, ptrVElem)
	fmt.Printf("The ptrVElem type = %T\n", ptrVElem)
	fmt.Println("Can pointer addressable ", ptrV.CanAddr())

	// use x's address to change its value
	px := ptrVElem.Addr().Interface().(*int) // px := &x
	*px = 3                                  // x = 3
	fmt.Println(x)                           // "3"

	// the other way to change x's value by reflect.Set
	// the ptrVElem is addressable, so we can use below way to configure x's value
	// ptrVElem is like "&x" now
	ptrVElem.Set(reflect.ValueOf(10))
	fmt.Println(x)

	// below will cause panic, because the "xReflectValue" is not addressable
	xReflectValue := reflect.ValueOf(x)
	// xReflectValue.Set(reflect.ValueOf(20))
	fmt.Printf("xReflectValue is able to address=%v, is able to set=%v \n",
		xReflectValue.CanAddr(),
		xReflectValue.CanSet())

	ptrVElem.SetInt(30)
	fmt.Println(x)

	//	reflect interface{} and assign it by reflect.Set()
	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// panic: SetInt called on interface Value
	// ry.SetInt(2)
	ry.Set(reflect.ValueOf(3)) // OK, y = int(3)
	fmt.Println("rv is", ry)

	// panic: SetString called on interface Value
	// ry.SetString("hello")
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
	fmt.Println("rv is", ry)

}
