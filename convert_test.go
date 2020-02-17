package cryptoconv

import (
	"log"
	"testing"
)

func TestFromWaves(t *testing.T) {
	res := From("100000000", "WVS")
	log.Println(res)
}

func TestToWaves(t *testing.T) {
	res := To("1", "WVS")
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
