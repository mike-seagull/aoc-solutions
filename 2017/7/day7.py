#!/usr/bin/env python3
import re


def parseInput(filename):
    programs = {}
    with open(filename) as f:
        for line in f:
            lineSplit = list(filter(None, re.split(' \(|, |\) -> |\)|\s', line)))
            programName = lineSplit.pop(0)
            programs[programName] = {}
            programs[programName]['weight'] = int(lineSplit.pop(0))
            if len(lineSplit) > 1:
                programs[programName]['children'] = lineSplit
            else:
                programs[programName]['children'] = []
    return programs


def part1():
    programs = parseInput("part1_input")
    all_programs = list(programs.keys())
    for programName, values in programs.items():
        if len(programs[programName]['children']) == 0:
            try:
                all_programs.remove(programName)
            except ValueError:  # some of the elements will already have been removed
                pass
        else:
            for program in values['children']:
                try:
                    all_programs.remove(program)
                except ValueError:
                    pass  # some of the elements will already have been removed
    return all_programs[0]

def part2():
    programs = parseInput('part1_input')
    #root = part1()
    for programName, values in programs.items():
        


def main():
    print("part1 = %s" % part1())


main()
