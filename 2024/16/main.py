from collections import defaultdict
from heapq import heappush, heappop
from math import inf

with open('input.txt') as file:
  maze = [list(x) for x in [x.rstrip() for x in file.readlines()]]

for i in range(len(maze)):
  for j in range(len(maze[i])):
    if maze[i][j] == 'E':
      E = (i, j)
      continue
    if maze[i][j] == 'S':
      S = (i, j)
      continue

dirs = [(0, 1), (-1, 0), (0, -1), (1, 0)]


def sum_tup(pos, d):
  return pos[0] + d[0], pos[1] + d[1]


def find_path(start, start_dir):
  dist = defaultdict(lambda: inf)
  dist_dir = dict()

  pq = []
  heappush(pq, (0, start, start_dir))

  while len(pq) > 0:
    cost, cur, cur_dir = heappop(pq)

    if maze[cur[0]][cur[1]] == '#':
      continue

    if cost >= dist[cur]:
      continue
    dist[cur] = cost
    dist_dir[cur] = cur_dir

    for d in dirs:
      if sum_tup(cur_dir, d) == (0, 0):
        continue

      if d != cur_dir:
        heappush(pq, (dist[cur] + 1001, sum_tup(cur, d), d))
        continue

      heappush(pq, (dist[cur] + 1, sum_tup(cur, d), d))

  return dist, dist_dir


distS, dist_dirS = find_path(S, dirs[0])
print(distS[E])

# part 2

for d in dirs:
  if sum_tup(d, dist_dirS[E]) == (0, 0):
    oppos_dirSE = d
    break

distE, dist_dirE = find_path(E, oppos_dirSE)

count = 0
for i in range(len(maze)):
  for j in range(len(maze[i])):
    s = distS[(i, j)] + distE[(i, j)]
    if s not in (distS[E], distS[E] - 1000):
      # print(maze[i][j], end='')
      continue
    # print('o', end='')
    count += 1
  # print()
print(count)
