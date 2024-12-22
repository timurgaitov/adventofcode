from collections import deque, defaultdict

with open('input.txt') as file:
  nums = [int(x) for x in file.read().splitlines()]


def secret(init):
  cur = init
  while True:
    cur = ((cur * 64) ^ cur) % 16777216
    cur = ((cur // 32) ^ cur) % 16777216
    cur = ((cur * 2048) ^ cur) % 16777216
    yield cur


def seq_price(init, num_secrets):
  i = 0
  prev = init % 10
  q = deque()

  res = defaultdict(lambda: 0)

  for s in secret(init):
    cur = s % 10
    change = cur - prev
    if len(q) == 4:
      q.popleft()
    q.append(change)

    if len(q) == 4 and tuple(q) not in res:
      res[tuple(q)] = cur

    prev = cur
    i += 1
    if i >= num_secrets:
      break

  return res


seq_prices = [seq_price(num, 2000) for num in nums]

sum = defaultdict(lambda: 0)
for sp in seq_prices:
  for k, v in sp.items():
    sum[k] += v

print(max(sum.items(), key=lambda x: x[1])[1])
