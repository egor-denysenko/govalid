# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-02-04  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	123867987	         9.679 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2365486	       484.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11181042	       108.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55848	     21101 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	653240310	         1.836 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	293571464	         4.157 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295922769	         4.051 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	320560426	         3.742 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21590716	        55.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1089 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1302408	       922.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   66363	     17538 ns/op	   15862 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	256491237	         4.686 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	428200845	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11323077	       106.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13794576	        88.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	428215460	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10938906	       110.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	30601495	        38.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9224904	       129.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	14105614	        87.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6863114	       175.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	148167985	         8.096 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9977743	       120.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4884099	       245.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74895	     16107 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	428662930	         2.800 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11358706	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	385262476	         3.112 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11211843	       107.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	203023078	         5.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8270268	       145.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37399899	        31.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8630415	       139.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4231506	       283.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73608	     16336 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	175157790	         6.847 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8542378	       141.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	47855552	        24.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10127950	       118.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4244151	       281.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	123896409	         9.682 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13776354	        87.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9452694	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71845	     16577 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296407322	         4.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7146729	       167.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642689877	         1.877 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75859	     15662 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20715324	        57.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2506796	       478.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   95451	     12648 ns/op	     147 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71623	     16732 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	22199534	        55.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2632834	       460.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3571850	       336.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71616	     16800 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 3.112 / 0 allocs | 107.3 / 0 allocs | **34.5x** | N/A | N/A | N/A | N/A |
| Enum | 4.686 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.36 / 0 allocs | 1089 / 89 B / 5 allocs | **19.7x** | 922.3 / 0 allocs | **16.7x** | 17538 / 15862 B / 76 allocs | **316.8x** |
| GTE | 2.801 / 0 allocs | 110.0 / 0 allocs | **39.3x** | N/A | N/A | N/A | N/A |
| MinLength | 24.63 / 0 allocs | 118.7 / 0 allocs | **4.8x** | 281.5 / 32 B / 2 allocs | **11.4x** | N/A | N/A |
| UUID | 55.24 / 0 allocs | 460.8 / 0 allocs | **8.3x** | 336.5 / 0 allocs | **6.1x** | 16800 / 15542 B / 76 allocs | **304.1x** |
| MaxItems | 5.915 / 0 allocs | 145.1 / 0 allocs | **24.5x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.77 / 0 allocs | 139.2 / 0 allocs | **4.4x** | 283.2 / 32 B / 2 allocs | **8.9x** | 16336 / 15648 B / 81 allocs | **514.2x** |
| LT | 2.800 / 0 allocs | 105.5 / 0 allocs | **37.7x** | N/A | N/A | N/A | N/A |
| MinItems | 6.847 / 0 allocs | 141.1 / 0 allocs | **20.6x** | N/A | N/A | N/A | N/A |
| Alpha | 9.679 / 0 allocs | 484.8 / 0 allocs | **50.1x** | 108.4 / 0 allocs | **11.2x** | 21101 / 16937 B / 101 allocs | **2180.1x** |
| Required | 4.046 / 0 allocs | 167.8 / 0 allocs | **41.5x** | 1.877 / 0 allocs | **0.5x** | 15662 / 15488 B / 73 allocs | **3871.0x** |
| IPV4 | 38.92 / 0 allocs | 129.7 / 0 allocs | **3.3x** | N/A | N/A | N/A | N/A |
| Length | 8.096 / 0 allocs | 120.0 / 0 allocs | **14.8x** | 245.6 / 32 B / 2 allocs | **30.3x** | 16107 / 15616 B / 79 allocs | **1989.5x** |
| IPV6 | 87.10 / 0 allocs | 175.5 / 0 allocs | **2.0x** | N/A | N/A | N/A | N/A |
| URL | 57.84 / 0 allocs | 478.2 / 144 B / 1 allocs | **8.3x** | 12648 / 147 B / 1 allocs | **218.7x** | 16732 / 15681 B / 77 allocs | **289.3x** |
| Numeric | 9.682 / 0 allocs | 87.21 / 0 allocs | **9.0x** | 128.3 / 0 allocs | **13.3x** | 16577 / 15574 B / 78 allocs | **1712.1x** |
| GT | 2.801 / 0 allocs | 106.8 / 0 allocs | **38.1x** | 88.04 / 0 allocs | **31.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.836 | 0 allocs |
| CELMultipleExpressions | 4.157 | 0 allocs |
| CELBasic | 4.051 | 0 allocs |
| CELCrossField | 3.742 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.000 | 0 allocs |

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

## Running Benchmarks

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=10s

# Run specific validator benchmarks
go test ./benchmark/ -bench=BenchmarkGoValid{ValidatorName} -benchmem
```
