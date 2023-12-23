def part1() -> None:
    lines = open("d23/in.txt").read().splitlines()

    start_x = lines[0].index(".")
    start_y = 0
    seen = set(((start_x, start_y),))
    path_lens = set()

    def get_neighbors(x, y):
        if lines[y][x] == ".":
            directions = ((0, 1), (0, -1), (1, 0), (-1, 0))
        elif lines[y][x] == ">":
            directions = ((1, 0),)
        elif lines[y][x] == "<":
            directions = ((-1, 0),)
        elif lines[y][x] == "^":
            directions = ((0, -1),)
        elif lines[y][x] == "v":
            directions = ((0, 1),)
        else:
            raise ValueError(f"Unknown character {lines[y][x]}")

        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            if (
                0 <= nx < len(lines[0])
                and 0 <= ny < len(lines)
                and lines[ny][nx] != "#"
                and (nx, ny) not in seen
            ):
                yield nx, ny

    def backtrack(x, y, steps):
        to_remove = set()
        while True:
            if y == len(lines) - 1:
                path_lens.add(steps)
                seen.difference_update(to_remove)
                return
            neighbors = list(get_neighbors(x, y))
            if len(neighbors) == 0:
                seen.difference_update(to_remove)
                return
            elif len(neighbors) == 1:
                x, y = neighbors[0]
                seen.add((x, y))
                to_remove.add((x, y))
                steps += 1
            else:
                for nx, ny in neighbors:
                    seen.add((nx, ny))
                    backtrack(nx, ny, steps + 1)
                    seen.remove((nx, ny))
                seen.difference_update(to_remove)
                return

    backtrack(start_x, start_y, 0)
    print("part 1:", max(path_lens))


def part2() -> None:
    lines = open("d23/in.txt").read().splitlines()

    nodes = {
        (x, y) for y, line in enumerate(lines) for x, c in enumerate(line) if c != "#"
    }
    edges = {
        (x, y): [
            ((x + dx, y + dy), 1)
            for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]
            if (x + dx, y + dy) in nodes
        ]
        for x, y in nodes
    }

    for node in nodes.copy():
        if len(edges[node]) == 0:
            del edges[node]
            nodes.remove(node)
        if len(edges[node]) == 2:
            (n1, d1), (n2, d2) = edges[node]
            edges[n1].remove((node, d1))
            edges[n2].remove((node, d2))
            edges[n1].append((n2, d1 + d2))
            edges[n2].append((n1, d1 + d2))
            del edges[node]
            nodes.remove(node)
    
    def backtrack(node, steps):
        if node[1] == len(lines) - 1:
            path_lens.append(steps)
            return
        for neighbor, dist in edges[node]:
            if neighbor not in seen:
                seen.add(neighbor)
                backtrack(neighbor, steps + dist)
                seen.remove(neighbor)

    start = (lines[0].index("."), 0)
    seen = {start}
    path_lens = []
    backtrack(start, 0)
    print("part 2:", max(path_lens))
 


part1()
part2()
