import numpy as np
data = []

max_level_change = 3

with open('aoc_2_2024.txt', 'r') as f:
    data = [x.strip().split() for x in f.readlines()]

for line_idx, line in enumerate(data):
    new_line = list(map(lambda x: int(x), line))
    data[line_idx] = new_line

#part 1
num_safe = 0
for line in data:
    safe = True
    sign_check = line[0] - line[1]
    if abs(sign_check) > max_level_change or sign_check == 0:
        safe = False
    for i in range(1, len(line)-1):
        diff = line[i] - line[i+1]
        if abs(diff) > max_level_change or diff * sign_check <= 0 or diff == 0:
            safe = False
    num_safe += safe

print(num_safe)

#part 2

num_safe = 0
for line in data:
    safe = False

    for i in range(len(line)):
        internal_safe_check = True
        np_line = np.array(line)
        new_line = np_line[np.arange(len(line))!=i]
        sign_check = new_line[0] - new_line[1]
        if abs(sign_check) > max_level_change or sign_check == 0:
            internal_safe_check = False
            continue
        for i in range(1, len(new_line)-1):
            diff = new_line[i] - new_line[i+1]
            if abs(diff) > max_level_change or diff * sign_check <= 0 or diff == 0:
                internal_safe_check = False
                continue
        if internal_safe_check:
            safe = True
            break
    num_safe += safe

print(num_safe)