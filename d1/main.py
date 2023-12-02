import re

nums = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
FIRST_LAST_PATTERN = re.compile(rf'(?=({"[1-9]|" + "|".join(nums)}))')
DIGIT_MAP = {str(i): i for i in range(1, 10)} | {n: i for i, n in enumerate(nums, 1)}

def get_line_digits(line: str) -> int:
    matches = FIRST_LAST_PATTERN.findall(line)
    return DIGIT_MAP[matches[0]] * 10 + DIGIT_MAP[matches[-1]]

print(sum(get_line_digits(line) for line in open("d1/in.txt").readlines()))
