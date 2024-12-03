from py_aoc.helpers import get_data
from copy import deepcopy

"""
The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?
"""

test_data, data = get_data(2024, 2)

def compose_report(report: str) -> list[int]:
    return report.split(' ')

def is_report_safe_part_1(report: str) -> bool:
    data = compose_report(report)
    count_inc = 0
    count_dec = 0
    for i in range(0, len(data)-1):
        diff = int(data[i]) - int(data[i+1])
        if diff > 0:
            count_inc += 1
        elif diff < 0:
            count_dec += 1
        else:
            return False

        if count_inc > 0 and count_dec > 0:
            return False

        if abs(diff) > 3:
            return False
    return True


def run_part_1(data):
    num_safe = 0
    for row in data:
        if is_report_safe_part_1(row):
            num_safe+=1

    return num_safe

assert run_part_1(test_data) == 2
print(f'Day 2 Part 1: {run_part_1(data)} safe!')

def compose_report_v2(report: str) -> list[int]:
    str_report = report.split(' ')
    return [int(n) for n in str_report]
    
def get_adjacent_indexes(report: list, err_index: int) -> list:
    adjacent_indexes = []
    if 0 < err_index < len(report) - 1:
        adjacent_indexes = [err_index - 1, err_index, err_index + 1]
    elif err_index == 0:
        adjacent_indexes = [err_index, err_index + 1]
    elif err_index == len(report) - 1:
        adjacent_indexes = [err_index-1, err_index]

    return adjacent_indexes

def is_report_safe_part_2(report: list, problem_dampener=False) -> bool:
    count_inc = 0
    count_dec = 0
    print(f"{report} checking")

    for i in range(0, len(report)-1):
        diff = report[i] - report[i+1]

        if abs(diff) > 3:
            print('more than 3')
            if problem_dampener:
                print(f"failed on retry! {report}")
                return False
            else:
                adj_indexes = get_adjacent_indexes(report, i)
                print(f"> 3 Error found! Checking adjacent indexes {adj_indexes}")
                for index in adj_indexes:
                    copy = deepcopy(report)
                    del(copy[index])
                    if is_report_safe_part_2(copy, problem_dampener=True):
                        return True
                return False
                
        if diff > 0:
            count_dec += 1
        elif diff < 0:
            count_inc += 1
        else:
            print(f'same number! {report[i]} and {report[i+1]}')
            if problem_dampener:
                print(f"failed on retry! {report}")
                return False
            else:
                adj_indexes = get_adjacent_indexes(report, i)
                print(f"Same number error found! Checking adjacent indexes {adj_indexes}")
                for index in adj_indexes:
                    copy = deepcopy(report)
                    del(copy[index])
                    if is_report_safe_part_2(copy, problem_dampener=True):
                        return True
                return False
                                    
        if count_inc > 0 and count_dec > 0:
            # print('both inc and dec!')
            if problem_dampener:
                print(f"failed on retry! {report}")
                return False
            else:
                adj_indexes = get_adjacent_indexes(report, i)
                print(f"Inc/Dec error found! Checking adjacent indexes {adj_indexes}")
                for index in adj_indexes:
                    copy = deepcopy(report)
                    del(copy[index])
                    if is_report_safe_part_2(copy, problem_dampener=True):
                        return True
                return False
                    
    print(f"{report} report safe!")
    return True

def run_part_2(data):
    num_safe = 0
    for row in data:
        report = compose_report_v2(row)
        print(f"{report} checking top level")
        if is_report_safe_part_2(report):

            num_safe+=1
    
    print("num safe ", num_safe)
    return num_safe

assert run_part_2(test_data) == 4
print(f'day 2 part 2: {run_part_2(data)} safe!')