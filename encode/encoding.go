package encode

import (
	"encoding/json"
	"fmt"
)

func EncodeUser(u User) ([]byte, error) {
	fmt.Println("In EncodeUser")

	d, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Encoded: %v\nOutput: %v\n", u, string(d))
	return d, nil
}

func EncodeUsers(users []User) ([]byte, error) {
	fmt.Println("In EncodeUsers")

	d, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return d, nil
}
