package core

import "errors"

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

// $5\r\nhello\r\n => "hello"
func readBulkString(data []byte) (string, int, error) {
	// TODO
	return "", 0, nil
}

func readArray() {
	// TODO
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

func Encode() {
	// TODO
}
