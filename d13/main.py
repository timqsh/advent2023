def block_value(rows: list[str], target_fails=0) -> int:
    cols = list(zip(*rows))
    for lines, coefficient in [(rows, 100), (cols, 1)]:
        for mid in range(1, len(lines)):
            fails = 0
            offset = 0
            while (lo := mid - offset - 1) >= 0 and (hi := mid + offset) < len(lines):
                diff = sum(lines[lo][i] != lines[hi][i] for i in range(len(lines[0])))
                fails += diff
                if fails > target_fails:
                    break
                offset += 1
            if fails == target_fails:
                return coefficient * mid
    raise ValueError(f"Can't find symmetry in block:\n{rows}\n")


def main() -> None:
    text = open("d13/in.txt").read()
    blocks = [bb.splitlines() for bb in text.split("\n\n")]
    print("part 1:", sum(block_value(b) for b in blocks))
    print("part 2:", sum(block_value(b, target_fails=1) for b in blocks))


if __name__ == "__main__":
    main()
