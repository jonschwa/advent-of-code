"""
Day 7: https://adventofcode.com/2022/day/7
"""

from dataclasses import dataclass
from typing import List
from py_aoc.helpers import read_file_to_list


def fetch_data(path) -> list:
    return read_file_to_list(path)


test_data = fetch_data('2022/day7/test.txt')
data = fetch_data('2022/day7/data.txt')


@dataclass
class File:
    name: str
    size: int

@dataclass
class Command:
    method: str
    arg: str = None

class Node:

    data = None

    def __init__(self, label, data=None, parent=None):
        
        self.parent = parent
        self.children: List[Node] = []
        self.label: str = label
        if data:
            self.data = data
            

    def insert(self, new_node):
        self.children.append(new_node)

    def traverse(self, label):
        for child in self.children:
            if child.label == label:
                return child

    def home(self):
        top = False
        pwd = self
        while not top:
            if parent := pwd.parent:
                pwd = parent
            else:
                top = True
                return pwd

    def size(self) -> int:
        size = 0
        if self.data:
            size = self.data.size
        if len(self.children) > 0:
            size = self.dir_size()
        return size

    def dir_size(self) -> int:
        size = 0
        for child in self.children:
            size += child.size()
        return size
        

    def parent_count(self) -> int:
        top = False
        pwd = self
        parent_count = 0
        while not top:
            if parent := pwd.parent:
                pwd = parent
                parent_count += 1
            else:
                top = True
        
        return parent_count

    def print(self):
        print(str(self))
        for child in self.children:
            child.print()

    def __str__(self) -> str:

        padding = ""
        pc = self.parent_count()
        if pc > 0:
            padding = "-" * pc
        if not self.data:
            return(f"{padding} {self.label} (dir)")
        else:
            return(f"{padding} {self.data.name} {self.data.size}")

def parse_command(command: str) -> Command:
    return Command(*command[2:].split(" "))

def add_node(input: str, pwd: Node) -> Node:
    split = input.split(" ")
    if split[0] == "dir": # adding a dir
        new_node = Node(label=split[1], parent=pwd)
    else:
        new_node = Node(label=split[1], data=File(size=int(split[0]), name=split[1]), parent=pwd)
    
    pwd.insert(new_node)
    return pwd


def build_fs(input):
    fs = Node("/", data=None)
    current_command = None
    pwd = fs

    for instruction in input:
        if instruction[0] == "$":
            command = parse_command(instruction)
            current_command = command
        if current_command.method == "cd":
            if current_command.arg == "..":
                pwd = pwd.parent
            elif current_command.arg == "/":
                pwd = pwd.home()
            else:
                pwd = pwd.traverse(current_command.arg)
        elif current_command.method == "ls":  # going to be adding some files to the current dir
            #JANK! make sure this isn't adding a command as a node
            if instruction[0] == "$":
                continue
            pwd = add_node(instruction, pwd)

    #return the fs @ the top level
    return pwd.home()
      

def test_part1():
    fs = build_fs(test_data)
    fs.print()

    #throwing a few assertions in here
    assert fs.traverse("a").size() == 94853
    assert fs.traverse("d").size() == 24933642
    assert fs.traverse("a").traverse("e").size() == 584
    assert fs.size() == 48381165


test_part1()