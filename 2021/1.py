from util import get_input_from_file

def part1(input_list):
    prev_depth = None
    count = 0
    for depth in input_list:
        depth = int(depth)
        if prev_depth is not None and depth > prev_depth:
            count = count + 1
        prev_depth = depth
    print(count)


def part2(input_list):
    prev_sum = None
    count = 0
    for i in range(2, len(input_list)):
        sum  = int(input_list[i]) + int(input_list[i - 1]) + int(input_list[i - 2])
        if prev_sum is not None and sum > prev_sum:
            count = count + 1
        prev_sum = sum   
    print(count)


def main():
    input = get_input_from_file("1.input.txt")
    # input = get_input_from_file("1.test.txt")
    part1(input)
    part2(input)


main()
