go_rng --- A pseudo-random number generator written in Golang

Author: Leesper

Inspired by:

1. StdRandom.java --- "Algorithms 4ed" by Prof Robert Sedgewick http://introcs.cs.princeton.edu/java/stdlib/StdRandom.java.html

2. Chapter 7 Random Numbers --- "Numerical Recipes" by William.H http://www.nr.com/

3. Random number generation --- http://en.wikipedia.org/wiki/Random_number_generation

4. Quantile function --- http://en.wikipedia.org/wiki/Quantile_function

5. Monte Carlo method --- http://en.wikipedia.org/wiki/Monte_Carlo_method

6. Pseudo-random number sampling --- http://en.wikipedia.org/wiki/Pseudo-random_number_sampling

7. Inverse transform sampling --- http://en.wikipedia.org/wiki/Inverse_transform_sampling

Supported Distributions and Functionalities:

1. uniform distribution

package rng
    import "rng"
TYPES

type UniformGenerator struct {
    // contains filtered or unexported fields
}
    UniformGenerator is a random number generator for uniform
    distribution. The zero value is invalid, use NewUniformGenerator to
    create a generator

func NewUniformGenerator(seed int64) *UniformGenerator
    NewUniformGenerator returns a uniform-distribution generator it is
    recommended using time.Now().UnixNano() as the seed, for example: urng
    := rng.NewUniformGenerator(time.Now().UnixNano())

func (ung UniformGenerator) Float32() float32
    Float32 returns a random float32 in [0.0, 1.0)

func (ung UniformGenerator) Float32Range(a, b float32) float32
    Float32Range returns a random float32 in [a, b)

func (ung UniformGenerator) Float32n(n float32) float32
    Float32n returns a random float32 in [0.0, n)

func (ung UniformGenerator) Float64() float64
    Float64 returns a random float64 in [0.0, 1.0)

func (ung UniformGenerator) Float64Range(a, b float64) float64
    Float32Range returns a random float32 in [a, b)

func (ung UniformGenerator) Float64n(n float64) float64
    Float64n returns a random float64 in [0.0, n)

func (ung UniformGenerator) Int32() int32
    Int32 returns a random uint32

func (ung UniformGenerator) Int32Range(a, b int32) int32
    Int32Range returns a random uint32 in [a, b)

func (ung UniformGenerator) Int32n(n int32) int32
    Int32n returns a random uint32 in [0, n)

func (ung UniformGenerator) Int64() int64
    Int64 returns a random uint64

func (ung UniformGenerator) Int64Range(a, b int64) int64
    Int64Range returns a random uint64 in [a, b)

func (ung UniformGenerator) Int64n(n int64) int64
    Int64n returns a random uint64 in [0, n)

bernoulli
binomial
geometric
poisson
exponential
gamma
normal
chi squared
cauchy
pareto
logistic
weibull
negative binomial
extreme value
lognormal
fisher f
student t
discrete
shuffle
