"""
Day 5!

The inputs (im gonna just not parse these lol)

live: 

[N]     [C]                 [Q]    
[W]     [J] [L]             [J] [V]
[F]     [N] [D]     [L]     [S] [W]
[R] [S] [F] [G]     [R]     [V] [Z]
[Z] [G] [Q] [C]     [W] [C] [F] [G]
[S] [Q] [V] [P] [S] [F] [D] [R] [S]
[M] [P] [R] [Z] [P] [D] [N] [N] [M]
[D] [W] [W] [F] [T] [H] [Z] [W] [R]
 1   2   3   4   5   6   7   8   9 

test: 

    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

"""
from collections import deque
from copy import deepcopy
from typing import List
from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)


test_data = fetch_data('2022/day5/test.txt')
data = fetch_data('2022/day5/data.txt')

# Just hardcoding this because I'm a busy person with things to do!
test_init = [deque("ZN"), deque("MCD"), deque("P")]
init = [
    deque("DMSZRFWN"),
    deque("WPQGS"),
    deque("WRVQFNJC"),
    deque("FZPCGDL"),
    deque("TPS"),
    deque("HDFWRL"),
    deque("ZNDC"),
    deque("WNRFVSJQ"),
    deque("RMSGZWV"),
]


def move_single(num: int, origin: int, destination: int, state: List[str]) -> List[str]:
    i = 1
    # shift the index
    origin -= 1
    destination -= 1
    while i <= num:
        to_move = state[origin].pop()
        state[destination].append(to_move)
        i += 1

    return state


def move_multiple(num: int, origin: int, destination: int, state: List[str]) -> List[str]:
    i = 1
    # shift the index
    origin -= 1
    destination -= 1

    group = deque()
    while i <= num:
        to_move = state[origin].pop()
        group.appendleft(to_move)
        i += 1

    state[destination].extend(group)
    return state


def parse_instruction(instruction: str) -> List[int]:
    # quick & dirty!
    split = instruction.split(" ")
    return [int(split[1]), int(split[3]), int(split[5])]


def return_message(init_state, data, multiple=False):
    move_strategy = move_multiple if multiple else move_single
    state = init_state
    for instruction in data:
        [num, origin, destination] = parse_instruction(instruction)
        # print([num, origin, destination])
        state = move_strategy(num, origin, destination, state)

    message = ""
    for col in state:
        message += col[-1]

    return message


def test_part_1():
    start = deepcopy(test_init)
    return return_message(start, test_data, multiple=False)


def part_1():
    start = deepcopy(init)
    print("Part 1:", return_message(start, data))


def test_part_2():
    start = deepcopy(test_init)
    return return_message(start, test_data, multiple=True)


def part_2():
    start = deepcopy(init)
    print("Part 2:", return_message(start, data, multiple=True))


assert test_part_1() == "CMZ"
part_1()
assert test_part_2() == "MCD"
part_2()
