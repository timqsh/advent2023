from dataclasses import dataclass
import heapq

Vec = tuple[int, int]

def vec_add(v1: Vec, v2: Vec) -> Vec:
    x1, y1 = v1
    x2, y2 = v2
    return (x1+x2, y1+y2)

LEFT = (-1, 0)
RIGHT = (1, 0)
UP = (0, -1)
DOWN = (0, 1)
NEXT_DIR = {
    LEFT: [DOWN, UP],
    RIGHT: [UP, DOWN],
    UP: [LEFT, RIGHT],
    DOWN: [RIGHT, LEFT]
}

@dataclass(frozen=True, order=True)
class State:
    pos: Vec
    dir: Vec
    steps: int

class Heap:
    def __init__(self, data=None):
        if data is None:
            data = []
        heapq.heapify(data)
        self.data = data
    
    def push(self, elem):
        heapq.heappush(self.data, elem)
    
    def pop(self):
        return heapq.heappop(self.data)


def main() -> None:
    lines = open("d17/in.txt").read().splitlines()
    graph = [[int(c) for c in line] for line in lines]

    def at(pos: Vec) -> int | None:
        x, y = pos
        if 0 <= y < len(graph) and 0 <= x < len(graph[1]):
            return graph[y][x]
        return None

    heap = Heap([
        (0, State((0,0), (RIGHT), 0)),
        (0, State((0,0), (DOWN), 0)),
    ])
    seen = set()
    MAX_STRAIGHT = 10
    MIN_STRAIGHT = 4
    while heap:
        state: State
        dist, state = heap.pop()

        if state in seen:
            continue
        seen.add(state) 

        if state.pos == (len(graph[0])-1, len(graph)-1) and state.steps >= MIN_STRAIGHT:
            print(dist)
            break

        # Move forward
        if state.steps < MAX_STRAIGHT:
            next_pos = vec_add(state.pos, state.dir)
            dist_add = at(next_pos)
            if dist_add is not None:
                next_state = State(pos=next_pos, dir=state.dir, steps=state.steps+1)
                heap.push((dist+dist_add, next_state))

        # Turn
        if state.steps >= MIN_STRAIGHT:
            for next_dir in NEXT_DIR[state.dir]:
                next_pos = vec_add(state.pos, next_dir)
                dist_add = at(next_pos)
                if dist_add is not None:
                    next_state = State(pos=next_pos, dir=next_dir, steps=1)
                    heap.push((dist+dist_add, next_state))

main()
