package device

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/pbkdf2"
)

type EmuDevice struct {
}

func emulateDeviceRaw(service []byte) []byte {
	fmt.Printf("[dev emu] request for \"%s\"\n", service)
	if len(service) == 0 {
		return []byte("I AM SOMETHING")
	}
	// BIP-39
	// final secret flashed onto firmware
	mnemonic := "ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic"
	mnemonic_pass := "TREZOR"
	secret := pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+mnemonic_pass), 2048, 32, sha512.New)

	b, err := blake2s.New256(secret)
	if err != nil {
		panic(err)
	}
	b.Write(service)
	key := b.Sum(nil)
	fmt.Printf("[dev emu] outputting %X...\n", key[:4])
	return []byte(fmt.Sprintf("%X", b.Sum(nil)))
}

func (dev EmuDevice) GetKey(service string) ([]byte, error) {
	res := emulateDeviceRaw([]byte(service))
	if len(res) != 64 {
		return nil, errors.New("invalid device response")
	}
	key := make([]byte, 32)
	n, err := hex.Decode(key, res)
	if err != nil {
		return nil, err
	}
	if n != 32 {
		return nil, errors.New("invalid device response")
	}
	return key, nil
}
