package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 3
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Printf("The type of integer [i] : %s\n", t)
	fmt.Printf("The value of integer [i] : %v\n", v)

	t1 := v.Type() // t1 is the same as t
	fmt.Println(t1)

	iface := v.Interface() // iface now is an "interface{}" type

	// because the iface is an "interface{}", so we can
	// do the type asertion
	v1 := iface.(int)

	fmt.Printf("v1 type: %T, value: %v\n", v1, v1)

	// do the normal integer reflection
	v1Type := getReflectType(v)
	fmt.Println(v1Type)

	// do the map reflection
	mType := map[string]int{}
	mType["1st"] = 1
	mType["2nd"] = 2
	mReflectValue := reflect.ValueOf(mType)
	mReflectValueType := getReflectType(mReflectValue)
	fmt.Println(mReflectValueType)

	// "reflect.Value" type support many methods, however,
	// if the method does not match the current type, then it return panic

	// e.g. use MapKeys() on "slice" type
	// arr := []int{1, 2, 3}
	// reflect.ValueOf(arr).MapKeys() // this will got

}

func getReflectType(v reflect.Value) string {
	// use v.Kind() to do the switch
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int32, reflect.Int64:
		return "int number"
	case reflect.Map:
		fmt.Println("[getReflectType] map keys", v.MapKeys())
		return "got the map"
	default:
		return "not implemented"
	}
}
