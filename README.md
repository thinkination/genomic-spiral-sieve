Welcome to the genomic-spiral-sieve wiki!

#Motivation

![Genomic Spiral Sieve](http://4.bp.blogspot.com/-RW9YrT1B10M/VCEEzAJAFLI/AAAAAAAARKY/QS0XsPCCZYY/s1600/human%2Bdna%2Bsequence.jpg)

Genomic medicine is the eventual future of intelligent healing. It could find cures for diseases which our current generation of humans thought was not curable.

#Introduction

Living organisms tend to have a DNA.

A DNA is made of a long running chain of [nucleotides](http://en.wikipedia.org/wiki/Nucleotide). Often this is addressed as [DNA Sequence](http://en.wikipedia.org/wiki/Nucleic_acid_sequence).

The following image is a textual representation of how a DNA sequence looks like:

![Genome Sequence](http://1.bp.blogspot.com/-u3eFlcCjKEk/VAnEQC7kb5I/AAAAAAAARJI/P50usFLeSoQ/s1600/DNA_sequences.gif)

#Problem

These DNA sequences are really long. Comparing two DNAs or finding a sub-sequence of disease causing sequence in a person's DNA takes too much time.

#Solution

In Information theory, there is a concept of compression.

There are two basic types of information compression:

1. [Loss-less compression](https://en.wikipedia.org/wiki/Lossless_compression)

1. [Lossy Compression](http://en.wikipedia.org/wiki/Lossy_compression)

The well known ZIP file is just application of Loss-Less compression algorithm on a given file to reduce its size on file system without losing any data.

This is the first step of Genomic Spiral Sieve is to compress the DNA sequences without any loss.

# What is Genomic Spiral Seive

Genome - is a sequence of nucleotides - A T G C.
Spiral Sieves - are mathematical systems that give out numbers of a specific pattern.

Genomic Spiral Seive is a system which provides a compressed analysis mechanism for quick operations on a large DNA sequences using information compression techniques.

The ASCII character code for ATGC are

> A 1000001

> T 1010100

> G 1000111

> C 1000011


# Compression Algorithm

A simple hash compression :

> A 1000001 after compression 00

> T 1010100 after compression 01 

> G 1000111 after compression 10

> C 1000011 after compression 11


So a simple ATGC sequence becomes

> ATGC after compression looks like **00011011**
> To decompress - **00** **01** **10** **11** = A T G C 

# What's the value?

## Lesser Space

Instead of using 16 bits to represent a nucleotide GSS uses just 2 bits.

That's 14 bits saved for a single character. One typical DNA might have close to 6 billion of them.

Now imagine the space saved for a whole DNA database!

## Lesser Time

A CPU typically compares two given character A & A by comparing all the 8 bits used in ASCII character code. 

After GSS compression to compare two characters, CPU takes just 2 bits to compare. 

Hence GSS saves atleast 6 CPU instructions per nucleotide comparision.

I would leave the reader to compute the total instructions which could be saved if GSS based platform is used for DNA searches.

# Tip of the ice-berg

GSS compression algorithm is just the base. There is a long way to go.

Saving 6 bits is good. Saving 6 instructions is good. This is just the tip of the ice-berg.

Could we create a Unix **"grep"** like super fast string search algorithms leveraging the high-concurrency abilities of languages like golang? 

Would be be able to check a given DNA and check for all disease causing gnome sequences in seconds? That's GSS ultimate goal!
