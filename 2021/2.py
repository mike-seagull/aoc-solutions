from util import get_input_from_file

def part1(input):
    x = y = 0
    for line in input:
        command, move_value = line.split()
        if command == "forward":
            x += int(move_value)
        elif command == "up":
            y -= int(move_value)
        elif command == "down":
            y += int(move_value)
    print(x*y)

def part2(input):
    x = y = aim = 0
    for line in input:
        command, move_value = line.split()
        if command == "forward":
            x += int(move_value)
            y += (int(move_value) * aim)
        elif command == "up":
            aim -= int(move_value)
        elif command == "down":
            aim += int(move_value)
        # print(x, y, aim)

    print(x*y)
def main():
    # input = get_input_from_file("2.test.txt")
    input = get_input_from_file("2.txt")
    part1(input)
    part2(input)

main()
