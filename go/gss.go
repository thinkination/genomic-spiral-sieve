//Thanks to http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write#18479916

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "io/ioutil"
)

var bitMap map[byte]string = map[byte]string{ byte(27):"ATGC" , byte(30):"ATCG" , byte(39):"AGTC" , byte(45):"AGCT" , byte(54):"ACTG" , byte(57):"ACGT" , byte(75):"TAGC" , byte(78):"TACG" , byte(99):"TGAC" , byte(108):"TGCA" , byte(114):"TCAG" , byte(120):"TCGA" , byte(135):"GATC" , byte(141):"GACT" , byte(147):"GTAC" , byte(156):"GTCA" , byte(177):"GCAT" , byte(180):"GCTA" , byte(198):"CATG" , byte(201):"CAGT" , byte(210):"CTAG" , byte(216):"CTGA" , byte(225):"CGAT" , byte(228):"CGTA"}

//TODO fix the decommression piece
func decompress(bytes *[]byte, n int) (*[]string) {
	byts := *bytes
	// fmt.Println(byts)
	quads := make([]string,n) 
	for i,byt := range byts {
		if bitMap[byt] != "" {
			quads[i] = bitMap[byt]	
		}		
	}

	/*
		bitMap[byte(00011011)=ATGC
		bitMap[byte(00011110)=ATCG
		bitMap[byte(00100111)=AGTC
		bitMap[byte(00101101)=AGCT
		bitMap[byte(00110110)=ACTG
		bitMap[byte(00111001)=ACGT
		bitMap[byte(01001011)=TAGC
		bitMap[byte(01001110)=TACG
		bitMap[byte(01100011)=TGAC
		bitMap[byte(01101100)=TGCA
		bitMap[byte(01110010)=TCAG
		bitMap[byte(01111000)=TCGA
		bitMap[byte(10000111)=GATC
		bitMap[byte(10001101)=GACT
		bitMap[byte(10010011)=GTAC
		bitMap[byte(10011100)=GTCA
		bitMap[byte(10110001)=GCAT
		bitMap[byte(10110100)=GCTA
		bitMap[byte(11000110)=CATG
		bitMap[byte(11001001)=CAGT
		bitMap[byte(11010010)=CTAG
		bitMap[byte(11011000)=CTGA
		bitMap[byte(11100001)=CGAT
		bitMap[byte(11100100)=CGTA

	*/
	// x := byte(27)
	// fmt.Println(bitMap[x])
	return &quads
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// A 1000001 after compression 00
// T 1010100 after compression 01 
// G 1000111 after compression 10
// C 1000011 after compression 11

//ATGC after compression looks like 00011011
//To decompress - 00 01 10 11 = A T G C 


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

func printBits(b byte) {
	
	fmt.Printf("%.8b",b)
	fmt.Println()
}


func main() {
	
	lines, err := readLines("../seq.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {
		var master []byte	
		master = *(compress(&line))

		// for _,masterByte := range master{			
		// 	printBits(masterByte);		
		// }

		decompressed := *(decompress(&master, len(master)))
		for _,symbol:= range decompressed{
			fmt.Print(symbol)
			fmt.Print(" ")
		}
		// err = ioutil.WriteFile("./dat1", master, 0644)
  //   	if err != nil {
  //       	panic(err)
  //   	 }    		

	}


	// if err := writeLines(lines, "foo.out.txt"); err != nil {
	// 	log.Fatalf("writeLines: %s", err)
	// }
}
