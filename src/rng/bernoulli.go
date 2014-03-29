package rng

import (
	"fmt"
)

// UniformGenerator is a random number generator for uniform distribution.
// The zero value is invalid, use NewBernoulliGenerator to create a generator
type BernoulliGenerator struct {
	uniform		*UniformGenerator
}

// NewBernoulliGenerator returns a bernoulli-distribution generator
// it is recommended using time.Now().UnixNano() as the seed, for example:
// urng := rng.NewBernoulliGenerator(time.Now().UnixNano())
func NewBernoulliGenerator(seed int64) *BernoulliGenerator {
	urng := NewUniformGenerator(seed)
	return &BernoulliGenerator{ urng }
}

// bernoulli returns a bool, which is true with probablity 0.5
func (beng BernoulliGenerator) Bernoulli() bool {
	return beng.Bernoulli_P(0.5)
}

// bernoulli_P returns a bool, which is true with probablity p
func (beng BernoulliGenerator) Bernoulli_P(p float32) bool {
	if !(0.0 <= p && p <= 1.0) {
		panic(fmt.Sprintf("Invalid probability", p))
	}
	return beng.uniform.Float32() < p
}
