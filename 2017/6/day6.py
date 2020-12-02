#!/usr/bin/env python3
import operator

def part1():
    input_file = open("part1_input")
    puzzle_input = input_file.readline()
    bank = [int(numeric_string) for numeric_string in puzzle_input.split()]
    old_banks = {}
    bank_string = ''.join(str(e) for e in bank)
    old_banks[bank_string] = 1
    while True:
        # get the first max value and its index
        index, value = max(enumerate(bank), key=operator.itemgetter(1))
        iterator_index = index
        # while bank[index] > 1:
        for i in range(value):
            iterator_index = iterator_index + 1
            if iterator_index > len(bank) - 1:  # loop back to the beginning
                iterator_index = 0

            bank[iterator_index] = bank[iterator_index] + 1
            bank[index] = bank[index] - 1

        bank_string = ''.join(str(e) for e in bank)
        if bank_string in old_banks:
            break
        else:
            old_banks[bank_string] = 1
    return (len(old_banks))


print("Part1 = %s" % part1())

