from itertools import product

with open('input.txt') as file:
  inputs = file.read().split('\n\n')
  wires = dict([(y[0], int(y[1])) for y in
                [x.split(': ') for x in inputs[0].splitlines()]])
  gates = dict([(y[1], y[0].split(' ')) for y in
                [x.split(' -> ') for x in inputs[1].splitlines()]])


def rec(cur, d):
  if d > 45:
    return 0

  if cur not in gates:
    return wires[cur]

  gate = gates[cur]

  # if d < 4:
  #   print(' |' * d, cur + ':', gate[1], gate[0], gate[2])

  if gate[1] == 'AND':
    return rec(gate[0], d + 1) & rec(gate[2], d + 1)

  if gate[1] == 'OR':
    return rec(gate[0], d + 1) | rec(gate[2], d + 1)

  if gate[1] == 'XOR':
    return rec(gate[0], d + 1) ^ rec(gate[2], d + 1)

  assert False


res = 0
bit = 0
while True:
  z = f'z{bit:02}'
  if z not in gates:
    break
  res |= (rec(z, 0) << bit)
  bit += 1


# print(res)


# part 2
def same_operands(gate1, gate2):
  return (gate1[0] == gate2[0] and gate1[2] == gate2[2]
          or gate1[0] == gate2[2] and gate1[2] == gate2[0])


def is_valid_bit(bit, invalid):
  valid = True
  z = gates[f'z{bit:02}']

  if z[1] != 'XOR':
    invalid.add(f'z{bit:02}')
    return False

  if z[0] not in gates or z[2] not in gates:
    return False

  t = sorted(((gates[z[0]], z[0]), (gates[z[2]], z[2])), key=lambda x: x[0][1])
  xor_g = t[1]
  or_g = t[0]
  if (xor_g[0][1] != 'XOR' or
      not same_operands(xor_g[0], (f'x{bit:02}', '', f'y{bit:02}'))):
    invalid.add(xor_g[1])
    valid = False
  if or_g[0][1] != 'OR':
    invalid.add(or_g[1])
    valid = False

  if not valid:
    return False

  t = sorted([(gates[or_g[0][0]], or_g[0][0]), (gates[or_g[0][2]], or_g[0][2])],
             key=lambda x: 0 if x[0][0].startswith(('x', 'y')) else 1)
  carry = t[0]
  prev = t[1]
  if (carry[0][1] != 'AND' or
      not same_operands(carry[0], (f'x{bit - 1:02}', '', f'y{bit - 1:02}'))):
    invalid.add(carry[1])
    valid = False
  if prev[0][1] != 'AND' or not same_operands(prev[0], gates[f'z{bit - 1:02}']):
    invalid.add(prev[1])
    valid = False
  return valid


def errors():
  invalid = set()
  for bit in range(2, 45):
    is_valid_bit(bit, invalid)
  return invalid


invalid = list(errors())


def swap(key1, key2):
  gates[key1], gates[key2] = gates[key2], gates[key1]


for c in product(invalid, repeat=8):
  if len(set(c)) != len(c):
    continue

  swap(c[0], c[1])
  swap(c[2], c[3])
  swap(c[4], c[5])
  swap(c[6], c[7])

  if len(errors()) == 0:
    print('found')
    print(','.join(sorted(c)))
    break

  swap(c[0], c[1])
  swap(c[2], c[3])
  swap(c[4], c[5])
  swap(c[6], c[7])
