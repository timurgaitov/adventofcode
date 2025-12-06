import sys

with open(sys.argv[1], 'r') as file:
	lines = file.readlines()
	x = 50
	cnt = 0
	for line in lines:
		direct = line[0]
		assert direct in ('L', 'R')
		num = int(line[1:])
		rot = num // 100
		num %= 100
		before = x
		if direct == 'L':
			x = (x - num + 100) % 100
			if x == 0 or (before != 0 and x > before):
				rot += 1
		elif direct == 'R':
			x = (x + num) % 100
			if x < before:
				rot += 1
		cnt += rot
	print(cnt)	

