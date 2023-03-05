# neo4j-test

## cpu
AMD Ryzen 3 3200U with Radeon Vega Mobile Gfx  

## restrictions
NEO4J_dbms_memory_pagecache_size=1G1 <br/>
NEO4J_dbms.memory.heap.initial_size=1G <br/>
NEO4J_dbms_memory_heap_max__size=1G

## benchmarks
### binary tree
10k users traversal
```
BenchmarkGetPartnersFrom1To5Lvl-4            157           7686271 ns/op // 0,00769s
BenchmarkGetPartnersFrom5To10Lvl-4            21          62051439 ns/op // 0,06205s
BenchmarkGetPartnersFrom10To14Lvl-4            5         209788815 ns/op // 0,20979s
BenchmarkGetPartnersFrom1To14Lvl-4             5         234384866 ns/op // 0,23438s
```