# Print a banner with length of 8 for the string 'Hi ☺'
# Should output
#   Hi ☺
# --------

message = 'Hi ☺'
width = 8
padding = (width - len(message)) // 2
print(' ' * padding, end='')
print(message)
print('-' * width)
