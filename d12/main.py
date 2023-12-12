from functools import cache

def arrangements(line: str) -> int:
    left, right = line.split()
    left = "?".join([left for _ in range(5)])
    right = ",".join([right for _ in range(5)])
    broken = [int(c) for c in right.split(",")]

    @cache
    def solve(left_idx: int, broken_idx: int) -> int:
        # 1. skip dots
        # 2.1 if # then continue until .
        # 2.2 if ? then two options:
        #   recursive calc (left_idx+steps, broken_idx)
        #   assume ? is # and continue until .
        while left_idx < len(left) and left[left_idx] == ".":
            left_idx += 1
        
        if left_idx >= len(left):
            return broken_idx > len(broken) - 1
        
        if broken_idx >= len(broken):
            return not any(left[i] == "#" for i in range(left_idx, len(left)))
        
        if left[left_idx] == "?":
            first_option = solve(left_idx+1, broken_idx) # assume ? is .
        elif left[left_idx] == "#":
            first_option = 0
        else:
            raise ValueError("unknown symbol")
        
        target_broken = broken[broken_idx]
        target_broken -= 1
        while target_broken > 0:
            left_idx += 1
            if left_idx >= len(left):
                return first_option # still need some target broken, but string ended
            if left[left_idx] == ".":
                return first_option # still need some target broken, but sequence stopped
            target_broken -= 1
        if left_idx + 1 >= len(left):
            is_ok = broken_idx == len(broken) - 1
            return is_ok + first_option 
        if left[left_idx+1] != "#":
            return first_option + solve(left_idx+2, broken_idx+1)
        else:
            return first_option
    
    res = solve(0, 0)
    return res

def main():
    lines = open("d12/in.txt").read().splitlines()
    total = 0
    for l in lines:
        a = arrangements(l)
        total += a
    print(total)


if __name__ == "__main__":
    main()
