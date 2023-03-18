package keygen

import (
	"baskets/device"
	"crypto/sha256"
	"golang.org/x/crypto/scrypt"
)

type Keygen struct {
	dev device.Device
	DevKey []byte
	BaseKey []byte
	Params [3]int
}

func (kg *Keygen) SiteKey(site string, offset string) ([]byte, error) {
	dkey, err := kg.dev.GetKey(site)
	if err != nil {
		return nil, err
	}

	toHash := make([]byte, len(kg.BaseKey))
	copy(toHash, kg.BaseKey)
	toHash = append(toHash, dkey...)
	toHash = append(toHash, offset...)
	key := sha256.Sum256(toHash)
	return key[:], nil
}

func NewKeygen(dev device.Device, pass string, scrypt_params [3]int) (*Keygen, error) {
	dk, err := dev.GetKey("baskets")
	if err != nil {
		return nil, err
	}

	kg := new(Keygen)
	kg.dev = dev
	kg.DevKey = dk
	kg.Params = scrypt_params
	sp := scrypt_params
	kg.BaseKey, err = scrypt.Key([]byte(pass), dk, sp[0], sp[1], sp[2], 32)
	if err != nil {
		return nil, err
	}

	return kg, nil
}
