BigInt structure contains __data__ and __ParseBy4Bits__,
- data - is just an array of uint64 which stores big number
- ParseBy4Bits - parameters tells how many chars we use in 1 __bigint.data__ cell. 
For 3 - it is 3 * 4 = 12 bits, for 16 - it is 16 * 4 = 64 bits (uint64).   
 




    bigHex1 := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	bigHex2 := "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c"
	fmt.Println(bigHex1, bigHex2)
    
    // Create bigint    
	bi1 := bigint.BigInt{ParseBy4Bits: 16}
	bi2 := bigint.BigInt{ParseBy4Bits: 16}

Basic Operations:

	// 1. SetHex, GetHex implementation
	bi1.SetHex(bigHex1)
	bi2.SetHex(bigHex2)
	fmt.Println(bi1.GetHex())

SetHex reads hex string and uses __ParseBy4Bits__ to read the exact amount 
of bits in a cell from __bigint.data__. For __ParseBy4Bits = 16__ we get 16 * 4 = 64 bits per a cell.

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