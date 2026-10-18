# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2026-10-18  
**Platform:** Linux 6.17.0-1022-azure x86_64  
**Go version:** go1.25.6

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	79611495	        14.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2768341	       438.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11581515	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   51108	     21184 ns/op	   16815 B/op	      97 allocs/op
BenchmarkGoValidCELConcurrent-4            	273279553	         4.389 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	97005684	        11.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	100000000	        12.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	147801988	         8.117 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21115621	        57.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1070 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1343268	       891.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   75235	     15777 ns/op	   15735 B/op	      74 allocs/op
BenchmarkGoValidEnum-4                     	92889278	        12.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	213440337	         5.619 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10796986	       110.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13879988	        86.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	213624432	         5.623 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10403468	       115.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV4-4                     	27230618	        45.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV4-4                	 9341829	       129.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidIPV6-4                     	13827632	        86.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundIPV6-4                	 6713460	       179.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	125785287	         9.541 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	11576775	       104.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4884180	       244.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   80706	     15007 ns/op	   15600 B/op	      78 allocs/op
BenchmarkGoValidLT-4                       	202785734	         5.933 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11357274	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	213315626	         5.622 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10698882	       111.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	98485222	        12.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8788416	       136.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	33335772	        35.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8495509	       140.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4572822	       261.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   78468	     15141 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-4                 	94356030	        12.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8668436	       136.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	62753832	        19.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9223485	       130.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4372288	       273.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	100000000	        11.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13092578	        88.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9692058	       125.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   76083	     15453 ns/op	   15533 B/op	      76 allocs/op
BenchmarkGoValidRequired-4                 	128049188	         9.372 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8600461	       139.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	549340968	         2.185 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   81001	     14616 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-4                      	19234611	        62.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2441295	       491.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   93309	     11998 ns/op	     150 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   77734	     15541 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-4                     	18708241	        64.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2784801	       432.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3601612	       334.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   77072	     15425 ns/op	   15501 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 5.622 / 0 allocs | 111.5 / 0 allocs | **19.8x** | N/A | N/A | N/A | N/A |
| Enum | 12.80 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.31 / 0 allocs | 1070 / 88 B / 5 allocs | **18.7x** | 891.7 / 0 allocs | **15.6x** | 15777 / 15735 B / 74 allocs | **275.3x** |
| GTE | 5.623 / 0 allocs | 115.4 / 0 allocs | **20.5x** | N/A | N/A | N/A | N/A |
| MinLength | 19.01 / 0 allocs | 130.2 / 0 allocs | **6.8x** | 273.1 / 32 B / 2 allocs | **14.4x** | N/A | N/A |
| UUID | 64.09 / 0 allocs | 432.5 / 0 allocs | **6.7x** | 334.4 / 0 allocs | **5.2x** | 15425 / 15501 B / 74 allocs | **240.7x** |
| MaxItems | 12.18 / 0 allocs | 136.6 / 0 allocs | **11.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 35.85 / 0 allocs | 140.8 / 0 allocs | **3.9x** | 261.7 / 32 B / 2 allocs | **7.3x** | 15141 / 15632 B / 80 allocs | **422.3x** |
| LT | 5.933 / 0 allocs | 106.1 / 0 allocs | **17.9x** | N/A | N/A | N/A | N/A |
| MinItems | 12.48 / 0 allocs | 136.2 / 0 allocs | **10.9x** | N/A | N/A | N/A | N/A |
| Alpha | 14.74 / 0 allocs | 438.8 / 0 allocs | **29.8x** | 110.5 / 0 allocs | **7.5x** | 21184 / 16815 B / 97 allocs | **1437.2x** |
| Required | 9.372 / 0 allocs | 139.8 / 0 allocs | **14.9x** | 2.185 / 0 allocs | **0.2x** | 14616 / 15472 B / 72 allocs | **1559.5x** |
| IPV4 | 45.45 / 0 allocs | 129.9 / 0 allocs | **2.9x** | N/A | N/A | N/A | N/A |
| Length | 9.541 / 0 allocs | 104.0 / 0 allocs | **10.9x** | 244.8 / 32 B / 2 allocs | **25.7x** | 15007 / 15600 B / 78 allocs | **1572.9x** |
| IPV6 | 86.97 / 0 allocs | 179.0 / 0 allocs | **2.1x** | N/A | N/A | N/A | N/A |
| URL | 62.10 / 0 allocs | 491.2 / 144 B / 1 allocs | **7.9x** | 11998 / 150 B / 1 allocs | **193.2x** | 15541 / 15641 B / 75 allocs | **250.3x** |
| Numeric | 11.01 / 0 allocs | 88.51 / 0 allocs | **8.0x** | 125.0 / 0 allocs | **11.4x** | 15453 / 15533 B / 76 allocs | **1403.5x** |
| GT | 5.619 / 0 allocs | 110.8 / 0 allocs | **19.7x** | 86.87 / 0 allocs | **15.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 4.389 | 0 allocs |
| CELMultipleExpressions | 11.89 | 0 allocs |
| CELBasic | 12.01 | 0 allocs |
| CELCrossField | 8.117 | 0 allocs |
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
