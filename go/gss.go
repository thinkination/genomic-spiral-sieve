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

// A 1000001
// T 1010100
// G 1000111
// C 1000011

func main() {
	lines, err := readLines("../seq.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		bytes := ([]byte)(line)

		for k := 0; k <len(bytes) ; k++ {
				// fmt.Printf("%.4b ",bytes[k])
				switch(bytes[k]){
					case 65: 
						fmt.Printf("%.2b",0)
					case 84:
						fmt.Printf("%.2b",1)
					case 71:
						fmt.Printf("%.2b",2)
					case 67:
						fmt.Printf("%.2b",3)
				}
				fmt.Println()
				// fmt.Println((string)(line[k]))
		}
		// fmt.Println(i, line)
	}

	if err := writeLines(lines, "foo.out.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
