import itertools github

#Reads the sequences
def getSequence():
	'''
		Reads the Nuecliotide sequence from a file 'seq.txt'
		seq.txt contains human genome sequence from http://www.ncbi.nlm.nih.gov/gene/100189401
	'''
	content = None

	with open('seq.txt') as f:
		content = f.readlines()
	return content
    # print content

def printSequence():
	'''
		Prints the Nuecliotide sequence
	'''
	seq = getSequence()	
	print seq

def sets():
	'''
		Generates all 4 character dot-products for ['A','T','G','C']
	'''

	#Using Python's itertools to generate dot product of the matrix ['A','T','G','C']
	sets = itertools.product('ATGC', repeat=4)

	#Printing the combination of sequences
	for seq in sets:
		print seq

#printSequence()
sets()