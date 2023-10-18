package main

import (
	"fmt"
	"go-large-math/bigint"
)

func main() {
	bigHex1 := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	bigHex2 := "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c"
	fmt.Println(bigHex1, bigHex2)

	bi1 := bigint.BigInt{ParseBy4Bits: 16}
	bi2 := bigint.BigInt{ParseBy4Bits: 16}

	// 1. SetHex, GetHex implementation
	bi1.SetHex(bigHex1)
	bi2.SetHex(bigHex2)
	fmt.Println(bi1.GetHex())

	// 2. Xor, Inv, Or, And, ShiftR, ShiftL
	bi1.XOR(&bi2)
	fmt.Println("XOR", bi1.GetHex())
	bi1.ShiftR(10)
	fmt.Println("ShiftR >> 10", bi1.GetHex())
	bi1.ShiftR(1)
	fmt.Println("ShiftR >> 1", bi1.GetHex())
	bi1.INV()
	fmt.Println("INV", bi1.GetHex())
	bi1.OR(&bi2)
	fmt.Println("OR", bi1.GetHex())
	bi1.AND(&bi2)
	fmt.Println("AND", bi1.GetHex())

}
