import re

w = 101
h = 103

def read():
  robots = []
  with open("input.txt") as file:
    lines = file.readlines()
    for line in lines:
      nums = [int(n) for n in re.findall("-?\d+", line)]
      pj, pi, vj, vi = nums
      robots.append((pj, pi, vj, vi))
  return robots

def pos_after(nums, after_sec):
  pj, pi, vj, vi = nums[0], nums[1], nums[2], nums[3]
  return (pj + vj * after_sec) % w, (pi + vi * after_sec) % h

def safe_fac(robots, after_sec):
  count_q1 = 0
  count_q2 = 0
  count_q3 = 0
  count_q4 = 0
  for robot in robots:
    after_j, after_i = pos_after(robot, after_sec)

    if after_j == w//2 or after_i == h//2:
      continue

    if after_j < w//2 and after_i < h//2:
      count_q1 += 1

    if after_j < w//2 and after_i > h//2:
      count_q2 += 1

    if after_j > w//2 and after_i < h//2:
      count_q3 += 1

    if after_j > w//2 and after_i > h//2:
      count_q4 += 1
  return count_q1 * count_q2 * count_q3 * count_q4

def print_robot_after(robots, after_sec):
  robots_after = []
  for i in range(h):
    robots_after.append([])
    for _ in range(w):
      robots_after[i].append(0)

  for robot in robots:
    aj, ai = pos_after(robot, after_sec)
    robots_after[ai][aj] += 1

  for i in range(len(robots_after)):
    print("".join(["." if x == 0 else str(x) for x in robots_after[i]]))

robots = read()

min_sf = 10000000000

for after in range(10000000):
  sf = safe_fac(robots, after)
  if sf < min_sf:
    min_sf = sf
    print("found new min", after)
    print_robot_after(robots, after)
