with open('input.txt') as file:
  nums = [int(x) for x in file.read().splitlines()]


def secret(init, i):
  cur = init
  for _ in range(i):
    cur = ((cur * 64) ^ cur) % 16777216
    cur = ((cur // 32) ^ cur) % 16777216
    cur = ((cur * 2048) ^ cur) % 16777216
  return cur


print(sum([secret(num, 2000) for num in nums]))
