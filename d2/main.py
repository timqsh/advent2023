from math import prod
import re

lines = [line.strip() for line in open("d2/test.txt").readlines()]
total = 0
for line in lines:
    maxes = {"red":0, "blue":0, "green":0}
    sets = re.split(", |; ", line.split(": ")[1])
    for s in sets:
        val, color = s.split(" ")
        maxes[color] = max(maxes[color], int(val))
    total += prod(maxes.values())
print(total)
