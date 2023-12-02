from aocd import get_data
inputlines = get_data(day=2, year=2023).splitlines()

def part1():
    total = 0
    for game in inputlines:
        tmp = game.split(":")
        game_number = tmp[0].strip().split(" ")[1]
        sets_of_cubes = tmp[1].strip().split(";")
        game_is_possible = True
        for set_of_cubes in sets_of_cubes:
            grab = set_of_cubes.split(",")
            for cubes_by_color in grab:
                tmp = cubes_by_color.strip().split(" ")
                num_of_cubes = int(tmp[0].strip())
                color_of_cubes = tmp[1].strip()
                if color_of_cubes == "red" and num_of_cubes > 12:
                    game_is_possible = False
                if color_of_cubes == "green" and num_of_cubes > 13:
                    game_is_possible = False
                if color_of_cubes == "blue" and num_of_cubes > 14:
                    game_is_possible = False
        if game_is_possible:
            total += int(game_number)
    print(total)
        

def part2():
    total = 0
    for game in inputlines:
        min_by_color = {
            "red": 1,
            "green": 1,
            "blue": 1
        }
        tmp = game.split(":")
        game_number = tmp[0].strip().split(" ")[1]
        sets_of_cubes = tmp[1].strip().split(";")
        for set_of_cubes in sets_of_cubes:
            grab = set_of_cubes.split(",")
            for cubes_by_color in grab:
                tmp = cubes_by_color.strip().split(" ")
                num_of_cubes = int(tmp[0].strip())
                color_of_cubes = tmp[1].strip()
                if num_of_cubes > min_by_color[color_of_cubes]:
                    min_by_color[color_of_cubes] = num_of_cubes
        total += (min_by_color["blue"] * min_by_color["green"] * min_by_color["red"])
    print(total)