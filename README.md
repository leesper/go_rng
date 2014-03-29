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

package rng
    
import "rng"

Uniform Distribution

1) struct UniformGenerator 

UniformGenerator is a random number generator for uniform
distribution. The zero value is invalid, use NewUniformGenerator to
create a generator

2) func NewUniformGenerator(seed int64) *UniformGenerator

NewUniformGenerator returns a uniform-distribution generator it is
recommended using time.Now().UnixNano() as the seed, for example: urng
:= rng.NewUniformGenerator(time.Now().UnixNano())

3) func (ung UniformGenerator) Float32() float32

Float32 returns a random float32 in [0.0, 1.0)

4) func (ung UniformGenerator) Float32Range(a, b float32) float32
    
Float32Range returns a random float32 in [a, b)

5) func (ung UniformGenerator) Float32n(n float32) float32
    
Float32n returns a random float32 in [0.0, n)

6) func (ung UniformGenerator) Float64() float64
    
Float64 returns a random float64 in [0.0, 1.0)

7) func (ung UniformGenerator) Float64Range(a, b float64) float64
    
Float32Range returns a random float32 in [a, b)

8) func (ung UniformGenerator) Float64n(n float64) float64
    
Float64n returns a random float64 in [0.0, n)

9) func (ung UniformGenerator) Int32() int32
    
Int32 returns a random uint32

10) func (ung UniformGenerator) Int32Range(a, b int32) int32
    
Int32Range returns a random uint32 in [a, b)

11) func (ung UniformGenerator) Int32n(n int32) int32
    
Int32n returns a random uint32 in [0, n)

12) func (ung UniformGenerator) Int64() int64
    
Int64 returns a random uint64

13) func (ung UniformGenerator) Int64Range(a, b int64) int64
    
Int64Range returns a random uint64 in [a, b)

14) func (ung UniformGenerator) Int64n(n int64) int64
    
Int64n returns a random uint64 in [0, n)

Bernoulli Distribution

1) struct BernoulliGenerator

UniformGenerator is a random number generator for uniform distribution.
The zero value is invalid, use NewBernoulliGenerator to create a
generator

2) func NewBernoulliGenerator(seed int64) *BernoulliGenerator

NewBernoulliGenerator returns a bernoulli-distribution generator it is
recommended using time.Now().UnixNano() as the seed, for example: urng
:= rng.NewBernoulliGenerator(time.Now().UnixNano())

3) func (beng BernoulliGenerator) Bernoulli() bool

bernoulli returns a bool, which is true with probablity 0.5

4) func (beng BernoulliGenerator) Bernoulli_P(p float32) bool

bernoulli_P returns a bool, which is true with probablity p

Binomial Distribution

1) struct BinomialGenerator

BinomialGenerator is a random number generator for binomial
distribution. The zero value is invalid, use NewBinomialGenerator to
create a generator

2) func NewBinomialGenerator(seed int64) *BinomialGenerator

NewBinomialGenerator returns a binomial-distribution generator it is
recommended using time.Now().UnixNano() as the seed, for example: urng
:= rng.NewBinomialGenerator(time.Now().UnixNano())

3) func (bing BinomialGenerator) Binomial(n int64, p float32) int64

Binomial returns a random number X ~ binomial(n, p)

Geometric Distribution

1) struct GeometricGenerator

GeometricGenerator is a random number generator for geometric
distribution. The zero value is invalid, use NewGeometryGenerator to
create a generator

2) func NewGeometricGenerator(seed int64) *GeometricGenerator

NewGeometricGenerator returns a geometric-distribution generator it is
recommended using time.Now().UnixNano() as the seed, for example: urng
:= rng.NewGeometricGenerator(time.Now().UnixNano())

3) func (grng GeometricGenerator) Geometric(p float64) int64

Geometric returns a random number X ~ binomial(n, p)

TODO:

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
