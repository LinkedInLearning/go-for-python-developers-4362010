from queue import Queue

queue = Queue()
queue.put(99)
val = queue.get()
print(f'got {val}')
