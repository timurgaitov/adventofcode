import itertools
from collections import defaultdict
from heapq import heappush, heappop
from math import inf

with open('example.txt') as file:
  track = [list(x) for x in file.read().splitlines()]

for i, j in itertools.product(range(len(track)), repeat=2):
  if track[i][j] == 'E':
    E = (i, j)
  elif track[i][j] == 'S':
    S = (i, j)

dirs = [(0, 1), (-1, 0), (0, -1), (1, 0)]


def mov(pos, d):
  return pos[0] + d[0], pos[1] + d[1]


def find_path(_map, _start):
  _dist = defaultdict(lambda: defaultdict(lambda: inf))

  pq = []
  heappush(pq, (0, _start, (0, 0)))

  while len(pq) > 0:
    cost, cur, cheat = heappop(pq)

    if (cur[0] == 0 or cur[1] == 0
        or cur[0] == len(_map) - 1 or cur[1] == len(_map) - 1):
      continue

    if _map[cur[0]][cur[1]] == '#':
      if cheat != (0, 0):
        continue
      cheat = cur

    if cost >= _dist[cheat][cur]:
      continue
    _dist[cheat][cur] = cost

    for d in dirs:
      heappush(pq, (_dist[cheat][cur] + 1, mov(cur, d), cheat[:]))

  return _dist


dist = find_path(track, S)

count = 0
for _, v in dist.items():
  if v[E] <= dist[(0, 0)][E] - 2:
    count += 1

print(count)
