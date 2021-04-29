package encode

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
}

func EncodeUser(data []byte) (User, error) {
	fmt.Println("In EncodeUser")
	var u1 User

	err := json.Unmarshal(data, &u1)
	if err != nil {
		return u1, err
	}

	fmt.Printf("Struct u1: %v\n", u1)
	fmt.Printf("ID: %v\nName: %v\nOccupation: %v\n",
		u1.Id, u1.Name, u1.Occupation)

	return u1, nil
}
