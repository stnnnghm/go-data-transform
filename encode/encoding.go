package encode

import (
	"encoding/json"
	"fmt"
)

// EncodeUser accepts a User and returns the response as []byte
func EncodeUser(u User) ([]byte, error) {
	d, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("error marshal %v: %v", u, err)
	}

	fmt.Printf("Encoded: %v\nOutput: %v\n", u, string(d))
	return d, nil
}

// EncodeUsers accepts []User and returns the response as []byte
func EncodeUsers(users []User) ([]byte, error) {
	d, err := json.Marshal(users)
	if err != nil {
		return nil, fmt.Errorf("error marshal %v: %v", users, err)
	}

	return d, nil
}
