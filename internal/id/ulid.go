package id

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/oklog/ulid/v2"
	"time"
)

type ulidGenerator struct {
}

func (s *ulidGenerator) Next() (string, error) {
	entropy := &ulid.LockedMonotonicReader{
		MonotonicReader: ulid.Monotonic(rand.Reader, 0),
	}
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}
	bytes, err := id.MarshalBinary()
	if err != nil {
		return "", err
	}
	var buf [36]byte
	encodeHex(buf[:], bytes)
	return string(buf[:]), nil
}

// Copy from github.com/google/uuid
func encodeHex(dst []byte, uuid []byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

var _ulidGenerator Generator

// ULIDGenerator creates a new id generator
func ULIDGenerator() Generator {
	if _ulidGenerator == nil {
		generator := Generator(&ulidGenerator{})
		_ulidGenerator = generator
	}
	return _ulidGenerator
}
