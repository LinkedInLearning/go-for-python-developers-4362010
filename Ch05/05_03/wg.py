from queue import Queue
from threading import Thread
from time import sleep


def worker(id, queue):
    while True:
        item = queue.get()
        sleep(0.1)
        print(f'{id} finished {item}')
        queue.task_done()


if __name__ == '__main__':
    queue = Queue()
    for id in range(3):
        Thread(target=worker, args=(id, queue), daemon=True).start()

    for msg in ['A', 'B', 'C', 'D', 'E', 'F']:
        queue.put(msg)

    queue.join()
    print('all jobs done')
