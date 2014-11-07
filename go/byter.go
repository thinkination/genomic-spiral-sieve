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
	// fmt.Println()
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

func encode(quadBase *string ) *byte {
	var base byte

	for _,char := range (*quadBase){
		// fmt.Println(char)
		switch(char){
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
	}

	return &base
}

func base() {
	quads := "ATGC"+"ATCG"+"AGTC"+"AGCT"+"ACTG"+"ACGT"+"TAGC"+"TACG"+"TGAC"+"TGCA"+"TCAG"+"TCGA"+"GATC"+"GACT"+"GTAC"+"GTCA"+"GCAT"+"GCTA"+"CATG"+"CAGT"+"CTAG"+"CTGA"+"CGAT"+"CGTA"
	values := [...]string{"ATGC","ATCG","AGTC","AGCT","ACTG","ACGT","TAGC","TACG","TGAC","TGCA","TCAG","TCGA","GATC","GACT","GTAC","GTCA","GCAT","GCTA","CATG","CAGT","CTAG","CTGA","CGAT","CGTA"}
	quads = "AAAA"+"AAAT"+"AAAG"+"AAAC"+"AATA"+"AATT"+"AATG"+"AATC"+"AAGA"+"AAGT"+"AAGG"+"AAGC"+"AACA"+"AACT"+"AACG"+"AACC"+"ATAA"+"ATAT"+"ATAG"+"ATAC"+"ATTA"+"ATTT"+"ATTG"+"ATTC"+"ATGA"+"ATGT"+"ATGG"+"ATGC"+"ATCA"+"ATCT"+"ATCG"+"ATCC"+"AGAA"+"AGAT"+"AGAG"+"AGAC"+"AGTA"+"AGTT"+"AGTG"+"AGTC"+"AGGA"+"AGGT"+"AGGG"+"AGGC"+"AGCA"+"AGCT"+"AGCG"+"AGCC"+"ACAA"+"ACAT"+"ACAG"+"ACAC"+"ACTA"+"ACTT"+"ACTG"+"ACTC"+"ACGA"+"ACGT"+"ACGG"+"ACGC"+"ACCA"+"ACCT"+"ACCG"+"ACCC"+"TAAA"+"TAAT"+"TAAG"+"TAAC"+"TATA"+"TATT"+"TATG"+"TATC"+"TAGA"+"TAGT"+"TAGG"+"TAGC"+"TACA"+"TACT"+"TACG"+"TACC"+"TTAA"+"TTAT"+"TTAG"+"TTAC"+"TTTA"+"TTTT"+"TTTG"+"TTTC"+"TTGA"+"TTGT"+"TTGG"+"TTGC"+"TTCA"+"TTCT"+"TTCG"+"TTCC"+"TGAA"+"TGAT"+"TGAG"+"TGAC"+"TGTA"+"TGTT"+"TGTG"+"TGTC"+"TGGA"+"TGGT"+"TGGG"+"TGGC"+"TGCA"+"TGCT"+"TGCG"+"TGCC"+"TCAA"+"TCAT"+"TCAG"+"TCAC"+"TCTA"+"TCTT"+"TCTG"+"TCTC"+"TCGA"+"TCGT"+"TCGG"+"TCGC"+"TCCA"+"TCCT"+"TCCG"+"TCCC"+"GAAA"+"GAAT"+"GAAG"+"GAAC"+"GATA"+"GATT"+"GATG"+"GATC"+"GAGA"+"GAGT"+"GAGG"+"GAGC"+"GACA"+"GACT"+"GACG"+"GACC"+"GTAA"+"GTAT"+"GTAG"+"GTAC"+"GTTA"+"GTTT"+"GTTG"+"GTTC"+"GTGA"+"GTGT"+"GTGG"+"GTGC"+"GTCA"+"GTCT"+"GTCG"+"GTCC"+"GGAA"+"GGAT"+"GGAG"+"GGAC"+"GGTA"+"GGTT"+"GGTG"+"GGTC"+"GGGA"+"GGGT"+"GGGG"+"GGGC"+"GGCA"+"GGCT"+"GGCG"+"GGCC"+"GCAA"+"GCAT"+"GCAG"+"GCAC"+"GCTA"+"GCTT"+"GCTG"+"GCTC"+"GCGA"+"GCGT"+"GCGG"+"GCGC"+"GCCA"+"GCCT"+"GCCG"+"GCCC"+"CAAA"+"CAAT"+"CAAG"+"CAAC"+"CATA"+"CATT"+"CATG"+"CATC"+"CAGA"+"CAGT"+"CAGG"+"CAGC"+"CACA"+"CACT"+"CACG"+"CACC"+"CTAA"+"CTAT"+"CTAG"+"CTAC"+"CTTA"+"CTTT"+"CTTG"+"CTTC"+"CTGA"+"CTGT"+"CTGG"+"CTGC"+"CTCA"+"CTCT"+"CTCG"+"CTCC"+"CGAA"+"CGAT"+"CGAG"+"CGAC"+"CGTA"+"CGTT"+"CGTG"+"CGTC"+"CGGA"+"CGGT"+"CGGG"+"CGGC"+"CGCA"+"CGCT"+"CGCG"+"CGCC"+"CCAA"+"CCAT"+"CCAG"+"CCAC"+"CCTA"+"CCTT"+"CCTG"+"CCTC"+"CCGA"+"CCGT"+"CCGG"+"CCGC"+"CCCA"+"CCCT"+"CCCG"+"CCCC"
	
	master := *(compress(&quads))
	for i,masterByte := range master{	
		// fmt.Print("bitMap[byte(")	
		fmt.Print(masterByte)	
		// fmt.Print(")]=")	
		fmt.Print(" ")
		fmt.Print(values[i])
		fmt.Println()

	}
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

	combinations := []string{"AAAA","AAAT","AAAG","AAAC","AATA","AATT","AATG","AATC","AAGA","AAGT","AAGG","AAGC","AACA","AACT","AACG","AACC","ATAA","ATAT","ATAG","ATAC","ATTA","ATTT","ATTG","ATTC","ATGA","ATGT","ATGG","ATGC","ATCA","ATCT","ATCG","ATCC","AGAA","AGAT","AGAG","AGAC","AGTA","AGTT","AGTG","AGTC","AGGA","AGGT","AGGG","AGGC","AGCA","AGCT","AGCG","AGCC","ACAA","ACAT","ACAG","ACAC","ACTA","ACTT","ACTG","ACTC","ACGA","ACGT","ACGG","ACGC","ACCA","ACCT","ACCG","ACCC","TAAA","TAAT","TAAG","TAAC","TATA","TATT","TATG","TATC","TAGA","TAGT","TAGG","TAGC","TACA","TACT","TACG","TACC","TTAA","TTAT","TTAG","TTAC","TTTA","TTTT","TTTG","TTTC","TTGA","TTGT","TTGG","TTGC","TTCA","TTCT","TTCG","TTCC","TGAA","TGAT","TGAG","TGAC","TGTA","TGTT","TGTG","TGTC","TGGA","TGGT","TGGG","TGGC","TGCA","TGCT","TGCG","TGCC","TCAA","TCAT","TCAG","TCAC","TCTA","TCTT","TCTG","TCTC","TCGA","TCGT","TCGG","TCGC","TCCA","TCCT","TCCG","TCCC","GAAA","GAAT","GAAG","GAAC","GATA","GATT","GATG","GATC","GAGA","GAGT","GAGG","GAGC","GACA","GACT","GACG","GACC","GTAA","GTAT","GTAG","GTAC","GTTA","GTTT","GTTG","GTTC","GTGA","GTGT","GTGG","GTGC","GTCA","GTCT","GTCG","GTCC","GGAA","GGAT","GGAG","GGAC","GGTA","GGTT","GGTG","GGTC","GGGA","GGGT","GGGG","GGGC","GGCA","GGCT","GGCG","GGCC","GCAA","GCAT","GCAG","GCAC","GCTA","GCTT","GCTG","GCTC","GCGA","GCGT","GCGG","GCGC","GCCA","GCCT","GCCG","GCCC","CAAA","CAAT","CAAG","CAAC","CATA","CATT","CATG","CATC","CAGA","CAGT","CAGG","CAGC","CACA","CACT","CACG","CACC","CTAA","CTAT","CTAG","CTAC","CTTA","CTTT","CTTG","CTTC","CTGA","CTGT","CTGG","CTGC","CTCA","CTCT","CTCG","CTCC","CGAA","CGAT","CGAG","CGAC","CGTA","CGTT","CGTG","CGTC","CGGA","CGGT","CGGG","CGGC","CGCA","CGCT","CGCG","CGCC","CCAA","CCAT","CCAG","CCAC","CCTA","CCTT","CCTG","CCTC","CCGA","CCGT","CCGG","CCGC","CCCA","CCCT","CCCG","CCCC"}

	for _,quadBase := range combinations{
		// quadBase := "ATGC"
		// fmt.Printf("%.8b ",*encode(&quadBase))
		// fmt.Printf("%s",&quadBase)
		fmt.Print("byte(",*encode(&quadBase),") : \"",quadBase+"\",")

	}

	/*codons := "ATGC"
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
	*/
	// fmt.Printf("%d",bytes)
}

func main() {
	basics()
	// base()
	// simple()
	// k := byte(17)
	// x := k >> 2
	// fmt.Printf("%.4b",x)
}