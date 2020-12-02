def question1():
    answer = 0
    with open("input.txt") as f:
        lines = f.readlines()
        for line in lines:
            line = line.strip()
            operator = line[0]
            number = line[1:]
            if operator == "+":
                answer += int(number)
            elif operator == "-":
                answer -= int(number)
    print(answer)


def question2():
    frequencies = {}
    answer = 0
    frequencies[answer] = None
    with open("input.txt") as f:
        lines = f.readlines()
        duplicate_found = False
        while duplicate_found == False:
            for line in lines:
                line = line.strip()
                operator = line[0]
                number = line[1:]
                if operator == "+":
                    answer += int(number)
                    if answer in frequencies:
                        print(answer)
                        duplicate_found = True
                        break
                    else:
                        frequencies[answer] = None
                elif operator == "-":
                    answer -= int(number)
                    if answer in frequencies:
                        print(answer)
                        duplicate_found = True
                        break
                    else:
                        frequencies[answer] = None
    # print(frequencies)
question1()
question2()
