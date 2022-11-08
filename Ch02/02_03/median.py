# Calculate the median of a random list of numbers

from random import random

size = 11
numbers = [random() * 100 for _ in range(size)]

numbers.sort()
i = len(numbers) // 2
if len(numbers) % 2 == 1:
    median = numbers[i]
else:
    median = (numbers[i-1] + numbers[i]) / 2
print(median)
