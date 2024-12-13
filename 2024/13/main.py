import numpy as np

with open("input.txt") as file:
  lines = file.readlines()
  i = 0
  
  res = 0
  while i < len(lines):
    a = [int(x.split("+")[1]) for x in lines[i].split(": ")[1].split(", ")]
    i += 1
    b = [int(x.split("+")[1]) for x in lines[i].split(": ")[1].split(", ")]
    i += 1
    p = [int(x.split("=")[1]) for x in lines[i].split(": ")[1].split(", ")]
    i += 1
    i += 1
    A = np.array([[a[0], b[0]], [a[1], b[1]]])
    B = np.array([p[0], p[1]])
    x = np.linalg.solve(A, B)
    x[0] = np.round(x[0], 1)
    x[1] = np.round(x[1], 1)
    if x[0] < 0 or x[1] < 0:
      continue
    if x[0] % 1 > 0 or x[1] % 1 > 0:
      continue
    if x[0] > 100 or x[1] > 100:
      continue
    res += 3*x[0]+x[1]
    
  print(res)