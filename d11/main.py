def main():
    EXPANSION_RATE = 1_000_000
    lines = open("d11/in.txt").read().splitlines()
    empty_rows = {row_idx for row_idx, row in enumerate(lines) if all(c=="." for c in row)}
    empty_cols = {col_idx for col_idx, col in enumerate(zip(*lines)) if all(c=="." for c in col)}
    galaxies: list[tuple[int,int]] = []
    for row_idx, row in enumerate(lines):
        for col_idx, char in enumerate(row):
            if char == "#":
                galaxies.append((col_idx, row_idx))

    dists = {}
    for i in range(len(galaxies)):
        for j in range(i+1, len(galaxies)):
            g1x, g1y = galaxies[i]
            g2x, g2y = galaxies[j]

            diap_x = range(g1x, g2x) if g2x >= g1x else range(g2x, g1x)
            diap_y = range(g1y, g2y) if g2y >= g1y else range(g2y, g1y)

            dist = 0
            for dx in diap_x:
                if dx in empty_cols:
                    dist += EXPANSION_RATE
                else:
                    dist += 1
            for dy in diap_y:
                if dy in empty_rows:
                    dist += EXPANSION_RATE
                else:
                    dist += 1
            dists[(i+1,j+1)] = dist
    print(sum(dists.values()))



if __name__ == "__main__":
    main()
