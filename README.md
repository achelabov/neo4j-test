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

100k users traversal <br/>
getting only vertices
```
BenchmarkGetPartnersFrom16To16Lvl100kVerticesTree-4            2         982297238 ns/op // 0,98229s
BenchmarkGetPartnersFrom13To16Lvl100kVerticesTree-4            1        2037339449 ns/op // 2,03734s
BenchmarkGetPartnersFrom10To13Lvl100kVerticesTree-4            4         343504682 ns/op // 0,34350s
BenchmarkGetPartnersFrom7To10Lvl100kVerticesTree-4            24          45678948 ns/op // 0,04568s
BenchmarkGetPartnersFrom1To16Lvl100kVerticesTree-4             1        2325799834 ns/op // 2,32580s
```

getting only edges
```
BenchmarkGetPartnersFrom16To16Lvl100kVerticesTree-4            1        10930561534 ns/op // 10,93056s
BenchmarkGetPartnersFrom13To16Lvl100kVerticesTree-4            1        22647509097 ns/op // 22,64751s
BenchmarkGetPartnersFrom10To13Lvl100kVerticesTree-4            1         3121065196 ns/op // 3,121065s
BenchmarkGetPartnersFrom7To10Lvl100kVerticesTree-4             3          354320346 ns/op // 0,354320s
BenchmarkGetPartnersFrom1To16Lvl100kVerticesTree-4             1        24313401696 ns/op // 24,31340s
```

getting vertices with edges
```
BenchmarkGetPartnersFrom16To16Lvl100kVerticesTree-4            1        12070492489 ns/op // 12,07049s
BenchmarkGetPartnersFrom13To16Lvl100kVerticesTree-4            1        25704931313 ns/op // 25,70493s
BenchmarkGetPartnersFrom10To13Lvl100kVerticesTree-4            1         3655351610 ns/op // 3,655352s
BenchmarkGetPartnersFrom7To10Lvl100kVerticesTree-4             3          408181904 ns/op // 0,408182s
BenchmarkGetPartnersFrom1To16Lvl100kVerticesTree-4             1        28312419840 ns/op // 28,31242s
```

getting the path to each vertex
```
BenchmarkGetPartnersFrom16To16Lvl100kVerticesTree-4            1        16692635436 ns/op // 16,69264s
BenchmarkGetPartnersFrom13To16Lvl100kVerticesTree-4            1        36030002807 ns/op // 36,03000s
BenchmarkGetPartnersFrom10To13Lvl100kVerticesTree-4            1         4827670956 ns/op // 4,827670s
BenchmarkGetPartnersFrom7To10Lvl100kVerticesTree-4             2          562781340 ns/op // 0,562781s
BenchmarkGetPartnersFrom1To16Lvl100kVerticesTree-4             1        38732582739 ns/op // 38,73258s
```

getting main bonus
```
BenchmarkGetMainBonusBronze-4                163           7810681 ns/op // 0,00781s
BenchmarkGetMainBonusBronzePro-4             156           8152347 ns/op // 0,00815s
BenchmarkGetMainBonusSilver-4                140           9102428 ns/op // 0,00910s
BenchmarkGetMainBonusSilverPro-4             120           8582092 ns/op // 0,00858s
BenchmarkGetMainBonusGold-4                  100          10623066 ns/op // 0,01062s
BenchmarkGetMainBonusGoldPro-4                62          17159427 ns/op // 0,01716s
BenchmarkGetMainBonusPlatinum-4               51          24927266 ns/op // 0,02493s
BenchmarkGetMainBonusPlatinimPro-4            25          55779393 ns/op // 0,05578s
BenchmarkGetMainBonusDiamond-4                12          99397576 ns/op // 0,09940s
```
