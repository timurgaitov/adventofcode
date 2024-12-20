import itertools
from collections import defaultdict
from heapq import heappush, heappop
from math import inf

with open('input.txt') as file:
  Map = [list(x) for x in file.read().splitlines()]

size = len(Map)
for i, j in itertools.product(range(size), repeat=2):
  if Map[i][j] == 'E':
    E = (i, j)
  elif Map[i][j] == 'S':
    S = (i, j)

dirs = [(0, -1), (0, 1), (-1, 0), (1, 0)]


def move(pos, d):
  return pos[0] + d[0], pos[1] + d[1]


def find_path(_map, _start):
  _dist = defaultdict(lambda: inf)

  pq = []
  heappush(pq, (0, _start))

  while len(pq) > 0:
    cost, cur = heappop(pq)

    if _map[cur[0]][cur[1]] == '#':
      continue

    if cost >= _dist[cur]:
      continue
    _dist[cur] = cost

    for d in dirs:
      heappush(pq, (_dist[cur] + 1, move(cur, d)))

  return _dist


dist_e = find_path(Map, E)
dist_s = find_path(Map, S)

save = 100
cheats = 20
ref_val = dist_s[E]

count = set()

for i, j, k, m in itertools.product(range(0, size), repeat=4):
  if Map[i][j] == '#' or Map[k][m] == '#':
    continue

  d = abs(i - k) + abs(j - m)

  if d == 0:
    continue

  if d > cheats:
    continue

  if dist_s[(i, j)] + dist_e[(k, m)] + d > ref_val - save:
    continue

  count.add(((i, j), (k, m)))

print(len(count))
