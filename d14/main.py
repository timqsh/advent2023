def main() -> None:
    text = open("d14/in.txt").read()
    lines = [list(s) for s in text.splitlines()]

    #print(*["".join(l) for l in lines], sep="\n")

    cycles = 1000
    for cycle in range(1, cycles):
        # North
        for row_idx, row in enumerate(lines):
            for col_idx, char in enumerate(row):
                if char != "O":
                    continue
                offset = 1
                while row_idx - offset >= 0 and lines[row_idx - offset][col_idx] == ".":
                    offset += 1
                if offset > 1:
                    lines[row_idx - offset + 1][col_idx] = "O"
                    lines[row_idx][col_idx] = "."
    
        # print()
        # print(*["".join(l) for l in lines], sep="\n")

        # West
        for row_idx, row in enumerate(lines):
            for col_idx, char in enumerate(row):
                if char != "O":
                    continue
                offset = 1
                while col_idx - offset >= 0 and lines[row_idx][col_idx - offset] == ".":
                    offset += 1
                if offset > 1:
                    lines[row_idx][col_idx - offset + 1] = "O"
                    lines[row_idx][col_idx] = "."

        # print()
        # print(*["".join(l) for l in lines], sep="\n")

        # South
        for row_idx, row in reversed(list(enumerate(lines))):
            for col_idx, char in enumerate(row):
                if char != "O":
                    continue
                offset = 1
                while row_idx + offset < len(lines) and lines[row_idx + offset][col_idx] == ".":
                    offset += 1
                if offset > 1:
                    lines[row_idx + offset - 1][col_idx] = "O"
                    lines[row_idx][col_idx] = "."

        # print()
        # print(*["".join(l) for l in lines], sep="\n")

        # East
        for row_idx, row in enumerate(lines):
            for col_idx, char in reversed(list(enumerate(row))):
                if char != "O":
                    continue
                offset = 1
                while col_idx + offset < len(lines[0]) and lines[row_idx][col_idx + offset] == ".":
                    offset += 1
                if offset > 1:
                    lines[row_idx][col_idx + offset - 1] = "O"
                    lines[row_idx][col_idx] = "."

        # print()
        # print(*["".join(l) for l in lines], sep="\n")
    
        total = 0
        for row_idx, row in enumerate(lines):
            for col_idx, char in enumerate(row):
                if char != "O":
                    continue
                total += len(lines) - row_idx
        print(total, cycle)



main()
