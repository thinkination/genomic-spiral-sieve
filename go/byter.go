package main

import (

	"fmt"
	"unsafe"
)

func AppendSymbol(b *byte, value *byte) {	
	*b = *b << 2
	*b = *b | *value
}

func printBits(b byte) {
	
	fmt.Printf("%.8b",b)
	fmt.Println()
}

func compress(seq *string) (*[]byte) {
	var master []byte
	codons := *seq

	var base byte = byte(0)
	// var bits byte
	var bitCounter uint8 = uint8(0);
	for i:=0 ; i<len(codons) ; i++ {
		code := codons[i]
		switch(code){
			case 65: //ASCI character code for A
				
				base <<= 2
				base |= 0 // Zero (00) is the hashed value for A 
				
			case 84: //ASCI character code for T
				
				base <<= 2
				base |= 1 // One (01) is the hashed value for T
				
			case 71: //ASCI character code for G
				
				base <<= 2
				base |= 2 // Two (10) is the hashed value for G
				
			case 67: //ASCI character code for C
				
				base <<= 2
				base |= 3 // Three (11) is the hashed value for C
				
		}
		

		bitCounter+=2
		if(bitCounter>=8){
			master=append(master,base)
			bitCounter = uint8(0)
			base = byte(0)
		}
		
	}

	return &master
}

func simple() {
	codons := "ATGC"+"TGAC"+"ATGC"+"TGAC"+"ATGC"+"TGAC"
	
	// master := make([]byte, len(codons)/4)

	// printBits(base)
	fmt.Println(len(codons))

	var master []byte
	master = *(compress(&codons))
	fmt.Println(codons)
		
	var newSize,oldSize uint


	fmt.Println()
	for _,masterByte := range master{
		newSize+= uint(unsafe.Sizeof(masterByte))
		printBits(masterByte);
	}

	for i:=0 ; i<len(codons) ; i++ {
		code := codons[i]
		oldSize += uint(unsafe.Sizeof(code))
	}
	fmt.Println("Size before compression :",oldSize)
	fmt.Println("Size after compression :",newSize)
}

func basics() {
	codons := "ATGC"
	// ucodon := (uint8)(codons)

	//whole string
	fmt.Println("Whole String")
	fmt.Println(unsafe.Sizeof(codons))
	fmt.Println("Single symbol")
	fmt.Println(unsafe.Sizeof(codons[0]))

	bytes := ([]byte)(codons)
	var b byte = (byte)(0)
	var addendum byte 
	printBits(b)
	
	addendum = byte(1)
	AppendSymbol(&b, &addendum)
	
	printBits(b)

	addendum = byte(2)

	AppendSymbol(&b, &addendum)
	printBits(b)

	// fmt.Println(len(bytes))
	for i:=0 ; i< len(bytes) ; i++ {
		b := bytes[i];

		fmt.Printf("%.4b ",b)		
		fmt.Println(codons[i])
	}
	
	// fmt.Printf("%d",bytes)
}

func main() {
	simple()
	// k := byte(17)
	// x := k >> 2
	// fmt.Printf("%.4b",x)
}