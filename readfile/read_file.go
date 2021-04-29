package readfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFromFile checks to see if the supplied file name is available and
// returns the response as []byte if successful
func ReadFromFile(fileName string) ([]byte, error) {
	_, err := fileExists(fileName)
	if err != nil {
		return nil, fmt.Errorf("supplied filename not found: %v", err)

	} else {
		filename, _ := os.Open(fileName)
		defer filename.Close()

		data, err := ioutil.ReadAll(filename)
		if err != nil {
			return nil, fmt.Errorf("error reading from file: %v", err)
		}

		fmt.Printf("Successful read from file %s\n", fileName)
		return data, nil
	}
}

func fileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	} else if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}
