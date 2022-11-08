package rtb

import "fmt"

func SecondBest(prices []int) (int, error) {
	if len(prices) == 0 {
		return 0, fmt.Errorf("SecondBest of empty prices")
	}

	first, second := prices[0], prices[0]
	for _, p := range prices[1:] {
		switch {
		case p > first:
			first, second = p, first
		case p > second:
			second = p
		}
	}

	return second, nil
}
