### Basic test command

All the tests must be wrote in test files `xxx_test.go` with `TestXxxx` format
```
// Example test function

func TestCalcAvg(t *testing.T) {
	tests := []struct {
		data []float64
		want float64
	}{
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 1, 1}, 1},
		{[]float64{1.5, 1.5, 1.5}, 1.5},
		// failure case
		//{[]float64{1.5, 1.5, 1.5}, 2},
	}

	for _, test := range tests {
		if actual := CalcAvg(test.data...); actual != test.want {
			t.Errorf("CalcAvg(): input %v, got actual= %v != want=%v", test.data, actual, test.want)
		}
	}
}

```

##### test command usage

```
# run the tests in the current package
go test -v -cover

# run all tests under current package (include sub packages) 
go test -v -cover ./.

# generate the coverage report
go test -cover -coverprofile=c.out

# parse the coverage report with html format
go tool cover -html=c.out

```

### Basic benchmark command
All the tests must be wrote in test files `xxx_test.go` with `BenchmarkXxxx` format
```
// Example test function

func BenchmarkCalcAvg(b *testing.B) {
	test := struct {
		data []float64
		want float64
	}{
		[]float64{301020.495, 19284.122, 45020101.111},
		2,
	}
	for i := 0; i < b.N; i++ {
		CalcAvg(test.data...)
	}
}
```
> We can have multiple benchmark function for an target function. e.g.: BenchmarkCalcAvgBasic, BenchmarkCalcHeavyLoad 


##### test command usage
```
go test -run={the test file name} -bench={the benchmark function name} -cpuprofile={the output filename} {the target package}

# run all the bechmark functions
go test -bench=.

# run all the bechmark functions with memory info
go test -bench=. -benchmem

# run the specific benchmark functions
go test -run=NONE -bench=BenchmarkCalcAvg

```
