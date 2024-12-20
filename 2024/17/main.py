with open('input.txt') as file:
  lines = file.readlines()

A = int(lines[0].split(': ')[1])
B = int(lines[1].split(': ')[1])
C = int(lines[2].split(': ')[1])
prog = [int(x) for x in lines[4].split(': ')[1].split(',')]


def combo(op, proc):
  assert 0 <= op < 7
  if op < 4:
    return op
  if op == 4:
    return proc[0]
  if op == 5:
    return proc[1]
  if op == 6:
    return proc[2]


def adv(op, proc):
  proc[0] >>= combo(op, proc)
  proc[3] += 2


def bxl(op, proc):
  proc[1] ^= op
  proc[3] += 2


def bst(op, proc):
  proc[1] = combo(op, proc) & 0b111
  proc[3] += 2


def jnz(op, proc):
  if proc[0] == 0:
    proc[3] += 2
    return
  proc[3] = op


def bxc(_, proc):
  proc[1] ^= proc[2]
  proc[3] += 2


def out(op, proc):
  proc[4].append(combo(op, proc) & 0b111)
  proc[3] += 2


def bdv(op, proc):
  proc[1] = proc[0] >> combo(op, proc)
  proc[3] += 2


def cdv(op, proc):
  proc[2] = proc[0] >> combo(op, proc)
  proc[3] += 2


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

proc = [A, B, C, 0, []]
while proc[3] < len(prog):
  opcode[prog[proc[3]]](prog[proc[3] + 1], proc)

print(','.join([str(x) for x in proc[4]]))


# part 2
# guessing three bits erased by division
def find_a(backw_i, a):
  if backw_i < 0:
    return a
  for x in range(0b111 + 1):
    cand_a = a << 3 | x

    proc = [cand_a, 0, 0, 0, []]
    for i in range(0, len(prog), 2):
      opcode[prog[i]](prog[i + 1], proc)

    if proc[4][0] == prog[backw_i]:
      result = find_a(backw_i - 1, cand_a)
      if result is None:
        continue
      return result


print(find_a(len(prog) - 1, 0))
