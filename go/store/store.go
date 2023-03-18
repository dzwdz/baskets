package store

import (
	"encoding/json"

	// used for the checksum
	"baskets/keygen"
	"crypto/sha256"
	"encoding/binary"
	"errors"
)

type Header struct {
	Scrypt [3]int `json:"scrypt"`
	DevKeyChecksum uint32 `json:"devKeyChecksum"`
	BaseKeyChecksum uint32 `json:"baseKeyChecksum"`
}

type Entry struct {
	Header *Header `json:"header,omitempty"`
	Site string `json:"site"`
	Offset string `json:"offset"`
}

func Parse(data []byte, ent *Entry) error {
	return json.Unmarshal(data, ent)
}


func checksum(key []byte) uint32 {
	hash := sha256.Sum256(key)
	return binary.BigEndian.Uint32(hash[len(hash)-4:])
}

func (hdr *Header) CheckCompat(kg *keygen.Keygen) error {
	if hdr.DevKeyChecksum != checksum(kg.DevKey) {
		return errors.New("bad device key (wrong device?)")
	}
	if hdr.BaseKeyChecksum != checksum(kg.BaseKey) {
		return errors.New("bad base key (incorrect passphrase?)")
	}
	if hdr.Scrypt != kg.Params {
		return errors.New("bad scrypt params")
	}
	return nil
}
