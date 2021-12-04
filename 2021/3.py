from util import get_input_from_file

def part1(input):
    zeros = [0] * len(input[0])
    ones = [0] * len(input[0])
    for bits in input:
        for i in range(0, len(bits)):
            bit = bits[i]
            if bit == "0":
                zeros[i] += 1
            elif bit == "1":
                ones[i] += 1
    gamma = epsilon = ""

    for i in range(0, len(zeros)):
        if zeros[i] > ones[i]:
            gamma += "0"
            epsilon += "1"
        elif zeros[i] < ones[i]:
            epsilon += "0"
            gamma += "1"
    
    print(int(gamma, 2) * int(epsilon, 2))    


def part2(oxygen, co2, i=0):
    # recursive
    if len(oxygen) == 1 and len(co2) == 1:
        print(int(oxygen[0], 2) * int(co2[0], 2))
    else:
        zeros = []
        ones = []
        if len(oxygen) > 1:
            for bits in oxygen:
                bit = bits[i]
                if bit == "0":
                    zeros.append(bits)
                elif bit == "1":
                    ones.append(bits)
            if len(zeros) > len(ones):
                oxygen = zeros
            else:
                oxygen = ones
        zeros = []
        ones = []
        if len(co2) > 1:
            for bits in co2:
                bit = bits[i]
                if bit == "0":
                    zeros.append(bits)
                elif bit == "1":
                    ones.append(bits)
            if len(zeros) > len(ones):
                co2 = ones
            else:
                co2 = zeros
        part2(oxygen, co2, i+1)


def main():
    # input = get_input_from_file("3.test.txt")
    input = get_input_from_file("3.txt")
    part1(input)
    part2(input, input)

main()