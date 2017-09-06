.PHONY: all

test:
	go test


# running all benchmark tests
bench:
	go test -bench=. -benchmem


# show the coverage report
report:
	go test -cover -coverprofile=c.out && go tool cover -html=c.out


# benchmark the function CalcAvg by benchmark function: "BenchmarkCalcAvg"
# the -run=NONE means only benchmark the target function
bench-CalcAvg:
	go test -run=NONE -bench=CalcAvg -cpuprofile=cpu.log


# show the benchmark report by log file: cpu.log
report-CalcAvg:
	go tool pprof -text -nodecount=10 ./testing.test cpu.log


# directly run benchmark on the function of specific package (this will generate an execute test file called lib.test)
profile-ToJSONArray:
	go test -run=NONE -bench=ToJSONArray -cpuprofile=profile.log golang-playground/practices/testing/lib


# show the benchmark report by the log file: parse.log
report-profile-ToJSONArray:
	go tool pprof -text -nodecount=10 ./lib.test profile.log