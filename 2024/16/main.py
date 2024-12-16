from collections import defaultdict
from sys import setrecursionlimit

setrecursionlimit(1000000000)


def apply_dir(a, b):
  return tuple(map(sum, zip(a, b)))


with open('example.txt') as file:
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


def print_trace(maze, trace):
  for i in range(len(maze)):
    for j in range(len(maze[i])):
      if (i, j) == trace[-1]:
        print('@', end='')
        continue
      if (i, j) in trace:
        print('o', end='')
        continue
      if maze[i][j] == '#':
        print('█', end='')
        continue
      if maze[i][j] in ('E', 'S'):
        print(maze[i][j], end='')
        continue
      print('.', end='')
    print()
  print()
  print()


bad_score = 1_000_000_000


def dir_cost(cur_dir, d):
  if cur_dir == d:
    return 1
  return 1001


def dfs(cur, cur_dir, trace, mem, target):
  if cur in trace:
    return bad_score

  if maze[cur[0]][cur[1]] == '#':
    return bad_score

  if maze[cur[0]][cur[1]] == target:
    return 0

  trace.add(cur)
  score = min(
      [dir_cost(cur_dir, d) + dfs(apply_dir(cur, d), d, trace, mem, target) for
       d in dirs])
  trace.remove(cur)

  return score


print(dfs(S, dirs[0], set(), defaultdict(), 'E'))
print(dfs(E, dirs[3], set(), defaultdict(), 'S'))
