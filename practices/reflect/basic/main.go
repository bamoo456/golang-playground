package main

import (
	"fmt"
	"reflect"
)

type book struct {
	name   string
	author string
}

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
	arr := []int{1, 2, 3}
	arrReflectValue := reflect.ValueOf(arr)
	// arrReflectValue.MapKeys() // this will got

	// walk through the map
	walk(mReflectValue)

	// walk through the slice
	walk(arrReflectValue)

	// walk through an customized struct
	b := book{"text book", "George"}
	bReflectValue := reflect.ValueOf(b)
	walk(bReflectValue)
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

// recursively walk through
func walk(v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Println("invalid data")
	case reflect.Slice, reflect.Array:
		fmt.Println("Trying to walk through an [slice]/[array] type")
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i))
		}
	case reflect.Struct:
		fmt.Println("Trying to walk through an [struct] type")
		// reflect also works on unexported fields
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("[field]=%s \n", v.Type().Field(i).Name)
			walk(v.Field(i))
		}
	case reflect.Map:
		fmt.Println("Trying to walk through an [map] type")
		for _, key := range v.MapKeys() {
			fmt.Printf("[key]=%s\n", key)
			walk(v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("data pointer = nil\n")
		} else {
			walk(v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("data interface = nil\n")
		} else {
			walk(v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("data value = %v \n", v)
	}
}
