from functools import cache
from math import inf

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


@cache
def mov(cur, target, use_num):
  coords = num_coords if use_num else arrow_coords

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


@cache
def rec(str, depth):
  if depth == 26:
    return len(str)
  cur = 'A'
  seq_len = 0
  for c in str:
    if c == cur:
      seq_len += 1
      continue

    min_len = inf
    for move in mov(cur, c, depth == 0):
      subseq_len = rec(move + 'A', depth + 1)
      if subseq_len < min_len:
        min_len = subseq_len

    seq_len += min_len
    cur = c
  return seq_len


def solv(line):
  seq_len = rec(line, 0)
  return seq_len * int(line[:-1])


print(sum([solv(line) for line in lines]))
