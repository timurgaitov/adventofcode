import re

w = 101
h = 103

def pos_after(nums, after):
  px = nums[0]
  py = nums[1]
  vx = nums[2]
  vy = nums[3]
  for i in range(after):
    px += vx
    py += vy
  return px % w, py % h

count_q1 = 0
count_q2 = 0
count_q3 = 0
count_q4 = 0

with open("input.txt") as file:
  lines = file.readlines()
  for line in lines:
    nums = [int(n) for n in re.findall("-?\d+", line)]
    
    after_x, after_y = pos_after(nums, 100)
    
    if after_x == w//2 or after_y == h//2:
      continue

    if after_x < w//2 and after_y < h//2:
      count_q1 += 1

    if after_x < w//2 and after_y > h//2:
      count_q2 += 1

    if after_x > w//2 and after_y < h//2:
      count_q3 += 1

    if after_x > w//2 and after_y > h//2:
      count_q4 += 1

print(count_q1*count_q2*count_q3*count_q4)
