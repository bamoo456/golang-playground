package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Director    string
	Title       string
	description string
}

func (m Movie) GetTitle() string {
	return m.Title
}

func (m Movie) GetDirector() string {
	return m.Director
}

func (m *Movie) SetTitle(s string) {
	m.Title = s
}

func main() {
	theMovie := Movie{"Jon Favreau", "Iron Man", ""}
	reflectMovieType := reflect.TypeOf(theMovie)
	reflectMovieValue := reflect.ValueOf(theMovie)

	// use the reflect type to get the number of fields and number of methods
	fmt.Printf("Number of field [%d], Number of method [%d]\n",
		reflectMovieType.NumField(),
		reflectMovieType.NumMethod())

	// use the reflect value to get the number of fields and number of methods
	fmt.Printf("Number of field [%d], Number of method [%d]\n",
		reflectMovieValue.NumField(),
		reflectMovieValue.NumMethod())

	fmt.Println("=====walk through method by value receiver========")
	for m := 0; m < reflectMovieType.NumMethod(); m++ {
		method := reflectMovieType.Method(m)
		//fmt.Println(method)
		//fmt.Println(method.Type)         // func(*main.MyStruct) string
		fmt.Println(method.Name) // GetName
		//fmt.Println(method.Type.NumIn()) // 参数个数
		//fmt.Println(method.Type.In(0))   // 参数类型
	}

	fmt.Println("=====walk through method by pointer receiver========")
	reflectMovieRefType := reflect.TypeOf(&theMovie)
	for m := 0; m < reflectMovieRefType.NumMethod(); m++ {
		method := reflectMovieRefType.Method(m)
		fmt.Println(method.Name) // GetName
		//fmt.Println(method.Type.NumIn()) // 参数个数
		fmt.Println(method.Type.In(0)) // 参数类型
	}

	// use reflect to dynamically call the function
	f := reflect.ValueOf(&theMovie).MethodByName("SetTitle")
	if f.IsValid() {
		v := []reflect.Value{reflect.ValueOf("New Title")}
		f.Call(v)
	} else {
		fmt.Println("====== not a valid function")
	}
	// check the result again
	fmt.Println("theMovie", theMovie)

	// the other way to call the function
	if method, ok := reflectMovieRefType.MethodByName("SetTitle"); !ok {
		fmt.Println("Target method does not exist")
	} else {
		v := []reflect.Value{
			// we must provide this pointer receiver as 1st argument
			reflect.ValueOf(&theMovie),
			reflect.ValueOf("New Movie Title"),
		}
		fmt.Println(method.Func.Call(v))
	}
	fmt.Println("theMovie", theMovie)

}
