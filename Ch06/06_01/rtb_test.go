package rtb

import (
	"fmt"
	"testing"
)

func TestSecondBest(t *testing.T) {
	prices := []int{1, 0, 2, 3}
	p, err := SecondBest(prices)
	if err != nil {
		t.Fatal(err)
	}
	expected := 2
	if p != expected {
		t.Fatalf("expected %d, got %d", p, expected)
	}
}

var secondBestCases = []struct {
	prices   []int
	expected int
}{
	{[]int{1}, 1},
	{[]int{1, 2, 3, 4, 2}, 3},
}

func TestSecondBestTable(t *testing.T) {
	for _, tc := range secondBestCases {
		name := fmt.Sprintf("%v", tc.prices)
		t.Run(name, func(t *testing.T) {
			p, err := SecondBest(tc.prices)
			if err != nil {
				t.Fatal(err)
			}
			if p != tc.expected {
				t.Fatalf("expected %d, got %d", p, tc.expected)

			}
		})
	}
}
