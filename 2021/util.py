def get_input_from_file(filename):
    with open(filename, "r") as f:
       lines = f.read()
       return lines.splitlines()
