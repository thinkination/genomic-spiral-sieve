//Thanks to http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write#18479916

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var bitMap map[byte]string = map[byte]string{byte(0): "AAAA", byte(1): "AAAT", byte(2): "AAAG", byte(3): "AAAC", byte(4): "AATA", byte(5): "AATT", byte(6): "AATG", byte(7): "AATC", byte(8): "AAGA", byte(9): "AAGT", byte(10): "AAGG", byte(11): "AAGC", byte(12): "AACA", byte(13): "AACT", byte(14): "AACG", byte(15): "AACC", byte(16): "ATAA", byte(17): "ATAT", byte(18): "ATAG", byte(19): "ATAC", byte(20): "ATTA", byte(21): "ATTT", byte(22): "ATTG", byte(23): "ATTC", byte(24): "ATGA", byte(25): "ATGT", byte(26): "ATGG", byte(27): "ATGC", byte(28): "ATCA", byte(29): "ATCT", byte(30): "ATCG", byte(31): "ATCC", byte(32): "AGAA", byte(33): "AGAT", byte(34): "AGAG", byte(35): "AGAC", byte(36): "AGTA", byte(37): "AGTT", byte(38): "AGTG", byte(39): "AGTC", byte(40): "AGGA", byte(41): "AGGT", byte(42): "AGGG", byte(43): "AGGC", byte(44): "AGCA", byte(45): "AGCT", byte(46): "AGCG", byte(47): "AGCC", byte(48): "ACAA", byte(49): "ACAT", byte(50): "ACAG", byte(51): "ACAC", byte(52): "ACTA", byte(53): "ACTT", byte(54): "ACTG", byte(55): "ACTC", byte(56): "ACGA", byte(57): "ACGT", byte(58): "ACGG", byte(59): "ACGC", byte(60): "ACCA", byte(61): "ACCT", byte(62): "ACCG", byte(63): "ACCC", byte(64): "TAAA", byte(65): "TAAT", byte(66): "TAAG", byte(67): "TAAC", byte(68): "TATA", byte(69): "TATT", byte(70): "TATG", byte(71): "TATC", byte(72): "TAGA", byte(73): "TAGT", byte(74): "TAGG", byte(75): "TAGC", byte(76): "TACA", byte(77): "TACT", byte(78): "TACG", byte(79): "TACC", byte(80): "TTAA", byte(81): "TTAT", byte(82): "TTAG", byte(83): "TTAC", byte(84): "TTTA", byte(85): "TTTT", byte(86): "TTTG", byte(87): "TTTC", byte(88): "TTGA", byte(89): "TTGT", byte(90): "TTGG", byte(91): "TTGC", byte(92): "TTCA", byte(93): "TTCT", byte(94): "TTCG", byte(95): "TTCC", byte(96): "TGAA", byte(97): "TGAT", byte(98): "TGAG", byte(99): "TGAC", byte(100): "TGTA", byte(101): "TGTT", byte(102): "TGTG", byte(103): "TGTC", byte(104): "TGGA", byte(105): "TGGT", byte(106): "TGGG", byte(107): "TGGC", byte(108): "TGCA", byte(109): "TGCT", byte(110): "TGCG", byte(111): "TGCC", byte(112): "TCAA", byte(113): "TCAT", byte(114): "TCAG", byte(115): "TCAC", byte(116): "TCTA", byte(117): "TCTT", byte(118): "TCTG", byte(119): "TCTC", byte(120): "TCGA", byte(121): "TCGT", byte(122): "TCGG", byte(123): "TCGC", byte(124): "TCCA", byte(125): "TCCT", byte(126): "TCCG", byte(127): "TCCC", byte(128): "GAAA", byte(129): "GAAT", byte(130): "GAAG", byte(131): "GAAC", byte(132): "GATA", byte(133): "GATT", byte(134): "GATG", byte(135): "GATC", byte(136): "GAGA", byte(137): "GAGT", byte(138): "GAGG", byte(139): "GAGC", byte(140): "GACA", byte(141): "GACT", byte(142): "GACG", byte(143): "GACC", byte(144): "GTAA", byte(145): "GTAT", byte(146): "GTAG", byte(147): "GTAC", byte(148): "GTTA", byte(149): "GTTT", byte(150): "GTTG", byte(151): "GTTC", byte(152): "GTGA", byte(153): "GTGT", byte(154): "GTGG", byte(155): "GTGC", byte(156): "GTCA", byte(157): "GTCT", byte(158): "GTCG", byte(159): "GTCC", byte(160): "GGAA", byte(161): "GGAT", byte(162): "GGAG", byte(163): "GGAC", byte(164): "GGTA", byte(165): "GGTT", byte(166): "GGTG", byte(167): "GGTC", byte(168): "GGGA", byte(169): "GGGT", byte(170): "GGGG", byte(171): "GGGC", byte(172): "GGCA", byte(173): "GGCT", byte(174): "GGCG", byte(175): "GGCC", byte(176): "GCAA", byte(177): "GCAT", byte(178): "GCAG", byte(179): "GCAC", byte(180): "GCTA", byte(181): "GCTT", byte(182): "GCTG", byte(183): "GCTC", byte(184): "GCGA", byte(185): "GCGT", byte(186): "GCGG", byte(187): "GCGC", byte(188): "GCCA", byte(189): "GCCT", byte(190): "GCCG", byte(191): "GCCC", byte(192): "CAAA", byte(193): "CAAT", byte(194): "CAAG", byte(195): "CAAC", byte(196): "CATA", byte(197): "CATT", byte(198): "CATG", byte(199): "CATC", byte(200): "CAGA", byte(201): "CAGT", byte(202): "CAGG", byte(203): "CAGC", byte(204): "CACA", byte(205): "CACT", byte(206): "CACG", byte(207): "CACC", byte(208): "CTAA", byte(209): "CTAT", byte(210): "CTAG", byte(211): "CTAC", byte(212): "CTTA", byte(213): "CTTT", byte(214): "CTTG", byte(215): "CTTC", byte(216): "CTGA", byte(217): "CTGT", byte(218): "CTGG", byte(219): "CTGC", byte(220): "CTCA", byte(221): "CTCT", byte(222): "CTCG", byte(223): "CTCC", byte(224): "CGAA", byte(225): "CGAT", byte(226): "CGAG", byte(227): "CGAC", byte(228): "CGTA", byte(229): "CGTT", byte(230): "CGTG", byte(231): "CGTC", byte(232): "CGGA", byte(233): "CGGT", byte(234): "CGGG", byte(235): "CGGC", byte(236): "CGCA", byte(237): "CGCT", byte(238): "CGCG", byte(239): "CGCC", byte(240): "CCAA", byte(241): "CCAT", byte(242): "CCAG", byte(243): "CCAC", byte(244): "CCTA", byte(245): "CCTT", byte(246): "CCTG", byte(247): "CCTC", byte(248): "CCGA", byte(249): "CCGT", byte(250): "CCGG", byte(251): "CCGC", byte(252): "CCCA", byte(253): "CCCT", byte(254): "CCCG", byte(255): "CCCC"}

func decrypt(bits *byte) *string {
	var s string
	if bitMap[*bits] != "" {
		s = (bitMap[*bits])
	}
	return &s
}

//TODO fix the decommression piece
func decompress(bytes *[]byte, n int) *[]string {
	byts := *bytes
	// fmt.Println(byts)
	quads := make([]string, n)
	for i, byt := range byts {
		if bitMap[byt] != "" {
			quads[i] = bitMap[byt]
		}
	}

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
		fmt.Fprint(w, line)
	}
	return w.Flush()
}

// A 1000001 after compression 00
// T 1010100 after compression 01
// G 1000111 after compression 10
// C 1000011 after compression 11

//ATGC after compression looks like 00011011
//To decompress - 00 01 10 11 = A T G C

func compress(seq *string) *[]byte {
	var master []byte
	codons := *seq

	var base byte = byte(0)
	// var bits byte
	var bitCounter uint8 = uint8(0)
	for i := 0; i < len(codons); i++ {
		code := codons[i]
		switch code {
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

		bitCounter += 2
		if bitCounter >= 8 {
			master = append(master, base)
			bitCounter = uint8(0)
			base = byte(0)
		}

	}

	return &master
}

func printBits(b byte) {

	fmt.Printf("%.8b", b)
	fmt.Println()
}

func main() {
	var master []byte
	lines, err := readLines("../seq.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	} else {
		for _, line := range lines {
			master = *(compress(&line))			
		}
	}

	err = ioutil.WriteFile("./encoded.txt", master, 0644)
	if err != nil {
		panic(err)
	}
	 
	 dat, _ := ioutil.ReadFile("./encoded.txt")
	 var decryptedText []string
	 for _,byt := range dat{
	 	decryptedText = append(decryptedText,bitMap[byt])
	 	// fmt.Print(bitMap[byt])
	 }
	// fmt.Print(decryptedText)
	// fmt.Println(len(decryptedText))

	writeLines(decryptedText,"./decoded.txt")
}
