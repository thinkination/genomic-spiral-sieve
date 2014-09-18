import itertools 

def getSequence():
	content = None
	with open('seq.txt') as f:
		content = f.readlines()
	return content
    # print content

def run():
	seq = getSequence()
	compressed = encode(seq)
	print compressed

def sets():
	sets = itertools.product('ATGC', repeat=4)
	for seq in sets:
		print seq

#run()
sets()