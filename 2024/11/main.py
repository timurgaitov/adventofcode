from functools import cache

@cache
def rec(n, b):
  if b == 0:
    return 1
  if n == 0:
    return rec(1, b - 1)
  n_str = str(n)
  if len(n_str) % 2 == 0:
    half = len(n_str) // 2
    return rec(int(n_str[:half]), b - 1) + rec(int(n_str[half:]), b - 1)
  return rec(n * 2024, b - 1)

with open("input.txt") as file:
  nums = [int(num_str) for num_str in file.readline().split()]
  res = sum(rec(num, 75) for num in nums)
  print(res)
