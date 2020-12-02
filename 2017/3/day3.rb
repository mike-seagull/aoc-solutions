#/usr/bin/env ruby
def find_level_and_last_number_in_level(number)
    last_number_of_level = 1
    level = 0
    square_root_of_last_number = 1
    while last_number_of_level < number
        level = level + 1
        square_root_of_last_number = square_root_of_last_number + 2
        last_number_of_level = square_root_of_last_number * square_root_of_last_number
    end
    return level, last_number_of_level, square_root_of_last_number
end

def find_distance_from_midpoint(square_root_of_last_number, high_corner, number, level)
	last_number_of_level = high_corner
	lower_corner = high_corner
	for i in 1..4
		lower_corner = lower_corner - (square_root_of_last_number - 1)
		if i == 4
			lower_corner = high_corner
			high_corner = last_number_of_level
			break
		elsif high_corner >= number and lower_corner <= number
			break
		else
			high_corner = lower_corner
		end

	end

	if number == lower_corner or number == high_corner
		if high_corner == last_number_of_level and lower_corner == high_corner - ((square_root_of_last_number - 1)*3)
			return level, (lower_corner - lower_corner), high_corner, lower_corner
		else
			return level, high_corner - level, high_corner, lower_corner
		end
	else

		# determine midpoint
		if high_corner == last_number_of_level and lower_corner == high_corner - ((square_root_of_last_number - 1)*3)
			high_corner = lower_corner
			lower_corner = high_corner - (square_root_of_last_number - 1)
		end
		midpoint = high_corner - ((high_corner - lower_corner)/2)
		return (number - midpoint), midpoint
	end
end

def part1()
	part1_input = 368078
	level, last_number_of_level, square_root_of_last_number = find_level_and_last_number_in_level(part1_input)
	distance_from_mid, midpoint, high_corner, lower_corner = find_distance_from_midpoint(square_root_of_last_number, last_number_of_level, part1_input, level)
	puts "part1 = #{(distance_from_mid).abs + level}"
end

part1()
