go_rng --- A pseudo-random number generator written in Golang
=============================================================

Author: Leesper

Inspired by:
------------

1.[StdRandom.java](http://introcs.cs.princeton.edu/java/stdlib/StdRandom.java.html)

2.[Numerical Recipes](http://www.nr.com/)

3.[Random number generation](http://en.wikipedia.org/wiki/Random_number_generation)

4.[Quantile function](http://en.wikipedia.org/wiki/Quantile_function)

5.[Monte Carlo method](http://en.wikipedia.org/wiki/Monte_Carlo_method)

6.[Pseudo-random number sampling](http://en.wikipedia.org/wiki/Pseudo-random_number_sampling)

7.[Inverse transform sampling](http://en.wikipedia.org/wiki/Inverse_transform_sampling)

Supported Distributions and Functionalities:
-------------------------------------------
### Uniform Distribution
### Bernoulli Distribution
### Chi-Squared Distribution
### Gamma Distribution
### Beta Distribution
### Fisher's F Distribution
### Cauchy Distribution
### Weibull Distribution
### Pareto Distribution
### Log Normal Distribution
### Exponential Distribution
### Student's t-Distribution
### Binomial Distribution
### Bernoulli Distribution
### Poisson Distribution
### Geometric Distribution
### Gaussian Distribution
### Logistic Distribution

Usage
-----

        package rng
        import "rng"

### Uniform Distribution

        1) struct UniformGenerator <br />
        UniformGenerator is a random number generator for uniform
        distribution. The zero value is invalid, use NewUniformGenerator to
        create a generator <br />
        
        2) func NewUniformGenerator(seed int64) *UniformGenerator <br />
        NewUniformGenerator returns a uniform-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: urng
        := rng.NewUniformGenerator(time.Now().UnixNano()) <br />
        
        3) func (ung UniformGenerator) Float32() float32 <br />
        Float32 returns a random float32 in [0.0, 1.0) <br />
        
        4) func (ung UniformGenerator) Float32Range(a, b float32) float32 <br />
        Float32Range returns a random float32 in [a, b) <br />
        
        5) func (ung UniformGenerator) Float32n(n float32) float32 <br />
        Float32n returns a random float32 in [0.0, n) <br />
        
        6) func (ung UniformGenerator) Float64() float64 <br />
        Float64 returns a random float64 in [0.0, 1.0) <br />
        
        7) func (ung UniformGenerator) Float64Range(a, b float64) float64 <br />
        Float32Range returns a random float32 in [a, b) <br />
        
        8) func (ung UniformGenerator) Float64n(n float64) float64 <br />
        Float64n returns a random float64 in [0.0, n) <br />
        
        9) func (ung UniformGenerator) Int32() int32 <br />
        Int32 returns a random uint32 <br />
        
        10) func (ung UniformGenerator) Int32Range(a, b int32) int32 <br />
        Int32Range returns a random uint32 in [a, b) <br />
        
        11) func (ung UniformGenerator) Int32n(n int32) int32 <br />
        Int32n returns a random uint32 in [0, n) <br />
        
        12) func (ung UniformGenerator) Int64() int64 <br />
        Int64 returns a random uint64 <br />
        
        13) func (ung UniformGenerator) Int64Range(a, b int64) int64 <br />
        Int64Range returns a random uint64 in [a, b) <br />
        
        14) func (ung UniformGenerator) Int64n(n int64) int64 <br />
        Int64n returns a random uint64 in [0, n) <br />
        
        15) func (ung UniformGenerator) Shuffle(arr []interface{}) <br />
        Shuffle rearrange the elements of an array in random order <br />
        
        16) func (ung UniformGenerator) ShuffleRange(arr []interface{}, low, high int) <br />
        Shuffle rearrange the elements of the subarray[low..high] in random order <br />

### Bernoulli Distribution

        1) struct BernoulliGenerator <br />
        UniformGenerator is a random number generator for uniform distribution.
        The zero value is invalid, use NewBernoulliGenerator to create a
        generator <br />
        
        2) func NewBernoulliGenerator(seed int64) *BernoulliGenerator <br />
        NewBernoulliGenerator returns a bernoulli-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: urng
        := rng.NewBernoulliGenerator(time.Now().UnixNano()) <br />
        
        3) func (beng BernoulliGenerator) Bernoulli() bool <br />
        bernoulli returns a bool, which is true with probablity 0.5 <br />
        
        4) func (beng BernoulliGenerator) Bernoulli_P(p float32) bool <br />
        bernoulli_P returns a bool, which is true with probablity p <br />

### Binomial Distribution

        1) struct BinomialGenerator <br />
        BinomialGenerator is a random number generator for binomial
        distribution. The zero value is invalid, use NewBinomialGenerator to
        create a generator <br />
        
        2) func NewBinomialGenerator(seed int64) *BinomialGenerator <br />
        NewBinomialGenerator returns a binomial-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: urng
        := rng.NewBinomialGenerator(time.Now().UnixNano()) <br />
        
        3) func (bing BinomialGenerator) Binomial(n int64, p float32) int64 <br />
        Binomial returns a random number X ~ binomial(n, p) <br />
        
### Geometric Distribution
        
        1) struct GeometricGenerator <br />
        GeometricGenerator is a random number generator for geometric
        distribution. The zero value is invalid, use NewGeometryGenerator to
        create a generator <br />
        
        2) func NewGeometricGenerator(seed int64) *GeometricGenerator <br />
        NewGeometricGenerator returns a geometric-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: urng
        := rng.NewGeometricGenerator(time.Now().UnixNano()) <br />
        
        3) func (grng GeometricGenerator) Geometric(p float64) int64 <br />
        Geometric returns a random number X ~ binomial(n, p) <br />

### Poisson Distribution

        1) struct PoissonGenerator <br />
        PoissonGenerator is a random number generator for possion distribution.
        The zero value is invalid, use NewPoissonGenerator to create a generator <br />
        
        2) func NewPoissonGenerator(seed int64) *PoissonGenerator <br />
        NewPoissonGenerator returns a possion-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: prng
        := rng.NewPoissonGenerator(time.Now().UnixNano()) <br />
        
        3) func (prng PoissonGenerator) Possion(lambda float64) int64
        Poisson returns a random number of possion distribution
        
### Exponential Distribution

        1) struct ExpGenerator <br />
        ExpGenerator is a random number generator for exponential distribution.
        The zero value is invalid, use NewExpGenerator to create a generator <br />
        
        2) func NewExpGenerator(seed int64) *ExpGenerator <br />
        NewExpGenerator returns a exponential-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: erng
        := rng.NewExpGenerator(time.Now().UnixNano()) <br />
        
        3) func (erng ExpGenerator) Exp(lambda float64) float64 <br />
        Exp returns a random number of exponential distribution <br />

### Cauchy Distribution

        1) struct CauchyGenerator <br />
        CauchyGenerator is a random number generator for cauchy distribution.
        The zero value is invalid, use NewCauchyGenerator to create a generator <br />
        
        2) func NewCauchyGenerator(seed int64) *CauchyGenerator <br />
        NewCauchyGenerator returns a cauchy-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: crng
        := rng.NewCauchyGenerator(time.Now().UnixNano()) <br />
        
        3) func (crng CauchyGenerator) Cauchy(x0, gamma float64) float64 <br />
        Cauchy returns a random number of cauchy distribution <br />
        
        4) func (crng CauchyGenerator) StandardCauchy() float64 <br />
        StandardCauchy() returns a random number of standard cauchy distribution (x0 = 0.0, gamma = 1.0) <br />
        
### Logistic Distribution

        1) struct LogisticGenerator <br />
        LogisticGenerator is a random number generator for cauchy distribution.
        The zero value is invalid, use NewLogisticGenerator to create a generator <br />
        
        2) func NewLogisticGenerator(seed int64) *LogisticGenerator <br />
        NewLogisticGenerator returns a logistic-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: lrng
        := rng.NewLogisticGenerator(time.Now().UnixNano()) <br />
        
        3) func (lrng LogisticGenerator) Logistic(mu, s float64) float64 <br />
        Logistic returns a random number of logistic distribution <br />

### Gaussian Distribution

        1) struct GaussianGenerator <br />
        GaussianGenerator is a random number generator for gaussian
        distribution. The zero value is invalid, use NewGaussianGenerator to
        create a generator <br />
        
        2) func NewGaussianGenerator(seed int64) *GaussianGenerator <br />
        NewGaussianGenerator returns a gaussian-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: crng
        := rng.NewGaussianGenerator(time.Now().UnixNano()) <br />
        
        3) func (grng GaussianGenerator) Gaussian(mean, stddev float64) float64 <br />
        Gaussian returns a random number of gaussian distribution Gauss(mean, stddev^2) <br />
        
        4)func (grng GaussianGenerator) StdGaussian() float64 <br />
        StdGaussian returns a random number of standard gaussian distribution <br />

### Pareto Distribution (type I)

        1) struct ParetoGenerator <br />
        ParetoGenerator is a random number generator for type I pareto
        distribution. The zero value is invalid, use NewParetoGenerator to
        create a generator <br />
        
        2) func NewParetoGenerator(seed int64) *ParetoGenerator <br />
        NewParetoGenerator returns a type I pareto-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: crng
        := rng.NewParetoGenerator(time.Now().UnixNano()) <br />
        
        3) func (prng ParetoGenerator) Pareto(alpha float64) float64 <br />
        Pareto returns a random number of type I pareto distribution (alpha > 0,0) <br />

### Weibull Distribution
        
        1) struct WeibullGenerator <br /> 
        WeibullGenerator is a random number generator for type weibull
        distribution. The zero value is invalid, use NewWeibullGenerator to
        create a generator <br />
        
        2) func NewWeibullGenerator(seed int64) *WeibullGenerator <br />
        NewWeibullGenerator returns a weibull-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: wrng
        := rng.NewWeibullGenerator(time.Now().UnixNano()) <br />

        3) func (wrng WeibullGenerator) Weibull(lambda, k float64) float64
        Weibull returns a random number of weibull distribution (lambda > 0.0 and k > 0.0) <br />

### Gamma Distribution
        1) struct GammaGenerator <br />
        GammaGenerator is a random number generator for gamma distribution. The
        zero value is invalid, use NewGammaGenerator to create a generator <br />

        2) func NewGammaGenerator(seed int64) *GammaGenerator <br />
        NewGammaGenerator returns a type I pareto-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: grng
        := rng.NewGammaGenerator(time.Now().UnixNano()) <br />
        
        3) func (grng GammaGenerator) Gamma(alpha, beta float64) float64 <br />
        Gamma returns a random number of gamma distribution (alpha > 0.0 and beta > 0.0) <br />

### Lognormal Distribution

        1) struct LognormalGenerator <br />
        LognormalGenerator is a random number generator for lognormal
        distribution. The zero value is invalid, use NewLognormalGenerator to
        create a generator <br />
        
        2) func NewLognormalGenerator(seed int64) *LognormalGenerator <br />
        NewLognormalGenerator returns a lognormal-distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: crng
        := rng.NewLognormalGenerator(time.Now().UnixNano()) <br />
        
        3) func (lnng LognormalGenerator) Lognormal(mean, stddev float64) float64 <br />
        Lognormal return a random number of lognormal distribution <br />

### Beta Distribution

        1) struct BetaGenerator struct <br />
        BetaGenerator is a random number generator for beta distribution. The
        zero value is invalid, use NewBetaGenerator to create a generator <br />
        
        2) func NewBetaGenerator(seed int64) *BetaGenerator <br />
        NewBetaGenerator returns a beta distribution generator it is recommended
        using time.Now().UnixNano() as the seed, for example: brng := 
        rng.NewBetaGenerator(time.Now().UnixNano()) <br />
        
        3) func (brng BetaGenerator) Beta(alpha, beta float64) float64 <br />
        Beta returns a random number of beta distribution (alpha > 0.0 and beta > 0.0) <br />

### Chi-Squared Distribution

        1) struct ChiSquaredGenerator <br />
        ChiSquaredGenerator is a random number generator for chi-squared
        distribution. The zero value is invalid, use NewChiSquaredGenerator to
        create a generator <br />
        
        2) func NewChiSquaredGenerator(seed int64) *ChiSquaredGenerator <br />
        NewChiSquaredGenerator returns a chi-squared distribution generator it
        is recommended using time.Now().UnixNano() as the seed, for example:
        crng := rng.NewChiSquaredGenerator(time.Now().UnixNano()) <br />
        
        3) func (crng ChiSquaredGenerator) ChiSquared(freedom int64) float64 <br />
        ChiSquared returns a random number of chi-squared distribution <br />

### Student's t-distribution

        1) struct StudentTGenerator <br />
        StudentTGenerator is a random number generator for student-t
        distribution. The zero value is invalid, use NewStudentTGenerator to
        create a generator <br />
        
        2) func NewStudentTGenerator(seed int64) *StudentTGenerator <br />
        NewStudentTGenerator returns a student-t distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: stng
        := rng.NewStudentTGenerator(time.Now().UnixNano()) <br />
        
        3) func (stng StudentTGenerator) Student(freedom int64) float64 <br />
        Student returns a random number of student-t distribution (freedom > 0.0) <br />

### Fisher's F Distribution

        1) struct FisherFGenerator <br />
        FisherFGenerator is a random number generator for Fisher's F
        distribution. The zero value is invalid, use NewFisherFGenerator to
        create a generator <br />
        
        2) func NewFisherFGenerator(seed int64) *FisherFGenerator <br />
        NewFisherFGenerator returns a Fisher's F distribution generator it is
        recommended using time.Now().UnixNano() as the seed, for example: frng
        := rng.NewFisherFGenerator(time.Now().UnixNano()) <br />
        
        3) func (frng FisherFGenerator) Fisher(d1, d2 int64) float64 <br />
        Fisher returns a random number of Fisher's F distribution (d1 > 0 and d2 > 0) <br />
