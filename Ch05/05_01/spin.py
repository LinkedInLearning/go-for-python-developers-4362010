from threading import Thread
from time import sleep


def worker(i):
    sleep(0.1)
    print(f'thread {i}')


for i in range(5):
    Thread(target=worker, args=(i,), daemon=True).start()
print('main')
