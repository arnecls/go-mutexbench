# go-mutexbench

Personal playground for concurrency benchmarks.
This code is purely educational and should not be used in any other project.

## Findings

1. Defer has a higher performance impact than expected
1. Atomic functions are "semantically inlined" by the compiler
1. Writing your own atomic functions does not help much because of missing inlining
1. Mutex with defer is faster than without defer under high congestion

## Results

Taken on an Intel Core i7 (Gen4) 2.2 Ghz using Golang 1.8.1.

- Mutex uses sync.Mutex
- Spinlock uses a minimal spinlock using atomic.CompareAndSwap
- Spinlock uses a minimal spinlock using the Intel BitTestAndSet instruction

### 10.000.000 loops single threaded

|test|ms|speedup|
|----|--|-------|
|plain           |442.04|1|
|mutex           |543.80|0.81|
|spinlock        |523.28|0.84|
|spinlock2       |627.49|0.7|
|mutex defer     |1053.86|0.41|
|spinlock defer  |1059.52|0.41|
|spinlock2 defer |1039.80|0.42|

### 1.000.000 loops 8 go routines

|test|ms|speedup|
|----|--|-------|
|plain           |1495.79|1|
|mutex           |1789.38|0.83|
|spinlock        |884.75|1.69|
|spinlock2       |926.29|1.61|
|mutex defer     |1818.81|0.82|
|spinlock defer  |1808.71|0.82|
|spinlock defer2 |1852.24|0.80|

### 1.000.000 loops 16 go routines

|test|ms|speedup|
|----|--|-------|
|plain           |3210.73|1|
|mutex           |4265.63|0.75|
|spinlock        |1803.39|1.78|
|spinlock2       |1882.31|1.70|
|mutex defer     |3889.62|0.82|
|spinlock defer  |3631.06|0.88|
|spinlock defer2 |3842.22|0.83|

### 1.000.000 loops 32 go routines

|test|ms|speedup|
|----|--|-------|
|plain           |6784.43|1|
|mutex           |9348.37|0.72|
|spinlock        |3737.34|1.81|
|spinlock2       |3916.39|1.73|
|mutex defer     |8037.36|0.84|
|spinlock defer  |7976.36|0.85|
|spinlock defer2 |8423.33|0.80|