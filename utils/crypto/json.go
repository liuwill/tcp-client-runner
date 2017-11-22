package crypto

import "encoding/json"

type JsonCoder struct{}

func (coder *JsonCoder) Decode(ciphertext []byte, target interface{}) error {
	return json.Unmarshal(ciphertext, target)
}

func (coder *JsonCoder) Encode(content interface{}) ([]byte, error) {
	return json.Marshal(content)
}

func init() {
	coderRegister["json"] = &JsonCoder{}
}
