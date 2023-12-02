from aocd import get_data
inputlines = get_data(day=1, year=2023).splitlines()

def part1():
    total = 0
    for line in inputlines:
        number = 0
        for c in line:
            if c.isdigit():
                number += int(c)*10
                break
        for c in reversed(line):
            if c.isdigit():
                number += int(c)
                break
        total += number
    print(total)

def part2():
    total = 0
    for line in inputlines:
        line = line.replace("one", "on1e").replace("two", "tw2o").replace("three", "thre3e").replace("four", "fou4r").replace("five", "fiv5e").replace("six", "si6x").replace("seven", "seve7n").replace("eight", "eigh8t").replace("nine", "nin9e")
        number = 0
        for c in line:
            if c.isdigit():
                number += int(c)*10
                break
        for c in reversed(line):
            if c.isdigit():
                number += int(c)
                break
        total += number
    print(total)    