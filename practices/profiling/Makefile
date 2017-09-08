.PHONY: all

# directly run benchmark on the function of specific package (this will generate an execute test file called profiling.test)
cpu-profiling:
	go test -run=NONE -bench=ToJSONArray -cpuprofile=cpu.log

cpu-profiling-report:
	go tool pprof -nodecount=10 ./profiling.test cpu.log

# show the benchmark report by the log file: parse.log
cpu-profiling-report-text:
	go tool pprof -text -nodecount=10 ./profiling.test cpu.log

cpu-profiling-report-pdf:
	go tool pprof -pdf -nodecount=10 ./profiling.test cpu.log > tmp.pdf

mem-profiling:
	go test -run=NONE -bench=ToJSONArray -memprofile=mem.log

mem-profiling-report:
	go tool pprof -nodecount=10 ./profiling.test mem.log

mem-profiling-report-text:
	go tool pprof -text -nodecount=10 ./profiling.test mem.log
