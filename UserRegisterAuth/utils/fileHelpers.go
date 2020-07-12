package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {

	if IsEmpty(fileName) {
		return "", errors.New("The File Name Should Not Empty")
	}

	bytes, _ := ioutil.ReadFile(fileName)
	// check the error
	//CheckError(err)

	return string(bytes), nil
}