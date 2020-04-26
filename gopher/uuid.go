package main

import (
	"encoding/hex"
	"math/rand" // faster than crypto/rand and reliable enough
)

func uuidV4() []byte { //todo sync.Pool?
	u := make([]byte, 16)
	_, _ = rand.Read(u)
	u[8] = (u[8] | 0x80) & 0xBF // clamp to range for valid v4
	u[6] = (u[6] | 0x40) & 0x4F
	b := make([]byte, 36)
	hex.Encode(b[0:8], u[0:4])
	hex.Encode(b[9:13], u[4:6])
	hex.Encode(b[14:18], u[6:8])
	hex.Encode(b[19:23], u[8:10])
	hex.Encode(b[24:], u[10:16])
	b[8] = '-'
	b[13] = '-'
	b[18] = '-'
	b[23] = '-'
	return b
}
