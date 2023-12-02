from math import prod

lines = [line.strip() for line in open("d2/test.txt").readlines()]
games = []
for line in lines:
    left, right = line.split(": ")
    id_ = int(left.removeprefix("Game "))
    sets_str = right.split("; ")
    sets = []
    for set in sets_str:
        colors = set.split(", ")
        cols = []
        for col in colors:
            val, color = col.split(" ")
            cols.append({"color": color, "value": int(val)})
        sets.append(cols)
    games.append({"id": id_, "sets": sets})

total = 0
for game in games:
    maxes = {"red":0, "blue":0, "green":0}
    for set in game["sets"]:
        for col in set:
            if maxes[col["color"]] < col["value"]:
                maxes[col["color"]] = col["value"]
    total += prod(maxes.values())
print(total)
