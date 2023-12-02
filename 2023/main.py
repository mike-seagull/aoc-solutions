import sys
day = __import__(sys.argv[1])

print(f"Day {sys.argv[1]}")

day.part1()
day.part2()