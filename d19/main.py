from dataclasses import dataclass
from operator import gt, lt
from typing import Literal, Callable, cast
import math

Categories = Literal["x"] | Literal["m"] | Literal["a"] | Literal["s"]
Part = dict[Categories, int]

@dataclass
class Rule:
    what: Categories
    op: Callable[[int, int], bool]
    val: int
    action: str

@dataclass
class LastRule:
    action: str

Workflows = dict[str, list[Rule | LastRule]]

def parse_workflows(s: str) -> Workflows:
    result: Workflows = {}
    lines = s.splitlines()
    for line in lines:
        name, rules = line.split("{")
        result[name] = []
        *head, tail = rules.strip("}{").split(",")
        for rule in head:
            assert rule[0] in "xmas"
            what = cast(Categories, rule[0])
            assert rule[1] in "<>"
            op = gt if rule[1] == ">" else lt
            val, action = rule[2:].split(":")
            result[name].append(Rule(what, op, int(val), action))
        action = tail
        result[name].append(LastRule(action))
    return result

def parse_parts(s: str) -> list[Part]:
    result: list[Part] = []
    lines = s.splitlines()
    for line in lines:
        line = line.strip("}{")
        part: Part = {}
        categories = line.split(",")
        for category in categories:
            what, val = category.split("=")
            assert what in "xmas"
            what = cast(Categories, what)
            part[what] = int(val)
        result.append(part)
    return result

def validate_part(workflows: Workflows, part: Part) -> bool:
    cur_workflow = workflows["in"]
    while True:
        for rule in cur_workflow:
            if isinstance(rule, Rule):
                if rule.op(part[rule.what], rule.val):
                    action = rule.action
                else:
                    continue
            else:
                action = rule.action
            if action == "A":
                return True
            elif action == "R":
                return False
            else:
                cur_workflow = workflows[action]
                break

def combinations(workflows: Workflows, cur: str, restrictions: dict[Categories, tuple[int, int]]) -> int:
    if cur == "A":
        assert all(hi>=lo for lo, hi in restrictions.values())
        return math.prod(hi-lo+1 for lo, hi in restrictions.values())
    if cur == "R":
        return 0

    total = 0
    for rule in workflows[cur]:
        if isinstance(rule, LastRule):
            return total + combinations(workflows, rule.action, restrictions)
        new_restrictions = restrictions.copy()
        if rule.op == gt:
            new_restrictions[rule.what] = (max(restrictions[rule.what][0], rule.val+1), restrictions[rule.what][1])
        else:
            new_restrictions[rule.what] = (restrictions[rule.what][0], min(restrictions[rule.what][1], rule.val-1))
        if new_restrictions[rule.what][0] <= new_restrictions[rule.what][1]:
            total += combinations(workflows, rule.action, new_restrictions)
            
        # update restrictions with inverse of rule
        if rule.op == gt:
            restrictions[rule.what] = (restrictions[rule.what][0], min(restrictions[rule.what][1], rule.val))
        else:
            restrictions[rule.what] = (max(restrictions[rule.what][0], rule.val), restrictions[rule.what][1])
        if restrictions[rule.what][0] > restrictions[rule.what][1]:
            return total
    return total


def main() -> None:
    workflows_str, parts_str = open("d19/in.txt").read().split("\n\n")
    workflows = parse_workflows(workflows_str)
    parts = parse_parts(parts_str)
    valid_parts = []
    for part in parts:
        if validate_part(workflows, part):
            valid_parts.append(part)
    print(sum(part[what] for part in valid_parts for what in part)) 

    restrictions: dict[Categories, tuple[int, int]] = {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 4000)}
    part2 = combinations(workflows, "in", restrictions)
    print(part2)

main()
