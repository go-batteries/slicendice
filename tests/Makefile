gen:
	go test -c -o reverse_bench


perf.half: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses ./reverse_bench -test.bench=BenchmarkReverseHalf > result_half.txt 2>&1

pref.full: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses ./reverse_bench -test.bench=BenchmarkReverseFull > result_full.txt 2>&1

pref.block: gen
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses ./reverse_bench -test.bench=BenchmarkReverseFull2 > result_full_2.txt 2>&1

perf.all: gen 
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses ./reverse_bench -test.bench=BenchmarkReverseFull > result_full.txt 2>&1
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses ./reverse_bench -test.bench=BenchmarkReverseFull2 > result_full_2.txt 2>&1
	perf stat -e cache-references,cache-misses,L1-dcache-loads,L1-dcache-load-misses ./reverse_bench -test.bench=BenchmarkReverseHalf > result_half.txt 2>&1
