package crypto

import (
	"errors"
)

const (
	CODE_FORMAT_JSON     = "json"
	CODE_FORMAT_PROTOBUF = "protobuf"
)

type Coder interface {
	Decode(ciphertext []byte, target interface{}) error
	Encode(content interface{}) ([]byte, error)
}

var (
	coderRegister = map[string]Coder{}
)

func GenerateCoder(format string) (Coder, error) {
	if coder, ok := coderRegister[format]; ok {
		return coder, nil
	}
	return nil, errors.New("coder is not exist")
}

func IsCoderExist(format string) bool {
	for key, _ := range coderRegister {
		if key == format {
			return true
		}
	}

	return false
}
