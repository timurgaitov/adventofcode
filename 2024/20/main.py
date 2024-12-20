import itertools
from collections import defaultdict
from heapq import heappush, heappop
from math import inf

with open('input.txt') as file:
  track = [list(x) for x in file.read().splitlines()]

size = len(track)
for i, j in itertools.product(range(size), repeat=2):
  if track[i][j] == 'E':
    E = (i, j)
  elif track[i][j] == 'S':
    S = (i, j)

dirs = [(0, 1), (-1, 0), (0, -1), (1, 0)]


def mov(pos, d):
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
      heappush(pq, (_dist[cur] + 1, mov(cur, d)))

  return _dist


dist_E_to_others = find_path(track, E)
dist_S_to_others = find_path(track, S)

count = 0

for i, j in itertools.product(range(1, size - 1), repeat=2):
  if track[i][j] != '#':
    continue

  left = mov((i, j), dirs[2])
  right = mov((i, j), dirs[0])
  top = mov((i, j), dirs[1])
  bottom = mov((i, j), dirs[3])

  saved = 100

  if dist_E_to_others[left] + dist_S_to_others[right] + 2 <= dist_S_to_others[
    E] - saved:
    count += 1

  if dist_E_to_others[right] + dist_S_to_others[left] + 2 <= dist_S_to_others[
    E] - saved:
    count += 1

  if dist_E_to_others[top] + dist_S_to_others[bottom] + 2 <= dist_S_to_others[
    E] - saved:
    count += 1

  if dist_E_to_others[bottom] + dist_S_to_others[top] + 2 <= dist_S_to_others[
    E] - saved:
    count += 1

print(count)
