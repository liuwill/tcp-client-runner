package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Pack(source []byte) ([]byte, error) {
	// 读取消息的长度
	var length int32 = int32(len(source))
	var pkg *bytes.Buffer = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, source)
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Unpack(reader *bufio.Reader) ([]byte, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return []byte{}, err
	}
	if int32(reader.Buffered()) < length+4 {
		return []byte{}, err
	}

	// 读取消息真正的内容
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return []byte{}, err
	}
	return pack[4:], nil
}
