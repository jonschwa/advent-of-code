"""
day 3 https://adventofcode.com/2022/day/3
"""

import string
from typing import Set, List
from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)


test_data = fetch_data('2022/day3/test.txt')
data = fetch_data('2022/day3/data.txt')


PRIORITIES = list(string.ascii_letters)
PRIORITY_MAP = dict(zip(PRIORITIES, list(range(1,53))))

def get_priority_of_item(item: str) -> int:
    # I assume range just uses a hashmap under the hood, but being explicit...
    #return PRIORITIES.index(item) + 1
    return PRIORITY_MAP[item]

def split_pack(pack: str) -> List[Set]:
    mid = int(len(pack)/2)
    return [set(pack[0:mid]), set(pack[mid:])]

def get_sum_of_dupe_priorities(data):
    sum = 0
    for pack in data:
        c1, c2 = split_pack(pack)
        dupe = c1.intersection(c2)
        sum += get_priority_of_item([*dupe][0])
    return sum

def get_sum_of_badge_priorities(data):
    sum = 0
    idx = 0
    while idx <= len(data) - 3:
        c1 = set(data[idx])
        c2 = set(data[idx+1])
        c3 = set(data[idx+2])
        badge = c1.intersection(c2,c3)
        sum += get_priority_of_item([*badge][0])
        idx += 3

    return sum


def test_part_1():
    return get_sum_of_dupe_priorities(test_data)

def part_1():
    return get_sum_of_dupe_priorities(data)

def test_part_2():
    return get_sum_of_badge_priorities(test_data)

def part_2():
    return get_sum_of_badge_priorities(data)

#priority tests
assert get_priority_of_item("p") == 16
assert get_priority_of_item("L") == 38

assert test_part_1() == 157
print("Part 1:", part_1())

assert test_part_2() == 70
print("Part 2:", part_2())