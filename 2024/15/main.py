with open("input.txt") as file:
  lines = [x.rstrip() for x in file.readlines()]
split = lines.index("")
Map = [list(x) for x in lines[:split]]
Moves = list(sum([list(x) for x in lines[split + 1:]], []))


def find_place_to_move(pos, arrow):
  cand = None
  if arrow == "<":
    cand = (pos[0], pos[1] - 1)
  if arrow == ">":
    cand = (pos[0], pos[1] + 1)
  if arrow == "^":
    cand = (pos[0] - 1, pos[1])
  if arrow == "v":
    cand = (pos[0] + 1, pos[1])

  if Map[cand[0]][cand[1]] == ".":
    Map[cand[0]][cand[1]] = Map[pos[0]][pos[1]]
    Map[pos[0]][pos[1]] = "."
    return True

  if Map[cand[0]][cand[1]] == "#":
    return False

  if not find_place_to_move(cand, arrow):
    return False

  Map[cand[0]][cand[1]] = Map[pos[0]][pos[1]]
  Map[pos[0]][pos[1]] = "."
  return True


for move in Moves:
  a_pos = None
  for i in range(len(Map)):
    for j in range(len(Map[i])):
      if Map[i][j] == "@":
        a_pos = (i, j)
  find_place_to_move(a_pos, move)
  # for i in range(len(Map)):
  #   print(''.join(Map[i]))

coord_sum = 0

for i in range(len(Map)):
  for j in range(len(Map[i])):
    if Map[i][j] == "O":
      coord_sum += 100 * i + j

print(coord_sum)
