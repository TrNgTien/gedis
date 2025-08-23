package core

import (
	"bytes"
	"errors"
	"fmt"
	"gedis/internal/constant"
)

const CRLF string = "\r\n"

// +OK\r\n => OK, 5 bytes
func readSimpleString(data []byte) (string, int, error) {
	start := 1
	for data[start] != '\r' {
		start++
	}

	return string(data[1:start]), start + 2, nil
}

// :123\r\n => 123, 5 bytes
func readInt64(data []byte) (int64, int, error) {
	var rs int64 = 0

	start := 1
	for data[start] != '\r' {
		rs = rs*10 + int64(data[start]-'0')
		start++
	}

	return rs, start + 2, nil
}
func readSimpleErrors(data []byte) (string, int, error) {
	return readSimpleString(data)
}

// $5\r\nhello\r\n => 5, 4
func readLen(data []byte) (int, int) {
	len, pos, _ := readInt64(data)
	return int(len), pos
}

// $5\r\nhello\r\n => "hello"
func readBulkString(data []byte) (string, int, error) {
	len, pos := readLen(data)

	return string(data[pos:(pos + len)]), pos + 2, nil
}

func Decode(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("no data")
	}

	switch string(data[0]) {
	case "+":
		return readSimpleString(data)
	case "-":
		return readSimpleErrors(data)
	case ":":
		return readInt64(data)
	case "$":
		return readBulkString(data)

	}
	return nil, 0, nil
}

func encodeString(v string) []byte {
	return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v))
}

// ["SET", "mykey", "Hello"]
// *3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nHello\r\n
func encodeStringArray(arr []string) []byte {
	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf("*%d\r\n", len(arr)))

	for _, s := range arr {
		buf.Write(encodeString(s))
	}
	return buf.Bytes()
}

func Encode(value interface{}) []byte {
	switch v := value.(type) {
	case []string:
		return encodeStringArray(value.([]string))
	case int64, int32, int16, int8, int:
		return []byte(fmt.Sprintf(":\r\n", v))
	case error:
		return []byte(fmt.Sprintf("-%s\r\n", v))
	default:
		return constant.RespNil
	}

}
