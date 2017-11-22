package crypto

import (
	"encoding/json"
	"log"

	"tcp-client-runner/resource"

	"github.com/golang/protobuf/proto"
)

type ProtobufCoder struct{}

func (coder *ProtobufCoder) Decode(ciphertext []byte, target interface{}) error {
	requestData := &resource.RequestModule{}
	err := proto.Unmarshal(ciphertext, requestData)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
		return err
	}
	messageBody := map[string]interface{}{}

	if len(requestData.Body) > 0 {
		innerErr := json.Unmarshal([]byte(requestData.Body), &messageBody)
		if innerErr != nil {
			return innerErr
		}
	}

	targetData := map[string]interface{}{
		"type": requestData.Type,
		"body": messageBody,
	}
	targetByte, _ := json.Marshal(targetData)
	json.Unmarshal(targetByte, target)

	return nil
}

func (coder *ProtobufCoder) Encode(content interface{}) ([]byte, error) {
	responseContent, _ := content.(map[string]interface{})

	data, _ := responseContent["data"]

	dataContent, _ := json.Marshal(data)

	rawStatus, _ := responseContent["status"]
	rawType, _ := responseContent["type"]
	rawMsg, _ := responseContent["msg"]

	rStatus, _ := rawStatus.(int)
	rType, _ := rawType.(string)
	rMsg, _ := rawMsg.(string)

	newResponse := &resource.ResponseModule{
		Type:   rType,
		Data:   string(dataContent),
		Status: int32(rStatus),
		Msg:    rMsg,
	}

	return proto.Marshal(newResponse)
}

func init() {
	coderRegister["protobuf"] = &ProtobufCoder{}
}
