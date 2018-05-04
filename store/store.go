package store

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Weex struct {
	AccessID  string
	SecretKey string
}

type MarKetAPI struct {
	Weex Weex
}

var MarKetAPIs MarKetAPI

func Write() {
	file, err := os.OpenFile("data", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	enc := gob.NewEncoder(file)
	err2 := enc.Encode(MarKetAPI{Weex{AccessID: "e07e6ae5-2829-4d81-8121-a36492d102b9", SecretKey: "80DFCA087A4F4D2EBE10E91C1C6C8C5EE1AB7333E3A1B07EC4"}})
	fmt.Println(err2)
}

func Read() error {
	var err error
	file, err := os.Open("data")
	if err == nil {
		dec := gob.NewDecoder(file)
		err = dec.Decode(&MarKetAPIs)

	}
	defer file.Close()
	return err
}
func init() {
	Read()
}
