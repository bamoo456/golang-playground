package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

	fmt.Println("=====demo on json.encode=====")

	// need io.Writer as input to get json encoder
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(george); err != nil {
		fmt.Println("Failed to encode goerge strcut", err)
	}

	fmt.Println("=====demo on json.decode=====")
	emptyPerson := person{}
	jsonString := `{"FirstName": "David", "LastName": "Chen", "last_name": "Lee", "age": 20}`
	// get strings reader
	reader := strings.NewReader(jsonString)
	// need io.Reader as input to get json decoder
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(&emptyPerson); err != nil {
		fmt.Println("Failed to decode json string", err)
	} else {
		fmt.Println("Decode result on emptyPerson", emptyPerson)
	}
}
