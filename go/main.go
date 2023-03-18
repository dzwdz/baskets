package main

import (
	"baskets/device"
	"baskets/keygen"
	"baskets/store"
	"fmt"

	"os"
)

func main() {
	data, err := os.ReadFile("../example.json")
	if err != nil {
		panic(err)
	}

	var ent store.Entry
	err = store.Parse(data, &ent)
	if err != nil {
		panic(err)
	}

	dev := device.EmuDevice{}
	pass := "correct horse battery staple"
	kg, err := keygen.NewKeygen(dev, pass, ent.Header.Scrypt)
	if err != nil {
		panic(err)
	}

	err = ent.Header.CheckCompat(kg)
	if err != nil {
		panic(err)
	}

	sk, err := kg.SiteKey(ent.Site, ent.Offset)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sk = %x\n", sk)
}
