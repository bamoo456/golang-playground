package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	// conver to last_name when using encoding/json
	LastName string `json:"last_name"`
	// because age is lower case (private), so it will not be marshaled
	age int
}

func main() {
	george := person{"George", "Chen", 20}

	fmt.Println("=====demo on json.Marshal=====")
	b, err := json.Marshal(george)
	if err != nil {
		fmt.Println("Failed to do json.Marshal on george")
	} else {
		fmt.Println("Marshal byte on george: ", b)
		fmt.Println("Marshal string on george: ", string(b))
	}

	fmt.Println("=====demo on json.Unmarshal=====")
	emptyPerson := person{}
	jsonString := `{"FirstName": "David", "LastName": "Chen", "last_name": "Lee", "age": 20}`
	json.Unmarshal([]byte(jsonString), &emptyPerson)
	fmt.Println("Unmarshal json string:", jsonString)
	// because LastName has a json tag "last_name", so it got data "last_name"
	// age will not be Unmarshal in this case (private field)
	fmt.Println("Unmarshal done for person:", emptyPerson)
}
