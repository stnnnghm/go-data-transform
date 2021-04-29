package main

import (
	"fmt"
	"log"

	e "github.com/stnnnghm/go-data-transform/encode"
	f "github.com/stnnnghm/go-data-transform/readfile"
)

func main() {
	fmt.Println("Starting Go-Data-Transform")

	d := []byte(`{
		"Id": 1,
		"Name": "Stevan",
		"Occupation" : "UE"
	}`)

	user, err := e.DecodeUser(d)
	if err != nil {
		log.Fatalf("%v", err)
	}

	d2 := []byte(`
	[
		{"Id": 2, "Name": "Kim", "Occupation": "Other"},
		{"Id": 3, "Name": "Travis", "Occupation": "Tech"},
		{"Id": 4, "Name": "Derek", "Occupation": "Casino"}
	]`)

	users, err := e.DecodeUsers(d2)
	if err != nil {
		log.Fatalf("%v", err)
	}

	userByte, err := e.EncodeUser(user)
	if err != nil {
		log.Fatalf("%v", err)
	}

	userBytes, err := e.EncodeUsers(users)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fUser, err := f.ReadFromFile("sample_data.json")
	if err != nil {
		log.Fatalf("%v", err)
	}

	fileUser, err := e.DecodeUsers(fUser)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Main User: ", user)
	fmt.Printf("Main Users: %v\n", users)
	fmt.Println("Main userByte: ", string(userByte))
	fmt.Printf("Main userBytes: %v\n", string(userBytes))
	fmt.Println("Main fileUser: ", fileUser)

	fmt.Println("Exiting...")
}
