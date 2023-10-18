package bigint

import (
	"fmt"
	"strconv"
)

type BigInt struct {
	data         []uint64
	ParseBy4Bits uint8
}

func (bi *BigInt) SetHex(s string) {
	parseBy4Bits := int(bi.ParseBy4Bits)
	intPart := len(s) / parseBy4Bits

	for i := 0; i < intPart; i++ {
		partitionStr := s[i*parseBy4Bits : (i*parseBy4Bits)+parseBy4Bits]
		err := parseAndAppendStr(partitionStr, bi)
		if err != nil {
			fmt.Println("Error", err)
		}
	}

	if len(s) > intPart*16 {
		partitionStr := s[intPart*16:]
		err := parseAndAppendStr(partitionStr, bi)
		if err != nil {
			fmt.Println("Error", err)
		}
	}
}

func parseAndAppendStr(partitionStr string, bi *BigInt) error {
	parsedUint64, err := strconv.ParseUint(partitionStr, 16, 64)
	if err != nil {
		return err
	}
	bi.data = append(bi.data, parsedUint64)
	return nil
}

func (bi *BigInt) GetHex() string {
	var result string

	for _, num := range bi.data {
		hexStr := fmt.Sprintf("%x", num)
		result += hexStr
	}

	return result
}

func (bi *BigInt) XOR(bi2 *BigInt) {
	for i := 0; i < len(bi.data); i++ {
		bi.data[i] = bi.data[i] ^ bi2.data[i]
	}
}

func (bi *BigInt) INV() {
	for i := 0; i < len(bi.data); i++ {
		bi.data[i] = ^bi.data[i]
	}
}

func (bi *BigInt) AND(bi2 *BigInt) {
	for i := 0; i < len(bi.data); i++ {
		bi.data[i] = bi.data[i] & bi2.data[i]
	}
}

func (bi *BigInt) OR(bi2 *BigInt) {
	for i := 0; i < len(bi.data); i++ {
		bi.data[i] = bi.data[i] | bi2.data[i]
	}
}

func (bi *BigInt) ShiftR(numShift uint) {
	var mask uint64 = (1 << numShift) - 1

	for i := len(bi.data) - 1; i >= 0; i-- {
		// make shift
		bi.data[i] = bi.data[i] >> numShift

		// add bits from another cell to result
		if i > 0 {
			nextData := bi.data[i-1]

			nextDataMasked := nextData & mask
			// maybe here bi.ParseBy4Bits - numShift
			nextDataMasked = nextDataMasked << (uint(bi.ParseBy4Bits*4) - numShift)

			//bi.data[i] = bi.data[i] & mask
			bi.data[i] = bi.data[i] | nextDataMasked
		}
	}
}
