from collections import defaultdict
from heapq import heappush, heappop
from math import inf

with open('input.txt') as file:
  blocks = [(int(c[0]), int(c[1])) for c in [str_c.rstrip().split(',') for str_c
                                             in file.readlines()]]

size = 71

space = [[0 for x in range(size)] for y in range(size)]

for b in blocks[:1024]:
  if b[0] < 0 or b[0] >= size or b[1] < 0 or b[1] >= size:
    continue

  space[b[1]][b[0]] = 1

dirs = [(0, 1), (-1, 0), (0, -1), (1, 0)]


def mov(pos, d):
  return pos[0] + d[0], pos[1] + d[1]


def find_path(space_t, start):
  dist = defaultdict(lambda: inf)

  pq = []
  heappush(pq, (0, start))

  while len(pq) > 0:
    cost, cur = heappop(pq)

    if cur[0] < 0 or cur[0] >= size or cur[1] < 0 or cur[1] >= size:
      continue

    if space_t[cur[0]][cur[1]] == 1:
      continue

    if cost >= dist[cur]:
      continue
    dist[cur] = cost

    for d in dirs:
      heappush(pq, (dist[cur] + 1, mov(cur, d)))

  return dist


dist = find_path(space, (0, 0))
print(dist[(size - 1, size - 1)])

# part 2
space2 = [[0 for x in range(size)] for y in range(size)]

for b in blocks:
  if b[0] < 0 or b[0] >= size or b[1] < 0 or b[1] >= size:
    continue
  space2[b[1]][b[0]] = 1
  dist2 = find_path(space2, (0, 0))
  if dist2[(size - 1, size - 1)] == inf:
    print(b)
    break
