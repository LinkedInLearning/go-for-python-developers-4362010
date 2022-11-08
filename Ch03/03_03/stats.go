package main

import "fmt"

func Sum(values []float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}

func Mean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Mean of empty slice")
	}

	return Sum(values) / float64(len(values)), nil
}

func main() {
	values := []float64{2, 4, 8}
	m, err := Mean(values)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(m)
}
