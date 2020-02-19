package cryptoconv

import (
	"log"
	"testing"
)

func TestFromWaves(t *testing.T) {
	res := From("100000", "WVS")
	log.Println(res)
}

func TestToWaves(t *testing.T) {
	res := To("0.001", "WVS")
	log.Println(res)
}

func TestFromWei(t *testing.T) {
	res := From("1000000000000000000", "Ether")
	log.Println(res)
}

func TestToWei(t *testing.T) {
	res := To("1", "Ether")
	log.Println(res)
}

func TestFromDrop(t *testing.T) {
	res := From("1000000", "XRP")
	log.Println(res)
}

func TestToDrop(t *testing.T) {
	res := To("1", "XRP")
	log.Println(res)
}

func TestFromPrizm(t *testing.T) {
	res := From("100", "PZM")
	log.Println(res)
}

func TestToPrizm(t *testing.T) {
	res := To("1", "PZM")
	log.Println(res)
}
