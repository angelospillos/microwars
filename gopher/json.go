package main

import "strconv"

var jsonTemplate = []byte(`{"uuid":"","fib":""}`)

func toJson(uuid []byte, fib uint64) []byte {
	buf := make([]byte, 54, 64)
	copy(buf, jsonTemplate[0:9])
	copy(buf[9:], uuid)
	copy(buf[45:], jsonTemplate[9:18])
	buf = append(buf, strconv.FormatUint(fib, 10)... /* fast path for fib <= 10 */)
	buf = append(buf, jsonTemplate[18:]...)
	return buf
}
