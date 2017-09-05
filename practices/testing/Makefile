.PHONY: all

test:
	go test

# running all benchmark tests
bench:
	go test -bench=. -benchmem

# show the coverage report
report:
	go test -run=Coverage -coverprofile=c.out && go tool cover -html=c.out

# benchmark the function CalcAvg by benchmark function: "BenchmarkCalcAvg"
bench-CalcAvg:
	go test -run=NONE -bench=BenchmarkCalcAvg -cpuprofile=cpu.log

# show the benchmark report by log file: cpu.log
report-CalcAvg:
	go tool pprof -text -nodecount=10 cpu.log
