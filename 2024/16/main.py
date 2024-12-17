from collections import defaultdict
from heapq import heappush, heappop
from math import inf
from sys import setrecursionlimit

setrecursionlimit(1000000000)

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


def find_path(start, start_dir, target):
  dist = defaultdict(lambda: inf)

  pq = []
  heappush(pq, (0, start, start_dir))

  while len(pq) > 0:
    cost, cur, cur_dir = heappop(pq)

    if maze[cur[0]][cur[1]] == '#':
      continue

    if cost >= dist[cur]:
      continue
    dist[cur] = cost

    for d in dirs:
      if sum_tup(cur_dir, d) == (0, 0):
        continue

      if d != cur_dir:
        heappush(pq, (dist[cur] + 1001, sum_tup(cur, d), d))
        continue

      heappush(pq, (dist[cur] + 1, sum_tup(cur, d), d))

  return dist[target]


print(find_path(S, dirs[0], E))
# print(find_path(E, dirs[3], S))
