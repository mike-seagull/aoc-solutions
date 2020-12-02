def q1():
	with open("input.txt") as f:
	# with open("test.txt") as f:
		lines = f.readlines()
		twice = 0
		thrice = 0
		for line in lines:
			chars = {}
			line = line.strip()
			for char in line:
				if char not in chars:
					chars[char] = 0 
				chars[char] += 1
			has2 = False
			has3 = False
			for char in chars:
				if chars[char] == 2:
					if not has2:
						twice += 1
						has2 = True
				elif chars[char] == 3:
					if not has3:
						thrice += 1
						has3 = True
	print(twice * thrice)

def q2():
	""" super dirty way to do this """
	with open("input.txt") as f:
		lines = f.readlines()
		for box_id1 in lines:
			for box_id2 in lines:
				box_id1 = box_id1.strip()
				box_id2 = box_id2.strip()
				count = 0
				difference = [] 
				for a, b in zip(box_id1, box_id2):
					if a != b:
						count += 1
						difference =[a, b]
				if count == 1:
					string = ""
					for char in box_id1:
						if char != difference[0]:
							string += char
					return string

q1()
print(q2())