# Print character and frequency in percent of Rumi's poem
# Print in frequency descending order
# Assume ASCII text, ignore white space
from operator import itemgetter

poem = '''
those who do not feel this love
pulling them like a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let them sleep
'''

counts = {}
count = 0
for c in poem:
    if c.isspace():
        continue
    counts[c] = counts.get(c, 0) + 1
    count += 1

for c, frequency in sorted(counts.items(), key=itemgetter(1), reverse=True):
    frequency = frequency / count * 100
    print(f'{c}: {frequency:.2f}')
