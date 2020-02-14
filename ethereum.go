package go_cryptoconv

import (
	"fmt"
	"math/big"
	"strings"
)

// EthereumFactor type for enum
type EthereumFactor int

const (
	// Wei 1 wei
	Wei EthereumFactor = 0
	// KWei (babbage) 1e3 wei
	KWei EthereumFactor = 3
	// MWei (lovelace) 1e6 wei
	MWei EthereumFactor = 6
	// GWei (shannon) 1e9 wei
	GWei EthereumFactor = 9
	// Szabo (microether) 1e12 wei
	Szabo EthereumFactor = 12
	// Finney (milliether) 1e15 wei
	Finney EthereumFactor = 15
	// Ether 1e18 wei
	Ether EthereumFactor = 18
	// KEther 1e21 wei
	KEther EthereumFactor = 21
	// MEther 1e24 wei
	MEther EthereumFactor = 24
	// GEther 1e27 wei
	GEther EthereumFactor = 27
)

// EthereumUnit Base struct
type EthereumUnit struct {
	Name      string
	WeiFactor *big.Int
}

// FromWei func calculate number / unitfactor return string
func FromWei(number string, unitname string) string {
	unit := new(EthereumUnit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWeiFactor := new(big.Float).SetInt(unit.GetWeiFactor())
	tmp := new(big.Float)
	tmp.Quo(bigFloatNumber, unitWeiFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// ToWei func calculate number * unitfactor return string
func ToWei(number string, unitname string) string {
	unit := new(EthereumUnit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWeiFactor := new(big.Float).SetInt(unit.GetWeiFactor())
	tmp := new(big.Float)
	tmp.Mul(bigFloatNumber, unitWeiFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// InitUnit func param unitname(string) retun EthereumUnit struct
// by default Wei factor
func (unit *EthereumUnit) InitUnit(unitname string) *EthereumUnit {
	name := strings.ToLower(unitname)
	units := map[string]EthereumFactor{
		"wei":    Wei,
		"kwei":   KWei,
		"mwei":   MWei,
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
	unit.WeiFactor = exp

	return unit
}

// GetWeiFactor func return unit.WeiFactor
func (unit *EthereumUnit) GetWeiFactor() *big.Int {
	return unit.WeiFactor
}
