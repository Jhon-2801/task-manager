package jwt

import (
	"io/ioutil"
)

func LoadFiles(privateFile, PublicFile string) ([]byte, []byte, error) {
	var (
		privateBytes []byte
		publicBytes  []byte
		err          error
	)
	privateBytes, err = ioutil.ReadFile(privateFile)
	if err != nil {
		return privateBytes, publicBytes, err
	}
	publicBytes, err = ioutil.ReadFile(PublicFile)
	if err != nil {
		return privateBytes, publicBytes, err
	}

	return privateBytes, publicBytes, err
}
