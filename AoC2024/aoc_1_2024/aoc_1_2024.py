import numpy as np
from collections import Counter
lines = []

with open('aoc_1_2024.txt', 'r') as f:
    data = np.array([x.strip().split() for x in f.readlines()])
data = data.astype('int32')
fixed_order_data = np.zeros_like(data).astype('int32')

fixed_order_data[:,0] = sorted(data[:,0])
fixed_order_data[:,1] = sorted(data[:,1])
sum = np.sum(list(map(lambda x: abs(x[0] - x[1]), fixed_order_data)))

left_dict = Counter(fixed_order_data[:,0])
right_dict = Counter(fixed_order_data[:,1])

similarity = 0

for key in left_dict:
    similarity +=  key * left_dict.get(key) * right_dict.get(key, 0)

print(sum)

print(similarity)