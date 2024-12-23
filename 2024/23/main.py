import itertools
from collections import defaultdict

with open('input.txt') as file:
  pairs = [tuple(line.split('-')) for line in file.read().splitlines()]

net = defaultdict(lambda: set())

for pair in pairs:
  net[pair[0]].add(pair[1])
  net[pair[1]].add(pair[0])

res = set()

for comp, conns in net.items():
  conns_list = list(conns)
  for i, j in itertools.product(range(len(conns_list)), repeat=2):
    other1 = conns_list[i]
    other2 = conns_list[j]
    three = [comp, other1, other2]

    if not any(t.startswith('t') for t in three):
      continue

    if other1 in net[other2] and other2 in net[other1]:
      three.sort()
      res.add(tuple(three))

print(len(res))

# part 2

max_ic = set()
visited = set()

for comp, conns in net.items():
  ic = set()
  ic.add(comp)

  st = set()
  st |= conns
  while len(st) > 0:
    cur = st.pop()
    if cur in ic:
      continue
    diff = ic.difference(net[cur])
    if len(diff) > 0:
      continue
    ic.add(cur)
    st |= net[cur].difference(ic)

  if len(ic) > len(max_ic):
    max_ic = ic

  visited |= ic

print(len(max_ic))
print(','.join(sorted(max_ic)))
