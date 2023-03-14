package main

import (
	"baskets/device"
	"baskets/keygen"
	"fmt"
)

func main() {
	dev := device.EmuDevice{}
	kg, err := keygen.NewKeygen(dev, "correct horse battery staple")
	if err != nil {
		panic(err)
	}
	fmt.Printf("bk = %x\n", kg.BaseKey)

	sk, err := kg.SiteKey("lichess")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sk = %x\n", sk)

	sk, err = kg.SiteKey("lobste.rs")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sk = %x\n", sk)

	sk, err = kg.SiteKey("lichess")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sk = %x\n", sk)
}
