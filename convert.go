package cryptoconv

import (
	"fmt"
	"math/big"
	"strings"
)

// Factor type for enum
type Factor int

const (
	// Wei 1 wei
	Wei Factor = 0
	// Drop 1 drop
	Drop Factor = 0
	// Waves 1 waves
	Waves Factor = 0
	// KWei (babbage) 1e3 wei
	KWei Factor = 3
	// MWei (lovelace) 1e6 wei
	MWei Factor = 6
	// XRP (Ripple) 1e6 drop
	XRP Factor = 6
	// WVS 1e8 waves
	WVS Factor = 8
	// GWei (shannon) 1e9 wei
	GWei Factor = 9
	// Szabo (microether) 1e12 wei
	Szabo Factor = 12
	// Finney (milliether) 1e15 wei
	Finney Factor = 15
	// Ether 1e18 wei
	Ether Factor = 18
	// KEther 1e21 wei
	KEther Factor = 21
	// MEther 1e24 wei
	MEther Factor = 24
	// GEther 1e27 wei
	GEther Factor = 27
)

// Unit Base struct
type Unit struct {
	Name   string
	Factor *big.Int
}

// InitUnit func param unitname(string) retun Unit struct
// by default 0 factor
func (unit *Unit) InitUnit(unitname string) *Unit {
	name := strings.ToLower(unitname)
	units := map[string]Factor{
		"wei":    Wei,
		"drop":   Drop,
		"waves":  Waves,
		"kwei":   KWei,
		"mwei":   MWei,
		"xrp":    XRP,
		"wvs":    WVS,
		"gwei":   GWei,
		"szabo":  Szabo,
		"finney": Finney,
		"ether":  Ether,
		"kether": KEther,
		"mether": MEther,
		"gether": GEther,
	}

	var unitfactor, exp = big.NewInt(int64(units[name])), big.NewInt(10)
	exp.Exp(exp, unitfactor, nil)
	unit.Name = name
	unit.Factor = exp

	return unit
}

// GetFactor func return unit.Factor
func (unit *Unit) GetFactor() *big.Int {
	return unit.Factor
}

// From func calculate number / unitfactor return string
func From(number string, unitname string) string {
	unit := new(Unit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitFactor := new(big.Float).SetInt(unit.GetFactor())
	tmp := new(big.Float)
	tmp.Quo(bigFloatNumber, unitFactor)
	tmpstr := fmt.Sprintf("%.8f", tmp)

	return tmpstr
}

// To func calculate number * unitfactor return string
func To(number string, unitname string) string {
	unit := new(Unit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitFactor := new(big.Float).SetInt(unit.GetFactor())
	tmp := new(big.Float)
	tmp.Mul(bigFloatNumber, unitFactor)
	tmpstr := fmt.Sprintf("%.8f", tmp)

	return tmpstr
}
