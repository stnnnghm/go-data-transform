package readfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFromFile(fileName string) ([]byte, error) {
	fmt.Println("In ReadFromFile")

	filename, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer filename.Close()

	data, err := ioutil.ReadAll(filename)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successful read from file %s\n", fileName)

	return data, nil
}
