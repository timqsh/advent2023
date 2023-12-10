Pos = tuple[int, int]


def find_start(lines: list[str]) -> tuple[int, int]:
    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            if char == "S":
                return (x, y)
    raise ValueError("No 'S' in map")

DIRECTIONS: dict[str, list[Pos]] = {
    "-": [(1, 0), (-1, 0)],
    "|": [(0, 1), (0, -1)],
    "L": [(0, -1), (1, 0)],
    "J": [(0, -1), (-1, 0)],
    "7": [(0, 1), (-1, 0)],
    "F": [(0, 1), (1, 0)],
    ".": [],
}


def is_neighbor(lines: list[str], source: Pos, dest: Pos) -> bool:
    sx, sy = source
    in_map = 0 <= sy < len(lines) and 0 <= sx < len(lines[0])
    if not in_map:
        return False
    char = lines[sy][sx]
    for dx, dy in DIRECTIONS[char]:
        if (sx + dx, sy + dy) == dest:
            return True
    return False


def find_next(lines: list[str], source: Pos, prev: Pos) -> Pos:
    sx, sy = source
    char = lines[sy][sx]
    for dx, dy in DIRECTIONS[char]:
        candidate = (sx + dx, sy + dy)
        if candidate != prev:
            return candidate
    raise ValueError(f"No path from {source}")


def part1() -> None:
    lines = open("d10/in.txt").read().split()
    sx, sy = find_start(lines)
    for dx, dy in ((1, 0), (0, 1), (-1, 0), (0, -1)):
        if is_neighbor(lines, (sx + dx, sy + dy), (sx, sy)):
            cur = (sx + dx, sy + dy)
            break
    else:
        raise ValueError("no path from start")
    prev = (sx, sy)
    steps = 1
    while lines[cur[1]][cur[0]] != "S":
        tmp = find_next(lines, cur, prev)
        prev = cur
        cur = tmp
        steps += 1
    print(steps // 2)


def part2() -> None:
    lines = open("d10/in.txt").read().split()
    sx, sy = find_start(lines)
    start_neighbors = []
    for dx, dy in ((1, 0), (0, 1), (-1, 0), (0, -1)):
        if is_neighbor(lines, (sx + dx, sy + dy), (sx, sy)):
            cur = (sx + dx, sy + dy)
            start_neighbors.append(cur)
    for char, dirs in DIRECTIONS.items():
        if {(sx + dx, sy + dy) for dx, dy in dirs} == set(start_neighbors):
            start_char = char
            break
    else:
        raise ValueError("Can't find start char")
    lines[sy] = lines[sy].replace("S", start_char)

    prev = (sx, sy)
    dx, dy = DIRECTIONS[start_char][0]
    cur = (sx + dx, sy + dy)
    loop_parts: set[Pos] = {prev, cur}
    while cur != (sx, sy):
        tmp = find_next(lines, cur, prev)
        prev = cur
        cur = tmp
        loop_parts.add(cur)

    total = 0
    in_pipe_from_above: bool | None = None
    in_area = False
    for y, line in enumerate(lines):
        for x, char in enumerate(line):
            if in_pipe_from_above is None:  # not in pipe
                if (x, y) not in loop_parts:
                    total += in_area
                elif char == "|":
                    in_area = not in_area
                elif char == "L":
                    in_pipe_from_above = True
                elif char == "F":
                    in_pipe_from_above = False
                else:
                    raise RuntimeError(
                        f"unexpected {char=}, {in_area=}, {in_pipe_from_above=}, {x=}, {y=}"
                    )
            else:  # in pipe
                if (x, y) not in loop_parts:
                    raise RuntimeError(
                        f"unexpected not in loop {char=}, {in_area=}, {in_pipe_from_above=}, {x=}, {y=}"
                    )
                if char == "-":
                    continue
                elif char == "J":
                    if in_pipe_from_above is False:
                        in_area = not in_area
                    in_pipe_from_above = None
                elif char == "7":
                    if in_pipe_from_above is True:
                        in_area = not in_area
                    in_pipe_from_above = None
                else:
                    raise RuntimeError(
                        f"unexpected {char=}, {in_area=}, {in_pipe_from_above=}, {x=}, {y=}"
                    )
    print(total)


if __name__ == "__main__":
    part1()
    part2()
