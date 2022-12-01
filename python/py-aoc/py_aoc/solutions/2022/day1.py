"""
--- Day 1: Calorie Counting ---
Santa's reindeer typically eat regular reindeer food, but they need a lot of magical energy to deliver presents on Christmas. For that, their favorite snack is a special type of star fruit that only grows deep in the jungle. The Elves have brought you on their annual expedition to the grove where the fruit grows.
To supply enough magical energy, the expedition needs to retrieve a minimum of fifty stars by December 25th. Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.
Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).
The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.
For example, suppose the Elves finish writing their items' Calories and end up with the following list:
"""
from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)


test_data = fetch_data('2022/day1/test.txt')
data = fetch_data('2022/day1/data.txt')


def group_calories(data) -> list:
    elf_totals = []
    current_total = 0

    for i, cals in enumerate(data):
        if cals:
            current_total += int(cals)
            if i == len(data) - 1:
                elf_totals.append(current_total)
                continue
        else:
            elf_totals.append(current_total)
            current_total = 0
            continue

    return elf_totals


def get_max_cal_group(data) -> dict:
    elf_totals = group_calories(data)
    return max(elf_totals)


def get_sum_top_3_elves(data) -> int:
    elf_totals = group_calories(data)
    ordered = sorted(elf_totals, reverse=True)
    return sum(ordered[0:3])


def test_max_cals():
    assert get_max_cal_group(test_data) == 24000


def part_1():
    return get_max_cal_group(data)


def test_top_3_cals():
    assert get_sum_top_3_elves(test_data) == 45000


def part_2():
    return get_sum_top_3_elves(data)


test_max_cals()
print("part 1", part_1())
test_top_3_cals()
print("part 2", part_2())
