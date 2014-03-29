// Package rng implements a series of pseudo-random number generator 
// based on a variety of common probability distributions
// Author: Leesper
// Email: pascal7718@gmail.com 394683518@qq.com

package rng

import (
	"fmt"
)

// BinomialGenerator is a random number generator for binomial distribution.
// The zero value is invalid, use NewBinomialGenerator to create a generator
type BinomialGenerator struct {
	uniform		*UniformGenerator
}

// NewBinomialGenerator returns a binomial-distribution generator
// it is recommended using time.Now().UnixNano() as the seed, for example:
// urng := rng.NewBinomialGenerator(time.Now().UnixNano())
func NewBinomialGenerator(seed int64) *BinomialGenerator {
	urng := NewUniformGenerator(seed)
	return &BinomialGenerator{ urng }
}

// Binomial returns a random number X ~ binomial(n, p)
func (bing BinomialGenerator) Binomial(n int64, p float32) int64 {
	if !(0.0 <= p && p <= 1.0) {
		panic(fmt.Sprintf("Invalid probability p: ", p))
	}
	if n <= 0 {
		panic(fmt.Sprintf("Invalid parameter n: ", n))
	}
	var i, result int64
	for i = 0; i < n; i++ {
		if bing.uniform.Float32() < p {
			result++
		}
	}
	return result
}