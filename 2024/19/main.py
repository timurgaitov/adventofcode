from functools import cache

with open('input.txt') as file:
  lines = file.read().splitlines()

pat = set(lines[0].split(', '))
des = lines[2:]

count = 0


@cache
def rec(d):
  if d == '':
    return True
  for i in range(1, len(d) + 1):
    if d[:i] in pat and rec(d[i:]):
      return True
  return False


for d in des:
  if rec(d):
    count += 1

print(count)
