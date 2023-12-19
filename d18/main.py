from dataclasses import dataclass

@dataclass
class Cmd:
    dir: str
    dist: int
    color: str

def main() -> None:
    lines = open("d18/in.txt").read().splitlines()
    cmds1 = []
    cmds2 = []
    for line in lines:
        dir, dist, color = line.split()
        color = color[2:-1]
        assert len(color) == 6
        real_dist = int(color[:-1], base=16)
        real_dir = "RDLU"[int(color[-1])]
        cmds1.append(Cmd(dir, int(dist), color))
        cmds2.append(Cmd(real_dir, real_dist, color))
    part1(cmds1)
    part2(cmds2)

def part2(cmds: list[Cmd]) -> None:
    x, y = 0, 0
    vertices: list[tuple[int, int]] = [(0, 0)]
    for cmd in cmds:
        if cmd.dir == "U":
            y -= cmd.dist
            vertices.append((x, y))
        elif cmd.dir == "D":
            y += cmd.dist
            vertices.append((x, y))
        elif cmd.dir == "L":
            x -= cmd.dist
            vertices.append((x, y))
        elif cmd.dir == "R":
            x += cmd.dist
            vertices.append((x, y))

    # first vertex is repeated in vertices[-1]

    area = 0
    perimeter = 0
    for i in range(len(vertices) - 1):
        cx, cy = vertices[i]
        nx, ny = vertices[i + 1]

        perimeter += abs(cx - nx) + abs(cy - ny)
        area += (nx - cx) * cy  # either dx is 0 or cy == ny
    area = abs(area)

    print(area + perimeter//2 + 1)



def part1(cmds: list[Cmd]) -> None:
    x, y = 0, 0
    edges: set[tuple[int, int]] = {(0, 0)}
    for cmd in cmds:
        if cmd.dir == "U":
            for i in range(cmd.dist):
                y -= 1
                edges.add((x, y))
        elif cmd.dir == "D":
            for i in range(cmd.dist):
                y += 1
                edges.add((x, y))
        elif cmd.dir == "L":
            for i in range(cmd.dist):
                x -= 1
                edges.add((x, y))
        elif cmd.dir == "R":
            for i in range(cmd.dist):
                x += 1
                edges.add((x, y))
    
    max_x = max(edges, key=lambda x: x[0])[0] + 2
    min_x = min(edges, key=lambda x: x[0])[0] - 2
    max_y = max(edges, key=lambda x: x[1])[1] + 2
    min_y = min(edges, key=lambda x: x[1])[1] - 2

    seen: set[tuple[int, int]] = set()
    components: list[set[tuple[int, int]]] = []
    count = 0
    for y in range(min_y, max_y):
        for x in range(min_x, max_x):
            if (x, y) in edges or (x, y) in seen:
                continue
            count += 1
            component: set[tuple[int, int]] = set()
            stack: list[tuple[int, int]] = [(x, y)]
            while stack:
                node = stack.pop()
                seen.add(node)
                component.add(node)
                for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                    new_x, new_y = node[0] + dx, node[1] + dy
                    if not (min_x <= new_x <= max_x and min_y <= new_y <= max_y):
                        continue
                    neighbor = (node[0] + dx, node[1] + dy)
                    if neighbor in edges or neighbor in seen:
                        continue
                    stack.append((node[0] + dx, node[1] + dy))
            components.append(component)

    assert len(components) == 2
    min_component = min(components, key=lambda x: len(x))
    total = len(edges) + len(min_component)
    print(total)
        

main()
