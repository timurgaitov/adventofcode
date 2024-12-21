with open('example.txt') as file:
  lines = file.read().splitlines()

num_coords = {
  '7': (0, 0),
  '8': (0, 1),
  '9': (0, 2),
  '4': (1, 0),
  '5': (1, 1),
  '6': (1, 2),
  '1': (2, 0),
  '2': (2, 1),
  '3': (2, 2),
  '0': (3, 1),
  'A': (3, 2),
  '(x_x)': (3, 0),
}

arrow_coords = {
  '^': (0, 1),
  'A': (0, 2),
  '<': (1, 0),
  'v': (1, 1),
  '>': (1, 2),
  '(x_x)': (0, 0),
}


def mov(cur, target, coords):
  if cur == target:
    return ''

  cc = coords[cur]
  nc = coords[target]

  di = cc[0] - nc[0]
  dj = cc[1] - nc[1]

  if di < 0:
    ver = 'v' * abs(di)
  elif di > 0:
    ver = '^' * abs(di)
  else:
    ver = ''

  if dj < 0:
    hor = '>' * abs(dj)
  elif dj > 0:
    hor = '<' * abs(dj)
  else:
    hor = ''

  if (cc[0] - di, cc[1]) == coords['(x_x)']:
    moves = hor + ver
  else:
    moves = ver + hor

  return moves


def rec(str, coords, depth):
  if depth == 0:
    return str
  cur = 'A'
  seq = ''
  for c in str:
    seq += mov(cur, c, coords) + 'A'
    cur = c
  return rec(seq, coords, depth - 1)


def solv(line):
  str = rec(rec(line, num_coords, 1), arrow_coords, 2)
  return len(str) * int(line[:-1])


print(sum([solv(line) for line in lines]))
