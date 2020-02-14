package go_cryptoconv

import (
	"fmt"
	"math/big"
	"strings"
)

// WavesFactor type for enum
type WavesFactor int

const (
	// Waves 1 waves
	Waves WavesFactor = 0
	// WVS 1e8 waves
	WVS WavesFactor = 8
)

// WavesUnit Base struct
type WavesUnit struct {
	Name        string
	WavesFactor *big.Int
}

// FromWaves func calculate number / unitfactor return string
func FromWaves(number string, unitname string) string {
	unit := new(WavesUnit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWavesFactor := new(big.Float).SetInt(unit.GetWavesFactor())
	tmp := new(big.Float)
	tmp.Quo(bigFloatNumber, unitWavesFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// ToWaves func calculate number * unitfactor return string
func ToWaves(number string, unitname string) string {
	unit := new(WavesUnit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWavesFactor := new(big.Float).SetInt(unit.GetWavesFactor())
	tmp := new(big.Float)
	tmp.Mul(bigFloatNumber, unitWavesFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// InitUnit func param unitname(string) retun WavesUnit struct
// by default Waves factor
func (unit *WavesUnit) InitUnit(unitname string) *WavesUnit {
	name := strings.ToLower(unitname)
	units := map[string]WavesFactor{
		"waves": Waves,
		"wvs":   WVS,
	}

	var unitfactor, exp = big.NewInt(int64(units[name])), big.NewInt(10)
	exp.Exp(exp, unitfactor, nil)
	unit.Name = name
	unit.WavesFactor = exp

	return unit
}

// GetWavesFactor func return unit.WavesFactor
func (unit *WavesUnit) GetWavesFactor() *big.Int {
	return unit.WavesFactor
}
