def second_best(prices):
    """Return second best price.

    >>> second_best([2, 1, 3])
    2
    """
    if not prices:
        raise ValueError('second_best on empty prices')

    first = second = prices[0]
    for price in prices[1:]:
        if price > first:
            first, second = price, first
        elif price > second:
            second = price
    return second
