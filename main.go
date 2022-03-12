package main

import (
	"encoding/json"
	"example/myapp/pkg/common"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// Let's do something more inline with the Golang standards
	person := common.Person{
		FirstName: "Benjamin",
		LastName:  "Franklin",
		Address: common.Address{
			Street1: "1600 Pennyslvania Avenue",
			Street2: "PO Box: The WhiteHouse",
			City:    "Washington",
			State:   "DC",
			Zip5:    11111,
		},
	}

	b, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
