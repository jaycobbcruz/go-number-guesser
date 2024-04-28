package numberguesser

import (
	"errors"
	"math"
)

const (
	MaxLoop = 400
)

type Formula func(float64) float64

func Guess(function Formula, expectedResult float64) (float64, error) {
	guess := 1.0
	low := 1.0
	i := 0
	for MaxLoop > i {
		i++
		result := function(guess)
		if result == expectedResult {
			return guess, nil
		} else if result < expectedResult {
			low = guess
			guess = guess * 2
		} else if math.Mod(result, expectedResult) == 0 {
			return findNumber(function, 0, expectedResult, result/expectedResult, 0.0)
		} else {
			return findNumber(function, 0, expectedResult, low, low)
		}
	}

	if MaxLoop > i {
		return 0.0, errors.New("max loop exceeded")
	}
	return guess, nil
}

func findNumber(
	function Formula,
	maxLoopCounter float64,
	expected float64,
	guess float64,
	offset float64) (float64, error) {

	if maxLoopCounter > MaxLoop {
		return 0.0, errors.New("reached max loop count")
	}

	newOffset := offset / 2
	newGuess := guess + newOffset
	result := function(newGuess)

	if expected > result {
		return findNumber(function, maxLoopCounter+1, expected, newGuess, math.Abs(newOffset))
	} else if expected < result {
		return findNumber(function, maxLoopCounter+1, expected, newGuess, offset*-1)
	} else {
		return newGuess, nil
	}
}
