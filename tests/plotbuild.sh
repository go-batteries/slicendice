#!/bin/bash

# Function to extract data from a file using awk
extract_data() {
    input_file=$1
    output_file=$2

    awk -v output_file="$output_file" '
    BEGIN {
        # Print the header only once
        printf("Benchmark\tCacheReferences\tCacheMisses\tL1DcacheLoads\tL1DcacheMisses\tLLCLoads\tLLCMisses\tBranches\tCycles\tInstructions\n") > output_file
    }
    /BenchmarkReverse/ {
        benchmark = $1
    }
    /cache-references:u/ { gsub(",", "", $1); cache_references = $1 }
    /cache-misses:u/ { gsub(",", "", $1); cache_misses = $1 }
    /L1-dcache-loads:u/ { gsub(",", "", $1); l1_dcache_loads = $1 }
    /L1-dcache-load-misses:u/ { gsub(",", "", $1); l1_dcache_misses = $1 }
    /LLC-loads:u/ { gsub(",", "", $1); llc_loads = $1 }
    /LLC-load-misses:u/ { gsub(",", "", $1); llc_misses = $1 }
    /branches:u/ { gsub(",", "", $1); branches = $1 }
    /cycles:u/ { gsub(",", "", $1); cycles = $1 }
    /instructions:u/ { gsub(",", "", $1); instructions = $1 }
    END {
        if (length(benchmark) > 0) {
            printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
                   benchmark, cache_references, cache_misses, l1_dcache_loads,
                   l1_dcache_misses, llc_loads, llc_misses, branches, cycles, instructions) > output_file
        }
    }
    ' "${input_file}"
}

# Extract data from result_half.txt and result_full.txt
extract_data "result_half.txt" "data_half.txt"
extract_data "result_full.txt" "data_full.txt"

