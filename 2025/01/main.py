import sys

with open(sys.argv[1], 'r') as file:
	lines = file.readlines()
	x = 50
	cnt = 0
	for line in lines:
		direct = line[0]
		assert direct in ('L', 'R')
		num = int(line[1:])
		if direct == 'L':
			x = (x - num + 100) % 100
		elif direct == 'R':
			x = (x + num) % 100
		if x == 0:
			cnt += 1
	print(cnt)	

