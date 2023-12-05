text = open("d5/in.txt").read()
seed_block, *blocks = text.split("\n\n")
seeds = [int(s) for s in seed_block[7:].split()]
maps = [[[int(d) for d in l.split()] for l in b.strip().split("\n")[1:]] for b in blocks]

def F(seed: int) -> int:
    cur = seed
    for map_ in maps:
        for dest,source,rang in map_:
            if source <= cur < source + rang:
                cur = dest+cur-source
                break
    return cur

result = []
for seed in seeds:
    result.append(F(seed))
print(f"part 1: {min(result)}")

result = []
for start, count in zip(seeds[::2], seeds[1::2]):
    print(f"{count} interval...")
    for seed in range(start, start+count):
        result.append(F(seed)) 
    print(f"{count} done")
#part 2: 2520479
#pypy d5/main.py  3232.78s user 50.10s system 99% cpu 54:43.54 total
print(f"part 2: {min(result)}")
