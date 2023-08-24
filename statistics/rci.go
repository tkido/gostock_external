package statistics

import (
	"errors"
	"sort"
)

// Rci is Rank Correlation Index
func Rci(arr interface{}) (float64, error) {
	order, err := getOrder(arr)
	if err != nil {
		return 0, err
	}
	return rci(order), nil
}

// RciInt is Rci for int
func RciInt(arr []int) float64 {
	return rci(getOrderInt(arr))
}

// RciFloat is Rci for float64
func RciFloat(arr []float64) float64 {
	return rci(getOrderFloat(arr))
}

func rci(order []int) float64 {
	var sum int
	for i, num := range order {
		sum += square(i - num)
	}
	d := float64(sum)
	s := float64(len(order))
	return 1 - (6 * d / (s * (s*s - 1)))
}
func square(x int) int { return x * x }

func getOrder(arr interface{}) ([]int, error) {
	switch a := arr.(type) {
	case []int:
		return getOrderInt(a), nil
	case []float64:
		return getOrderFloat(a), nil
	default:
		return nil, errors.New("getOrder(): invalid type")
	}
}

type iPair struct {
	First int
	Last  int
}

func getOrderInt(arr []int) []int {
	pairs := make([]iPair, len(arr))
	for i, num := range arr {
		pairs[i] = iPair{i, num}
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].Last < pairs[j].Last
	})
	rsts := make([]int, len(pairs))
	for i, p := range pairs {
		rsts[i] = p.First
	}
	return rsts
}

type ifPair struct {
	First int
	Last  float64
}

func getOrderFloat(arr []float64) []int {
	pairs := make([]ifPair, len(arr))
	for i, num := range arr {
		pairs[i] = ifPair{i, num}
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].Last < pairs[j].Last
	})
	rsts := make([]int, len(pairs))
	for i, p := range pairs {
		rsts[i] = p.First
	}
	return rsts
}
