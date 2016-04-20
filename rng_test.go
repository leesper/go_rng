package rng

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestDirichletGenerator(t *testing.T) {
	fmt.Println("=====Testing for DirichletGenerator begin=====")
	epsilon := 0.001
	drng := NewDirichletGenerator(time.Now().UnixNano())
	fmt.Print("Dirichlet({1.0, 42.42, 244.24}): ")
	alphas := [...]float64{1.0, 42.42, 244.24}
	values := drng.Dirichlet(alphas[:])
	sum := SumFloat64Slice(values)
	fmt.Printf("Sum: %f\n", sum)
	if math.Abs(1.0-sum) > epsilon {
		t.Fatalf("Bad sum: %f", sum)
	}
	fmt.Print("SymmetricDirichlet(42, 4242): ")
	values = drng.SymmetricDirichlet(42, 4242)
	sum = SumFloat64Slice(values)
	fmt.Printf("Sum: %f\n", sum)
	if math.Abs(1.0-sum) > epsilon {
		t.Fatalf("Bad sum: %f", sum)
	}
	fmt.Print("FlatDirichlet(8): ")
	values = drng.FlatDirichlet(8)
	sum = SumFloat64Slice(values)
	fmt.Printf("Sum: %f (values: %v)\n", sum, values)
	if math.Abs(1.0-sum) > epsilon {
		t.Fatalf("Bad sum: %f", sum)
	}
	fmt.Println("=====Testing for DirichletGenerator end=====")
	fmt.Println()
}

func TestFisherFGenerator(t *testing.T) {
	fmt.Println("=====Testing for FisherFGenerator begin=====")
	frng := NewFisherFGenerator(time.Now().UnixNano())
	fmt.Println("Fisher(100, 100): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(frng.Fisher(100, 100))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for FisherFGenerator end=====")
	fmt.Println()
}

func TestStudentTGenerator(t *testing.T) {
	fmt.Println("=====Testing for StudentTGenerator begin=====")
	stng := NewStudentTGenerator(time.Now().UnixNano())
	fmt.Println("StudentT(5): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(stng.Student(5))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for StudentTGenerator end=====")
	fmt.Println()
}

func TestChiSquaredGenerator(t *testing.T) {
	fmt.Println("=====Testing for ChiSquaredGenerator begin=====")
	crng := NewChiSquaredGenerator(time.Now().UnixNano())
	fmt.Println("ChiSquared(1): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(crng.ChiSquared(1))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for ChiSquaredGenerator end=====")
	fmt.Println()
}

func TestBetaGenerator(t *testing.T) {
	// FIXME: I have no idea whether it is right :(
	fmt.Println("=====Testing for BetaGenerator begin=====")
	brng := NewBetaGenerator(time.Now().UnixNano())
	fmt.Println("Beta(2.0, 2.0): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(brng.Beta(2.0, 2.0))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for BetaGenerator end=====")
	fmt.Println()
}

func TestLognormalGenerator(t *testing.T) {
	fmt.Println("=====Testing for LognormalGenerator begin=====")
	lnng := NewLognormalGenerator(time.Now().UnixNano())
	fmt.Println("Lognormal(0.0, 0.5): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(lnng.Lognormal(0.0, 0.5))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for LognormalGenerator end=====")
	fmt.Println()
}

func TestGammaGenerator(t *testing.T) {
	fmt.Println("=====Testing for GammaGenerator begin=====")
	grng := NewGammaGenerator(time.Now().UnixNano())
	fmt.Println("Gamma(9.0, 0.5): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(grng.Gamma(9.0, 0.5))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for GammaGenerator end=====")
	fmt.Println()
}

func TestWeibullGenerator(t *testing.T) {
	fmt.Println("=====Testing for WeibullGenerator begin=====")
	wrng := NewWeibullGenerator(time.Now().UnixNano())
	fmt.Println("Weibull(1.0, 0.5): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(wrng.Weibull(1.0, 0.5))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for WeibullGenerator end=====")
	fmt.Println()
}

func TestParetoGenerator(t *testing.T) {
	fmt.Println("=====Testing for ParetoGenerator begin=====")
	prng := NewParetoGenerator(time.Now().UnixNano())
	fmt.Println("Pareto(2.0): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(prng.Pareto(2.0))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for ParetoGenerator end=====")
	fmt.Println()
}

func TestGaussianGenerator(t *testing.T) {
	fmt.Println("=====Testing for GaussianGenerator begin=====")
	grng := NewGaussianGenerator(time.Now().UnixNano())
	fmt.Println("Gaussian(5.0, 2.0): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(grng.Gaussian(5.0, 2.0))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/200))
	}

	fmt.Println("=====Testing for GaussianGenerator end=====")
	fmt.Println()
}

func TestLogisticGenerator(t *testing.T) {
	fmt.Println("=====Testing for LogisticGenerator begin=====")
	lrng := NewLogisticGenerator(time.Now().UnixNano())
	fmt.Println("Logistic(6.0, 2.0): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(lrng.Logistic(6.0, 2.0))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/100))
	}

	fmt.Println("=====Testing for LogisticGenerator end=====")
	fmt.Println()
}

func TestCauchyGenerator(t *testing.T) {
	fmt.Println("=====Testing for CauchyGenerator begin=====")
	crng := NewCauchyGenerator(time.Now().UnixNano())
	fmt.Println("Cauchy(-2.0, 1.0): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(crng.Cauchy(-2.0, 1.0))]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/100))
	}

	fmt.Println("=====Testing for CauchyGenerator end=====")
	fmt.Println()
}

func TestExpGenerator(t *testing.T) {
	fmt.Println("=====Testing for ExpGenerator begin=====")
	erng := NewExpGenerator(time.Now().UnixNano())
	fmt.Println("Poisson(1): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[int64(erng.Exp(1)*2)]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%.1f-%.1f:\t%s\n", float32(key)/2.0, float32(key+1)/2.0,
			strings.Repeat("*", hist[key]/100))
	}

	fmt.Println("=====Testing for ExpGenerator end=====")
	fmt.Println()
}

func TestPoissonGenerator(t *testing.T) {
	fmt.Println("=====Testing for PoissonGenerator begin=====")
	prng := NewPoissonGenerator(time.Now().UnixNano())
	fmt.Println("Poisson(4): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[prng.Poisson(4)]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/100))
	}

	fmt.Println("=====Testing for PoissonGenerator end=====")
	fmt.Println()
}

func TestGeometryGenerator(t *testing.T) {
	fmt.Println("=====Testing for GeometryGenerator begin=====")

	grng := NewGeometricGenerator(time.Now().UnixNano())
	fmt.Println("Geometry(0.2): ")
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[grng.Geometric(0.2)]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/100))
	}

	fmt.Println("=====Testing for GeometryGenerator end=====")
	fmt.Println()
}

func TestBinomialGenerator(t *testing.T) {
	fmt.Println("=====Testing for BinomialGenerator begin=====")
	bing := NewBinomialGenerator(time.Now().UnixNano())

	fmt.Println("Binomial(10^6, 0.02) = ", bing.Binomial(1000000, 0.02))
	var n int64 = 6
	var p float64 = 0.5
	fmt.Printf("X ~ Binomial(%d, %.2f): \n", n, p)
	hist := map[int64]int{}
	for i := 0; i < 10000; i++ {
		hist[bing.Binomial(n, p)]++
	}

	keys := []int64{}
	for k := range hist {
		keys = append(keys, k)
	}
	SortInt64Slice(keys)

	for _, key := range keys {
		fmt.Printf("%d:\t%s\n", key, strings.Repeat("*", hist[key]/100))
	}

	fmt.Println("=====Testing for BinomialGenerator end=====")
	fmt.Println()
}

func TestBernoulliGenerator(t *testing.T) {
	fmt.Println("=====Testing for BernoulliGenerator begin=====")
	beng := NewBernoulliGenerator(time.Now().UnixNano())
	hist := map[bool]int{}
	for i := 0; i < 10000; i++ {
		hist[beng.Bernoulli_P(0.25)]++
	}
	for k, v := range hist {
		fmt.Printf("%t:\t%s\n", k, strings.Repeat("*", v/500))
	}
	fmt.Printf("ratio: %.2f\n", float32(hist[true])/float32(hist[true]+hist[false]))
	fmt.Println("=====Testing for BernoulliGenerator end=====")
	fmt.Println()
}

func TestUniformGenerator(t *testing.T) {
	urng := NewUniformGenerator(time.Now().UnixNano())

	fmt.Println("=====Testing for UniformGenerator begin=====")
	fmt.Println("Generating 100 random int64s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\t", urng.Int64())
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Generating 100 random int32s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\t", urng.Int32())
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Println("Generating 100 random float64s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%.2f\t", urng.Float64())
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Println("Generating 100 random float32s: ")
	for i := 0; i < 100; i++ {
		fmt.Printf("%.2f\t", urng.Float32())
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Printf("Random number in [0, 1024): %d and %d\n",
		urng.Int32n(1024), urng.Int64n(1024))
	fmt.Printf("Random number in [33421, 987584): %d and %d\n",
		urng.Int32Range(33421, 987584), urng.Int64Range(33421, 987584))

	fmt.Printf("Random number in [0.0, 47.339): %.2f and %.2f\n",
		urng.Float32n(47.339), urng.Float64n(47.339))
	fmt.Printf("Random number in [45.485, 999.458): %.2f and %.2f\n",
		urng.Float32Range(45.485, 999.458), urng.Float64Range(45.485, 999.458))

	names := []string{"leesper", "charlie", "jack", "cage", "luck", "rice", "hasting", "dora", "selina"}
	names1 := make([]interface{}, len(names))
	names2 := make([]interface{}, len(names))

	for i, v := range names {
		names1[i] = interface{}(v)
		names2[i] = interface{}(v)
	}

	fmt.Printf("Original names1: %v\n", names1)
	for i := 0; i < 5; i++ {
		urng.Shuffle(names1)
		fmt.Printf("names1: %v\n", names1)
	}

	fmt.Printf("Original names2: %v\n", names2)
	for i := 0; i < 5; i++ {
		urng.ShuffleRange(names2, 2, 6)
		fmt.Printf("names2: %v\n", names2)
	}

	fmt.Println("=====Testing for UniformGenerator end=====")
}

func SortInt64Slice(slice []int64) {
	sort.Sort(int64slice(slice))
}

type int64slice []int64

func (slice int64slice) Len() int {
	return len(slice)
}

func (slice int64slice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice int64slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SumFloat64Slice(slice []float64) float64 {
	sum := 0.0
	for _, value := range slice {
		sum += value
	}
	return sum
}
