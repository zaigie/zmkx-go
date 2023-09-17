package utils

import (
	"bytes"
	"encoding/binary"
	"errors"

	"google.golang.org/protobuf/proto"
)

func DelimitedEncode(message proto.Message) ([]byte, error) {
	serialized, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	varintBuf := make([]byte, binary.MaxVarintLen64) // 用于存放varint的缓冲区
	n := binary.PutUvarint(varintBuf, uint64(len(serialized)))
	buf.Write(varintBuf[:n]) // 写入varint
	buf.Write(serialized)

	return buf.Bytes(), nil
}

func DelimitedDecode(buffer []byte) ([]byte, error) {
	size, n := binary.Uvarint(buffer)
	if n <= 0 {
		return nil, errors.New("failed to decode varint")
	}

	if int(size)+n > len(buffer) {
		return nil, errors.New("buffer too small")
	}

	msg := buffer[n : n+int(size)]
	return msg, nil
}

func Ljust(buf []byte, fill byte, fillSize int) []byte {
	if len(buf) >= fillSize {
		return buf
	}

	// 创建一个新的切片，大小为 fillSize，并使用 fill 字节进行填充
	filled := make([]byte, fillSize)
	for i := range filled {
		filled[i] = fill
	}

	// 将原始的 buf 复制到新的切片中
	copy(filled, buf)

	return filled
}
