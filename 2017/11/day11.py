#!/usr/bin/env python3


class Cube:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def steps(self):
        return (abs(self.x + 0) + abs(self.y + 0) + abs(self.z + 0)) / 2


class Axial:
    def __init__(self, q, r):
        self.q = q
        self.r = r
    def n(self):
        self.r = self.r - 1
    def s(self):
        self.r = self.r + 1
    def se(self):
        self.q = self.q + 1
    def nw(self):
        self.q = self.q - 1
    def sw(self):
        self.s()
        self.nw()
    def ne(self):
        self.se()
        self.n()

    def convertToCube(self):
        return Cube(self.q, ((self.q * -1) - self.r), self.r)

    def steps(self):
        x = self.q
        y = (self.q * -1) - self.r
        z = self.r
        return (abs(self.q) + abs((self.q * -1) - self.r) + abs(self.r)) / 2


def examples():
    ex1 = Axial(0, 0)
    ex1.nw()
    ex1.nw()
    ex1.nw()
    print("Example1: %i" % ex1.steps())

    ex2 = Axial(0, 0)
    ex2.ne()
    ex2.ne()
    ex2.sw()
    ex2.sw()
    print("Example2: %i" % ex2.steps())

    ex3 = Axial(0, 0)
    ex3.ne()
    ex3.ne()
    ex3.s()
    ex3.s()
    print("Example3: %i" % ex3.steps())

    ex4 = Axial(0, 0)
    ex4.se()
    ex4.sw()
    ex4.se()
    ex4.sw()
    ex4.sw()
    print("Example4: %i" % ex4.steps())


def part1():
    part1_input = open("part1_input").readline()
    movements = part1_input.split(",")
    child_process = Axial(0, 0)
    for move in movements:
        if move == "n":
            child_process.n()
        elif move == "s":
            child_process.s()
        elif move == "nw":
            child_process.nw()
        elif move == "sw":
            child_process.sw()
        elif move == "ne":
            child_process.ne()
        elif move == "sw":
            child_process.sw()
        else:
            "invalid move: %s" % move
    print("Part1: %i" % child_process.steps())

part1()
#examples()

