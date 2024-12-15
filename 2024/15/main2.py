with open("input.txt") as file:
  lines = [x.rstrip() for x in file.readlines()]
split = lines.index("")


def enlarge(x):
  return (x.replace("#", "##")
          .replace("O", "[]")
          .replace(".", "..")
          .replace("@", "@."))


Map = [list(enlarge(x)) for x in lines[:split]]
Moves = list(sum([list(x) for x in lines[split + 1:]], []))


def move_pos(pos, arrow):
  if arrow == "<":
    return pos[0], pos[1] - 1
  elif arrow == ">":
    return pos[0], pos[1] + 1
  elif arrow == "^":
    return pos[0] - 1, pos[1]
  elif arrow == "v":
    return pos[0] + 1, pos[1]
  else:
    return None


def can_move(pos, arrow):
  cand_pos = move_pos(pos, arrow)
  cand = Map[cand_pos[0]][cand_pos[1]]

  if cand == ".":
    return True

  if cand == "#":
    return False

  if arrow in ("^", "v"):
    if cand == "[":
      return can_move(cand_pos, arrow) and can_move(move_pos(cand_pos, ">"),
                                                    arrow)
    if cand == "]":
      return can_move(cand_pos, arrow) and can_move(move_pos(cand_pos, "<"),
                                                    arrow)

  return can_move(cand_pos, arrow)


def find_place_to_move(pos, arrow):
  cand_pos = move_pos(pos, arrow)
  cand = Map[cand_pos[0]][cand_pos[1]]

  if cand == ".":
    Map[cand_pos[0]][cand_pos[1]] = Map[pos[0]][pos[1]]
    Map[pos[0]][pos[1]] = "."
    return

  if cand == "#":
    return

  if not can_move(cand_pos, arrow):
    return

  other_part = None

  if cand == "[" and arrow in ("^", "v"):
    other_part = move_pos(cand_pos, ">")
    if not can_move(cand_pos, arrow) or not can_move(other_part, arrow):
      return

  if cand == "]" and arrow in ("^", "v"):
    other_part = move_pos(cand_pos, "<")
    if not can_move(cand_pos, arrow) or not can_move(other_part, arrow):
      return

  find_place_to_move(cand_pos, arrow)
  if other_part is not None:
    find_place_to_move(other_part, arrow)
    if Map[pos[0]][pos[1]] != "@":
      Map[other_part[0]][other_part[1]] = Map[pos[0]][pos[1]]
      Map[other_part[0]][other_part[1]] = "."

  Map[cand_pos[0]][cand_pos[1]] = Map[pos[0]][pos[1]]
  Map[pos[0]][pos[1]] = "."


move_i = 0
for move in Moves:
  a_pos = None
  for i in range(len(Map)):
    for j in range(len(Map[i])):
      if Map[i][j] == "@":
        a_pos = (i, j)
  find_place_to_move(a_pos, move)
  if move_i % 10000 == 0:
    print("move", move_i)
    for i in range(len(Map)):
      print(''.join(Map[i]))
  move_i += 1

coord_sum = 0

for i in range(len(Map)):
  for j in range(len(Map[i])):
    if Map[i][j] == "[":
      coord_sum += 100 * i + j

print(coord_sum)
