from collections import defaultdict

def hash_str(s: str) -> int:
    result = 0
    for char in s:
        result += ord(char)
        result *= 17
        result %= 256
    return result

def main():
    instructions = open("d15/in.txt").read().strip().split(",")

    print("part 1:", sum(hash_str(s) for s in instructions))

    data = defaultdict(dict)
    for instruction in instructions:
        sep = "-" if "-" in instruction else "="
        left, right = instruction.split(sep)
        box_id = hash_str(left)
        box = data[box_id]
        if sep == "-":
            if left in box:
                del box[left]
        else:
            focus = int(right)
            box[left] = focus
    result = 0
    for box_id, lenses in data.items():
        for count, focus in enumerate(lenses.values(), start=1):
            result += (1 + box_id) * count * focus
    print("part 2:", result)
            
main()
