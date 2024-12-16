import copy
import sys
from collections import defaultdict

sys.setrecursionlimit(1000000000)

# 280964

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


def dfs(cur, visited, prev_dir, mem):
  if (cur, prev_dir) in mem:
    return mem[(cur, prev_dir)]

  if cur in visited:
    return 1000000000
  visited.add(cur)

  if maze[cur[0]][cur[1]] == 'E':
    return 0

  if maze[cur[0]][cur[1]] == '#':
    return 1000000000

  min_score = 1000000000

  for dir in dirs:
    if dir[0] + prev_dir[0] == 0 and dir[1] + prev_dir[1] == 0:
      continue

    change = 1
    if dir != prev_dir:
      change += 1000

    score = change + dfs((cur[0] + dir[0], cur[1] + dir[1]), copy.copy(visited),
                         dir, mem)
    if score < min_score:
      min_score = score

  mem[(cur, prev_dir)] = min_score
  return min_score


print(dfs(S, set(), dirs[0], defaultdict()))
