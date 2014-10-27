//Thanks to http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write#18479916

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

		for _,masterByte := range master{			
			printBits(masterByte);
		}		
	}

	if err := writeLines(lines, "foo.out.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
