# What is the average of all the numbers between 1 to 100 that divide either
# by 3 or 5?

count, total = 0, 0
for n in range(1, 101):
    if n % 3 == 0 or n % 5 == 0:
        count += 1
        total += n

print(total/count)
