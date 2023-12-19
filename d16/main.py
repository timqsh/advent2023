Vec = tuple[int, int]

LEFT = (-1, 0)
RIGHT = (1, 0)
UP = (0, -1)
DOWN = (0, 1)

def vec_add(v1: Vec, v2: Vec) -> Vec:
    x1, y1 = v1
    x2, y2 = v2
    return (x1+x2, y1+y2)

def main():
    lines = open("d16/in.txt").read().splitlines()
    energized = set()
    was = set()

    def char_at(pos: Vec) -> str | None:
        x, y = pos
        if 0 <= y < len(lines) and 0 <= x < len(lines[0]):
            return lines[y][x]
        return None

    def debug():
        for y, row in enumerate(lines):
            for x, char in enumerate(row):
                if (x, y) in energized:
                    print("#", end="")
                else:
                    print(char, end="")
            print()
        print("-------------------------------")

    def follow(pos: Vec, dir: Vec):
        cur = pos
        while True:
            ch = char_at(cur)
            if ch is None:
                break

            energized.add(cur)

            if (cur, dir) in was:
                break
            was.add((cur, dir))

            if ch == "|" and (dir in [LEFT, RIGHT]):
                follow(vec_add(cur, UP), UP)
                follow(vec_add(cur, DOWN), DOWN)
                break
            if ch == "-" and (dir in [UP, DOWN]):
                follow(vec_add(cur, LEFT), LEFT)
                follow(vec_add(cur, RIGHT), RIGHT)
                break

            if ch == "/" and dir == LEFT:
                dir = DOWN
            elif ch == "/" and dir == RIGHT:
                dir = UP            
            elif ch == "\\" and dir == LEFT:
                dir = UP
            elif ch == "\\" and dir == RIGHT:
                dir = DOWN
            elif ch == "/" and dir == UP:
                dir = RIGHT
            elif ch == "/" and dir == DOWN:
                dir = LEFT 
            elif ch == "\\" and dir == UP:
                dir = LEFT
            elif ch == "\\" and dir == DOWN:
                dir = RIGHT

            cur = vec_add(cur, dir) 
    
    follow((0, 0), RIGHT)
    print("part 1:", len(energized))

    cur_max = 0
    for i in range(len(lines[0])):
        was = set()
        energized = set()
        follow((i, 0), DOWN)
        cur = len(energized)
        cur_max = max(cur_max, cur)
    for i in range(len(lines[0])):
        was = set()
        energized = set()
        follow((i, len(lines)-1), UP)
        cur = len(energized)
        cur_max = max(cur_max, cur)
    for i in range(len(lines)):
        was = set()
        energized = set()
        follow((0, i), RIGHT)
        cur = len(energized)
        cur_max = max(cur_max, cur)
    for i in range(len(lines)):
        was = set()
        energized = set()
        follow((len(lines[0]) - 1, i), LEFT)
        cur = len(energized)
        cur_max = max(cur_max, cur)

    print("part 2:", cur_max)

main()
