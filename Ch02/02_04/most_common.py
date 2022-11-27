# What is the most common word in Rumi's poem?

poem = '''
those who do not feel this love
pulling them like a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let them sleep
'''

frequency = {}
for word in poem.split():
    frequency[word] = frequency.get(word, 0) + 1

most_common = max(frequency, key=frequency.get)
print(most_common)
