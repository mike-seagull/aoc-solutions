from math import floor

def readfile(filepath):
    input = []
    with open(filepath) as fp:
       line = fp.readline()
       while line:
           input.append(int(line.strip()))
           line = fp.readline()
    return input

def part1(masses):
    total_fuel = 0
    for mass in masses:
        total_fuel += part1_func(mass)
    return total_fuel

def part1_func(mass):
    return floor(mass / 3) - 2


def part2(masses):
    total_fuel = 0
    for mass in masses:
        total_fuel += part2_func(mass)
    return total_fuel

def part2_func(mass):
    total_fuel = 0
    doAgain = True
    while doAgain == True:
        fuel = part1_func(mass)
        if fuel >= 0:
            total_fuel += fuel
            mass = fuel
        else:
            doAgain = False
    return total_fuel

def test_part1():
    assert part1_func(12) == 2
    assert part1_func(14) == 2
    assert part1_func(1969) == 654
    assert part1_func(100756) == 33583
def test_part2():
    assert part2_func(14) == 2
    assert part2_func(1969) == 966
    assert part2_func(100756) == 50346


if __name__ == '__main__':
    masses = readfile("day1_input.txt")
    print("Part1 = " + str(part1(masses)))
    print("Part2 = " + str(part2(masses)))

