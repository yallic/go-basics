package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Structs like class in oop programing

type User struct {

	// We can give a tag with json declaration like below
	// We can reach values with tags
	Name     string `json:"name"` // When we define tag, the tag is shown in json
	Surname  string // If we don't define any tag json takes default variable name as a json key
	Follower []User `json:"follower,omitempty"` // With this definiton we can ignore this variable if its null nil:)

	Begeni []Begen
}

type Begen struct {
	Date time.Time
}

func main() {

	user1 := User{
		Name:    "Go",
		Surname: "Turkiye",
		Follower: []User{
			{
				Name:    "Follower1",
				Surname: "1",
			},
		},
	}
	//fmt.Println(user1)

	// It returns a json from struct variable ehat we define in the struct definiton
	// arr is byte array
	arr, _ := json.Marshal(user1)
	fmt.Println(string(arr))
	fmt.Println(user1)
}
