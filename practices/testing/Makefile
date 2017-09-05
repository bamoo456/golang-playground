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
	go tool pprof -text -nodecount=10 ./testing.test cpu.log

# directly run benchmark on the function of specific package (this will generate an execute test file called lib.test)
bench-ToJSONArray:
	go test -run=NONE -bench=ToJSONArray -cpuprofile=parse.log golang-playground/practices/testing/lib

# show the benchmark report by the log file: parse.log
report-ToJSONArray:
	go tool pprof -text -nodecount=10 ./lib.test parse.log