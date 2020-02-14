package go_cryptoconv

import (
	"fmt"
	"math/big"
	"strings"
)

// RippleFactor type for enum
type RippleFactor int

const (
	// Drop 1 drop
	Drop RippleFactor = 0
	// XRP (Ripple) 1e6 drop
	XRP RippleFactor = 6
)

// RippleUnit Base struct
type RippleUnit struct {
	Name       string
	DropFactor *big.Int
}

// FromDrop func calculate number / unitfactor return string
func FromDrop(number string, unitname string) string {
	unit := new(RippleUnit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitDropFactor := new(big.Float).SetInt(unit.GetDropFactor())
	tmp := new(big.Float)
	tmp.Quo(bigFloatNumber, unitDropFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// ToDrop func calculate number * unitfactor return string
func ToDrop(number string, unitname string) string {
	unit := new(RippleUnit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitDropFactor := new(big.Float).SetInt(unit.GetDropFactor())
	tmp := new(big.Float)
	tmp.Mul(bigFloatNumber, unitDropFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// InitUnit func param unitname(string) retun RippleUnit struct
// by default Drop factor
func (unit *RippleUnit) InitUnit(unitname string) *RippleUnit {
	name := strings.ToLower(unitname)
	units := map[string]RippleFactor{
		"drop":  Drop,
		"drops": Drop,
		"xrp":   XRP,
	}

	var unitfactor, exp = big.NewInt(int64(units[name])), big.NewInt(10)
	exp.Exp(exp, unitfactor, nil)
	unit.Name = name
	unit.DropFactor = exp

	return unit
}

// GetDropFactor func return unit.DropFactor
func (unit *RippleUnit) GetDropFactor() *big.Int {
	return unit.DropFactor
}
