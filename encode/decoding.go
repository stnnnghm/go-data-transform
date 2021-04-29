package encode

import (
	"encoding/json"
	"fmt"
)

func DecodeUser(data []byte) (User, error) {
	fmt.Println("In DecodeUser")
	var u User

	err := json.Unmarshal(data, &u)
	if err != nil {
		return u, err
	}

	fmt.Printf("Struct u: %v\n", u)
	return u, nil
}

func DecodeUsers(data []byte) ([]User, error) {
	fmt.Println("In DecodeUsers")
	var u []User

	err := json.Unmarshal(data, &u)
	if err != nil {
		return u, err
	}

	for i := range u {
		fmt.Printf("User %d:\n\t ID: %v\n\t Name: %v\n\t Occupation: %v\n",
			u[i].Id, u[i].Id, u[i].Name, u[i].Occupation)
	}

	return u, nil
}
