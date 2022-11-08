def collatz_step(n):
    """Compute one step in Collatz conjecture"""
    if n % 2 == 0:
        return n // 2
    return n * 3 + 1


if __name__ == '__main__':
    print(collatz_step(4))
    print(collatz_step(5))
