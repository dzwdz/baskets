package keygen

import (
	"baskets/device"
	"crypto/sha256"
	"golang.org/x/crypto/scrypt"
)

type Keygen struct {
	dev device.Device
	BaseKey []byte
}

func (kg *Keygen) SiteKey(site string) ([]byte, error) {
	dkey, err := kg.dev.GetKey(site)
	if err != nil {
		return nil, err
	}

	toHash := make([]byte, len(kg.BaseKey))
	copy(toHash, kg.BaseKey)
	toHash = append(toHash, dkey...)
	toHash = append(toHash, "0"...)
	key := sha256.Sum256(toHash)
	return key[:], nil
}

func NewKeygen(dev device.Device, pass string) (*Keygen, error) {
	// TODO do some caching, so you don't have to repress the button on incorrect passwords
	// lowkey a security issue because of the mitm possibility
	salt, err := dev.GetKey("baskets")
	if err != nil {
		return nil, err
	}

	kg := new(Keygen)
	kg.dev = dev
	// scrypt parameters are stupidly small for demo purposes.
	kg.BaseKey, err = scrypt.Key([]byte(pass), salt, 1024, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return kg, nil
}
