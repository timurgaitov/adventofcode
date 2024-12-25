locks = []
keys = []

with open('input.txt') as file:
  items = file.read().split('\n\n')
  for item in items:
    lines = item.splitlines()
    if all(x == '#' for x in lines[0]):
      lock = [0] * len(lines[0])
      for j in range(len(lines[0])):
        for i in range(1, len(lines)):
          if lines[i][j] == '#':
            lock[j] += 1
      locks.append(lock)
    else:
      key = [0] * len(lines[0])
      for j in range(len(lines[0])):
        for i in range(0, len(lines) - 1):
          if lines[i][j] == '#':
            key[j] += 1
      keys.append(key)

count = 0
for lock in locks:
  for key in keys:
    if all(x <= 5 for x in map(sum, zip(lock, key))):
      count += 1

print(count)
