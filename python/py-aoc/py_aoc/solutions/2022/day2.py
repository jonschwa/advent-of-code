"""
day 2: https://adventofcode.com/2022/day/2

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

"""

from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)


test_data = fetch_data('2022/day2/test.txt')
data = fetch_data('2022/day2/data.txt')

ROCK = "X"
PAPER = "Y"
SCISSORS = "Z"

WIN = "W"
LOSE = "L"
DRAW = "D"

RESULT_SCORES = {
    WIN: 6,
    DRAW: 3,
    LOSE: 0,
}

PLAY_SCORES = {
    ROCK: 1,
    PAPER: 2,
    SCISSORS: 3
}

SYMBOL_WIN_MAP = {
    "X": LOSE, "Y": DRAW, "Z": WIN
}

# a mapping of which symbols W/L/D against which
CONDITIONS_MAP = {
    "A": {SCISSORS: LOSE, ROCK: DRAW, PAPER: WIN},
    "B": {ROCK: LOSE, PAPER: DRAW, SCISSORS: WIN},
    "C": {PAPER: LOSE, SCISSORS: DRAW, ROCK: WIN},
}

# a mapping of W/L/D for each symbol to the opposing symbol
INVERSE_CONDITIONS_MAP = {
    "A": {LOSE: SCISSORS, DRAW: ROCK, WIN: PAPER},
    "B": {LOSE: ROCK, DRAW: PAPER, WIN: SCISSORS},
    "C": {LOSE: PAPER, DRAW: SCISSORS, WIN: ROCK},
}


def calculate_score(match_data):
    score = 0
    for round in match_data:
        result = CONDITIONS_MAP[round[0]][round[-1]]
        score += RESULT_SCORES[result] + PLAY_SCORES[round[-1]]

    return score


def calculate_score_pt2(match_data):
    score = 0
    for round in match_data:
        opposing = INVERSE_CONDITIONS_MAP[round[0]][SYMBOL_WIN_MAP[round[-1]]]
        score += PLAY_SCORES[opposing] + \
            RESULT_SCORES[SYMBOL_WIN_MAP[round[-1]]]

    return score


def test_part1():
    return calculate_score(test_data)


def part1():
    score = calculate_score(data)
    print("Part 1: ", score)


def test_part2():
    return calculate_score_pt2(test_data)


def part2():
    score = calculate_score_pt2(data)
    print("Part 2: ", score)


assert test_part1() == 15
part1()
assert test_part2() == 12
part2()
