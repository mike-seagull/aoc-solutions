#!/usr/bin/env python3
import re


def parseInput(filename):
    graph = {}
    for line in open(filename):
        lineSplit = list(filter(None, re.split(' <-> |, |\n', line)))
        vertex = lineSplit.pop(0)
        graph[vertex] = lineSplit
    return graph


def iterative_dfs(graph, start_vertex):
    '''iterative depth first search from start'''
    path = []
    stack = [start_vertex]
    while stack:
        vertex = stack.pop(0)
        if vertex not in path:
            path.append(vertex)
            stack = graph[vertex] + stack  # faster than just pushing to the front
    return path


def test_input():
    graph = {
        '0': ['2'],
        '1': ['1'],
        '2': ['0', '3', '4'],
        '3': ['2', '4'],
        '4': ['2', '3', '6'],
        '5': ['6'],
        '6': ['4', '5']
    }
    print("** TEST INPUT **")
    return graph


def part1(graph):
    if not graph:
        graph = parseInput("part1_input")

    path = iterative_dfs(graph, '0')
    print("part 1 = %i" % len(path))


def part2(graph):
    groups = []
    for vertex in graph.keys():
        vertexIsInListOfGroups = False
        for group in groups:
            if vertex in group:
                vertexIsInListOfGroups = True
        if not vertexIsInListOfGroups:
            path = iterative_dfs(graph, vertex)
            groups.append(path)
    print("part 2 = %i" % len(groups))


def main():
    test_graph = test_input()
    puzzle_input = parseInput("part1_input")
    part1(puzzle_input)
    part2(puzzle_input)


main()
