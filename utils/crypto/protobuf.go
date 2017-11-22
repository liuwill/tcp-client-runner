package crypto

import (
	"encoding/json"
	"game-backend-server/resource"
	"log"

	"github.com/golang/protobuf/proto"
)

type ProtobufCoder struct{}

func (coder *ProtobufCoder) Decode(ciphertext []byte, target interface{}) error {
	responseData := &resource.ResponseModule{}
	err := proto.Unmarshal(ciphertext, responseData)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
		return err
	}
	messageBody := map[string]interface{}{}

	if len(responseData.Data) > 0 {
		innerErr := json.Unmarshal([]byte(responseData.Data), &messageBody)
		if innerErr != nil {
			return innerErr
		}
	}

	targetData := map[string]interface{}{
		"type":   responseData.Type,
		"status": responseData.Status,
		"msg":    responseData.Msg,
		"data":   messageBody,
	}
	targetByte, _ := json.Marshal(targetData)
	json.Unmarshal(targetByte, target)

	return nil
}

func (coder *ProtobufCoder) Encode(content interface{}) ([]byte, error) {
	responseContent, _ := content.(map[string]interface{})

	body, _ := responseContent["body"]

	bodyContent, _ := json.Marshal(body)

	rawType, _ := responseContent["type"]
	rType, _ := rawType.(string)

	newResponse := &resource.RequestModule{
		Type: rType,
		Body: string(bodyContent),
	}

	return proto.Marshal(newResponse)
}

func init() {
	coderRegister["protobuf"] = &ProtobufCoder{}
}
