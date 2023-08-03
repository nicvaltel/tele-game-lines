package utils

import (
	"errors"
	"math/rand"
	"time"
)

func DifferentRandomNumbers(min, max int, count int) ([]int, error) {
	if count > max-min+1 {
		return nil, errors.New("count should be less than or equal to range")
	}

	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, count)
	used := make(map[int]bool)
	randNum := rand.Intn(max-min+1) + min

	for i := 0; i < count; i++ {
		for used[randNum] {
			randNum = rand.Intn(max-min+1) + min
		}
		numbers[i] = randNum
		used[randNum] = true
	}
	return numbers, nil
}
