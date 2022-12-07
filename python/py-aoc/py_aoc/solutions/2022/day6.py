"""
Day 6: https://adventofcode.com/2022/day/6
"""

from collections import deque
from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)[0]


data = fetch_data('2022/day6/data.txt')


def check_if_unique(chars: list) -> bool:
    return len(chars) == len(set(chars))


def find_marker_position(input: str, msg_count):
    i = 1
    current_window = []
    while i <= len(input):
        current_window.append(input[i-1])

        if i > msg_count:
            current_window.pop(0)
            if check_if_unique(current_window):
                return i

        i += 1


assert find_marker_position('mjqjpqmgbljsphdztnvjfqwrcgsmlb', 4) == 7
assert find_marker_position('bvwbjplbgvbhsrlpgdmjqwftvncz', 4) == 5
assert find_marker_position('nppdvjthqldpwncqszvftbrmjlhg', 4) == 6
assert find_marker_position('nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg', 4) == 10
assert find_marker_position('zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw', 4) == 11

print("Part 1: ", find_marker_position(data, 4))

assert find_marker_position('mjqjpqmgbljsphdztnvjfqwrcgsmlb', 14) == 19
assert find_marker_position('bvwbjplbgvbhsrlpgdmjqwftvncz', 14) == 23
assert find_marker_position('nppdvjthqldpwncqszvftbrmjlhg', 14) == 23
assert find_marker_position('nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg', 14) == 29
assert find_marker_position('zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw', 14) == 26

print("Part 2: ", find_marker_position(data, 14))
