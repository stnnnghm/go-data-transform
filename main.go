package main

import (
	"fmt"
	"log"

	e "github.com/stnnnghm/go-data-transform/encode"
)

func main() {
	fmt.Println("Starting Go-Data-Transform")

	d := []byte(`{
		"Id": 1,
		"Name": "Stevan",
		"Occupation" : "UE"
	}`)

	u, err := e.EncodeUser(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Main User: ", u)

	fmt.Println("Exiting...")
}
