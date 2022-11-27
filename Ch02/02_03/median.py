# Calculate the median of a list of numbers


numbers = [2, 1, 3]

numbers.sort()
i = len(numbers) // 2
if len(numbers) % 2 == 1:
    median = numbers[i]
else:
    median = (numbers[i-1] + numbers[i]) / 2
print(median)
