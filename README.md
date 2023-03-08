# neo4j-test

## cpu
AMD Ryzen 3 3200U with Radeon Vega Mobile Gfx  

## restrictions
NEO4J_dbms_memory_pagecache_size=1G <br/>
NEO4J_dbms.memory.heap.initial_size=1G <br/>
NEO4J_dbms_memory_heap_max__size=1G

## benchmarks
### binary tree
10k users traversal
```
BenchmarkGetPartnersFrom1To5Lvl10kVerticesTree-4            157           7686271 ns/op // 0,00769s
BenchmarkGetPartnersFrom5To10Lvl10kVerticesTree-4            21          62051439 ns/op // 0,06205s
BenchmarkGetPartnersFrom10To14Lvl10kVerticesTree-4            5         209788815 ns/op // 0,20979s
BenchmarkGetPartnersFrom1To14Lvl10kVerticesTree-4             5         234384866 ns/op // 0,23438s
```

100k users traversal
```
BenchmarkGetPartnersFrom16To16Lvl100kVerticesTree-4            2         982297238 ns/op // 0,98229s
BenchmarkGetPartnersFrom13To16Lvl100kVerticesTree-4            1        2037339449 ns/op // 2,03734s
BenchmarkGetPartnersFrom10To13Lvl100kVerticesTree-4            4         343504682 ns/op // 0,34350s
BenchmarkGetPartnersFrom7To10Lvl100kVerticesTree-4            24          45678948 ns/op // 0,04568s
BenchmarkGetPartnersFrom1To16Lvl100kVerticesTree-4             1        2325799834 ns/op // 2,32580s
```