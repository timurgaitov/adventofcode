from functools import cache

with open('input.txt') as file:
  inputs = file.read().split('\n\n')
  wires = dict([x.split(': ') for x in inputs[0].splitlines()])
  gates = dict([(y[1], y[0].split(' ')) for y in
                [x.split(' -> ') for x in inputs[1].splitlines()]])


@cache
def rec(cur):
  if cur not in gates:
    return int(wires[cur])

  instr = gates[cur]

  if instr[1] == 'AND':
    return rec(instr[0]) & rec(instr[2])

  if instr[1] == 'OR':
    return rec(instr[0]) | rec(instr[2])

  if instr[1] == 'XOR':
    return rec(instr[0]) ^ rec(instr[2])

  assert False


res = 0
cur_bit = 0
while True:
  bit_key = f'z{cur_bit:02}'
  if bit_key not in gates:
    break
  res |= (rec(bit_key) << cur_bit)
  cur_bit += 1

print(res)
