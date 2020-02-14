package go_cryptoconv

import (
	"log"
	"testing"
)

func TestFromWaves(t *testing.T) {
	res := FromWaves("200000000", "WVS")
	log.Println(res)
}

func TestToWaves(t *testing.T) {
	res := ToWaves("2", "WVS")
	log.Println(res)
}

func TestToWei(t *testing.T) {
	res := ToWei("1", "Ether")
	log.Println(res)
}

func TestFromWei(t *testing.T) {
	res := FromWei("1000000000000000000", "Ether")
	log.Println(res)
}

func TestToDrop(t *testing.T) {
	res := ToDrop("1", "XRP")
	log.Println(res)
}

func TestFromDrop(t *testing.T) {
	res := FromDrop("1000000", "XRP")
	log.Println(res)
}
