import pytest
from rtb import second_best


def test_rtb():
    prices = [10, 20, 30]
    sb = second_best(prices)
    assert sb == 20


rtb_cases = [
    # prices, expected
    [[10, 20, 30], 20],
    [[1, 2, 3, 1, 2], 2],
]


@pytest.mark.parametrize('prices, expected', rtb_cases)
def test_rtb_table(prices, expected):
    sb = second_best(prices)
    assert sb == expected
