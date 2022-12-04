"""
day 4: https://adventofcode.com/2022/day/4
"""

from typing import List, Set
from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)


test_data = fetch_data('2022/day4/test.txt')
data = fetch_data('2022/day4/data.txt')


def convert_rangestr_to_set(rangestr: str) -> Set:
    rangelist = rangestr.split("-")
    return set(range(int(rangelist[0]), int(rangelist[-1]) + 1))


def split_into_sections(sections: str) -> List[Set]:
    split = sections.split(',')
    return [convert_rangestr_to_set(split[0]), convert_rangestr_to_set(split[1])]


def detect_any_intersection(s1: set, s2: set) -> bool:
    return s1.intersection(s2)


def detect_full_intersection(s1: set, s2: set) -> bool:
    if len(s1) < len(s2):
        subset = s1
        superset = s2
    else:
        subset = s2
        superset = s1

    return subset.issubset(superset)


def sum_intersections(data: list, full: bool):
    if full:
        strategy = detect_full_intersection
    else:
        strategy = detect_any_intersection

    sum = 0
    for sections in data:
        s1, s2 = split_into_sections(sections)
        if strategy(s1, s2):
            sum += 1

    return sum


def test_part1():
    return sum_intersections(test_data, full=True)


def part1():
    return sum_intersections(data, full=True)


def test_part2():
    return sum_intersections(test_data, full=False)


def part2():
    return sum_intersections(data, full=False)


assert test_part1() == 2
print("Part 1", part1())
assert test_part2() == 4
print("Part 2", part2())
