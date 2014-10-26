package main

import (

	"fmt"
	"unsafe"
)

func main() {
	codons := "ATGC"
	// ucodon := (uint8)(codons)
	bytes := ([]byte)(codons)
	// fmt.Println(len(bytes))
	for i:=0 ; i< len(bytes) ; i++ {
		b := bytes[i];
		fmt.Printf("%.4b ",b)
		fmt.Println(unsafe.Sizeof(b))
		fmt.Println(codons[i])
		fmt.Println(unsafe.Sizeof(codons[i]))
		fmt.Println()
	}
	
	// fmt.Printf("%d",bytes)
}