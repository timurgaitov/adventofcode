with open('input.txt') as file:
  lines = file.readlines()

A = int(lines[0].split(': ')[1])
B = int(lines[1].split(': ')[1])
C = int(lines[2].split(': ')[1])
prog = [int(x) for x in lines[4].split(': ')[1].split(',')]
pos = 0
output = []


def combo(op):
  assert 0 <= op < 7
  if op < 4:
    return op
  if op == 4:
    return A
  if op == 5:
    return B
  if op == 6:
    return C


def adv(op):
  global A, pos
  A //= 2 ** combo(op)
  pos += 2


def bxl(op):
  global B, pos
  B ^= op
  pos += 2


def bst(op):
  global B, pos
  B = combo(op) % 8
  pos += 2


def jnz(op):
  global A, pos
  if A == 0:
    pos += 2
    return
  pos = op


def bxc(_):
  global B, C, pos
  B ^= C
  pos += 2


def out(op):
  global output, pos
  output.append(combo(op) % 8)
  pos += 2


def bdv(op):
  global A, B, pos
  B = A // 2 ** combo(op)
  pos += 2


def cdv(op):
  global A, C, pos
  C = A // 2 ** combo(op)
  pos += 2


opcode = [
  adv,
  bxl,
  bst,
  jnz,
  bxc,
  out,
  bdv,
  cdv,
]

while pos < len(prog):
  opcode[prog[pos]](prog[pos + 1])
  print('', end='')

print(','.join([str(x) for x in output]))
