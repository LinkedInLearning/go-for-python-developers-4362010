from pathlib import Path
from queue import Queue
from threading import Thread


def worker(file_name, queue):
    size = Path(file_name).stat().st_size
    queue.put((file_name, size))


files = [
    'Ch05/challenge/files/a.txt',
    'Ch05/challenge/files/b.txt',
    'Ch05/challenge/files/c.txt',
    'Ch05/challenge/files/d.txt',
    'Ch05/challenge/files/e.txt',
]

queue = Queue()
for file_name in files:
    Thread(target=worker, args=(file_name, queue), daemon=True).start()

for _ in files:
    file_name, size = queue.get()
    print(f'{file_name} -> {size}')
