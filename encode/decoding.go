package encode

import (
	"encoding/json"
	"fmt"
)

// DecodeUser accepts []bytes and returns a User if possible
func DecodeUser(data []byte) (User, error) {
	var u User

	err := json.Unmarshal(data, &u)
	if err != nil {
		return u, fmt.Errorf("failure unmarshal: %v", err)
	}

	fmt.Printf("Struct u: %v\n", u)
	return u, nil
}

// DecodeUsers accepts []bytes and returns []User if possible
func DecodeUsers(data []byte) ([]User, error) {
	var u []User

	err := json.Unmarshal(data, &u)
	if err != nil {
		return u, fmt.Errorf("failure unmarshal in DecodeUsers: %v", err)
	}

	for i := range u {
		fmt.Printf("User %d:\n\t ID: %v\n\t Name: %v\n\t Occupation: %v\n",
			u[i].Id, u[i].Id, u[i].Name, u[i].Occupation)
	}

	return u, nil
}
