with open('input.txt') as file:
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


def mov(cur, target, depth):
  if depth == 0:
    coords = num_coords
  else:
    coords = arrow_coords

  moves = set()

  if cur == target:
    return moves

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

  if ver == '' and hor == '':
    return moves

  if ver == '':
    moves.add(hor)
    return moves

  if hor == '':
    moves.add(ver)
    return moves

  if (cc[0] - di, cc[1]) != coords['(x_x)']:
    moves.add(ver + hor)

  if (cc[0], cc[1] - dj) != coords['(x_x)']:
    moves.add(hor + ver)

  return moves


def rec(str, depth):
  if depth == 3:
    return str
  cur = 'A'
  seq = ''
  for c in str:
    moves = sorted(
        [rec(move + 'A', depth + 1) for move in mov(cur, c, depth)],
        key=len)
    if len(moves) == 0:
      seq += 'A'
      continue
    seq += moves[0]
    cur = c
  return seq


def solv(line):
  str = rec(line, 0)
  return len(str) * int(line[:-1])


print(sum([solv(line) for line in lines]))
