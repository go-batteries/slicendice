bench:
	go test -bench=.

gen:
	go test -c -o reverse_bench


perf.full: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses,LLC-loads,LLC-load-misses,branch-misses,branches,L1-icache-load-misses,cycles,instructions ./reverse_bench -test.bench=BenchmarkReverseFull > result_full.txt 2>&1

	cat result_full.txt


perf.half: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses,LLC-loads,LLC-load-misses,branch-misses,branches,L1-icache-load-misses,cycles,instructions ./reverse_bench -test.bench=BenchmarkReverseHalf > result_half.txt 2>&1

	cat result_half.txt


perf.all: gen 
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses,LLC-loads,LLC-load-misses,branch-misses,branches,L1-icache-load-misses,cycles,instructions ./reverse_bench -test.bench=BenchmarkReverseFull > result_full.txt 2>&1
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses,LLC-loads,LLC-load-misses,branch-misses,branches,L1-icache-load-misses,cycles,instructions ./reverse_bench -test.bench=BenchmarkReverseHalf > result_half.txt 2>&1

	cat result_half.txt result_full.txt

perf.dyn.full: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses,LLC-loads,LLC-load-misses,branch-misses,branches,L1-icache-load-misses,cycles,instructions ./reverse_bench -test.bench=BenchmarkReverseDynamicFull > result_dyn_full.txt 2>&1

	cat result_dyn_full.txt


perf.dyn.half: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses,LLC-loads,LLC-load-misses,branch-misses,branches,L1-icache-load-misses,cycles,instructions ./reverse_bench -test.bench=BenchmarkReverseDynamicHalf > result_dyn_half.txt 2>&1

	cat result_dyn_half.txt


